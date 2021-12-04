package main

import (
	"fmt"
	"github.com/AliceDiNunno/cc-user/src/adapters/persistence/postgres"
	"github.com/AliceDiNunno/cc-user/src/adapters/rest"
	"github.com/AliceDiNunno/cc-user/src/config"
	"github.com/AliceDiNunno/cc-user/src/core/usecases"
	"gorm.io/gorm"
	"log"
)

func main() {
	config.LoadEnv()

	ginConfiguration := config.LoadGinConfiguration()
	dbConfig := config.LoadGormConfiguration()

	var db *gorm.DB
	if dbConfig.Engine == "POSTGRES" {
		db = postgres.StartGormDatabase(dbConfig)
		/*err := db.AutoMigrate(&postgres.Project{})
		if err != nil {
			log.Fatalln(err)
		}*/
	} else {
		log.Fatalln(fmt.Sprintf("Database engine \"%s\" not supported", dbConfig.Engine))
	}

	_ = db

	usecasesHandler := usecases.NewInteractor()

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	rest.SetRoutes(restServer, routesHandler)

	restServer.Start()
}
