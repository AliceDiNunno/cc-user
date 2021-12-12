package main

import (
	"fmt"
	"github.com/AliceDiNunno/cc-user/src/adapters/persistence/postgres"
	"github.com/AliceDiNunno/cc-user/src/adapters/rest"
	"github.com/AliceDiNunno/cc-user/src/config"
	"github.com/AliceDiNunno/cc-user/src/core/usecases"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()

	ginConfiguration := config.LoadGinConfiguration()
	dbConfig := config.LoadGormConfiguration()
	initialUserConfiguration := config.LoadInitialUserConfiguration()

	var userRepo usecases.UserRepo
	var tokenRepo usecases.UserTokenRepo
	var jwtSignatureRepo usecases.JwtSignatureRepo

	var db *gorm.DB
	if dbConfig.Engine == "POSTGRES" {
		db = postgres.StartGormDatabase(dbConfig)
		err := db.AutoMigrate(&postgres.User{}, &postgres.JwtSignature{}, &postgres.UserToken{})
		if err != nil {
			log.Fatalln(err)
		}

		userRepo = postgres.NewUserRepo(db)
		tokenRepo = postgres.NewUserTokenRepo(db)
		jwtSignatureRepo = postgres.NewJwtSignatureRepo(db)
	} else {
		log.Fatalln(fmt.Sprintf("Database engine \"%s\" not supported", dbConfig.Engine))
	}

	usecasesHandler := usecases.NewInteractor(userRepo, tokenRepo, jwtSignatureRepo)

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	if initialUserConfiguration != nil {
		err := usecasesHandler.CreateInitialUser(initialUserConfiguration)
		if err != nil {
			log.Warning(err.Err.Error())
		}
	}

	rest.SetRoutes(restServer, routesHandler)

	restServer.Start()
}
