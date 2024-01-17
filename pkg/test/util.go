package test

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/quarkloop/quarkloop/pkg/db"
	"github.com/quarkloop/quarkloop/pkg/util"
)

func InitTestAuthDB() (context.Context, *pgx.Conn) {
	err := godotenv.Load(util.GetTestEnvFilePath())
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	db := db.NewAuthDatabase()
	db.Connect()

	return context.Background(), db.GetConnection()
}

func InitTestAuthzDB() (context.Context, *pgx.Conn) {
	err := godotenv.Load(util.GetTestEnvFilePath())
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	db := db.NewAuthzDatabase()
	db.Connect()

	return context.Background(), db.GetConnection()
}

func InitTestSystemDB() (context.Context, *pgx.Conn) {
	err := godotenv.Load(util.GetTestEnvFilePath())
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	db := db.NewSystemDatabase()
	db.Connect()

	return context.Background(), db.GetConnection()
}
