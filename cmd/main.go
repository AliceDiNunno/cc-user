package main

import (
	"crypto/rand"
	"fmt"
	"github.com/AliceDiNunno/cc-user/src/adapters/persistence/postgres"
	"github.com/AliceDiNunno/cc-user/src/adapters/rest"
	"github.com/AliceDiNunno/cc-user/src/config"
	"github.com/AliceDiNunno/cc-user/src/core/usecases"
	"gorm.io/gorm"
	"log"
	"math/big"
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret = append(ret, letters[num.Int64()])
	}

	return string(ret), nil
}

func main() {
	config.LoadEnv()

	ginConfiguration := config.LoadGinConfiguration()
	dbConfig := config.LoadGormConfiguration()
	initialUserConfiguration := config.LoadInitialUserConfiguration()

	var userRepo usecases.UserRepo

	var db *gorm.DB
	if dbConfig.Engine == "POSTGRES" {
		db = postgres.StartGormDatabase(dbConfig)
		err := db.AutoMigrate(&postgres.User{})
		if err != nil {
			log.Fatalln(err)
		}

		userRepo = postgres.NewUserRepo(db)
	} else {
		log.Fatalln(fmt.Sprintf("Database engine \"%s\" not supported", dbConfig.Engine))
	}

	_ = db

	usecasesHandler := usecases.NewInteractor(userRepo)

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	if initialUserConfiguration != nil {
		usecasesHandler.CreateInitialUser(initialUserConfiguration)
	}

	rest.SetRoutes(restServer, routesHandler)

	restServer.Start()
}
