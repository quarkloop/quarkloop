package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type systemDatabase struct {
	host       string
	user       string
	password   string
	port       string
	database   string
	connection *pgx.Conn
}

func NewSystemDatabase() *systemDatabase {
	return &systemDatabase{
		host:       os.Getenv("PG_HOST"),
		user:       os.Getenv("PG_USER"),
		password:   os.Getenv("PG_PASSWORD"),
		port:       os.Getenv("PG_PORT"),
		database:   os.Getenv("PG_QUARKLOOP_SYSTEM_DB"),
		connection: nil,
	}
}

func (db *systemDatabase) Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		db.host,
		db.user,
		db.password,
		db.port,
		db.database,
	)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.connection = conn
}

func (db *systemDatabase) Connection() *pgx.Conn {
	return db.connection
}
