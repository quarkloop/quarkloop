package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/db"
	"github.com/quarkloop/quarkloop/pkg/server"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

func init() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
}

func initDatabases(sysDatabse db.Database, projectDatabse db.Database, authDatabse db.Database) *repository.Repository {
	sysDatabse.Connect()
	projectDatabse.Connect()
	authDatabse.Connect()

	return &repository.Repository{
		SystemDbConn:  sysDatabse.GetConnection(),
		ProjectDbConn: projectDatabse.GetConnection(),
		AuthDbConn:    authDatabse.GetConnection(),
	}
}

func startApiServer() {
	repository := initDatabases(db.NewSystemDatabase(), db.NewProjectDatabase(), db.NewAuthDatabase())
	api := api.NewServerApi(repository)

	serve := server.NewDefaultServer(repository)
	serve.BindHandlers(&api)

	fmt.Println("Server running on port 8000")
	serve.Router().Run(":8000")
}

func main() {
	startApiServer()
}
