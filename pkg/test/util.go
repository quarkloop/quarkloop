package test

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/quarkloop/quarkloop/pkg/db"
	"github.com/quarkloop/quarkloop/pkg/util"
)

func InitTestSystemDB() (context.Context, *pgx.Conn) {
	err := godotenv.Load(util.GetDevEnvFilePath())
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	db := db.NewSystemDatabase()
	db.Connect()

	return context.Background(), db.Connection()
}
