package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type projectDatabase struct {
	host       string
	user       string
	password   string
	port       string
	database   string
	connection *pgx.Conn
}

func NewProjectDatabase() *projectDatabase {
	return &projectDatabase{
		host:       os.Getenv("PG_HOST"),
		user:       os.Getenv("PG_USER"),
		password:   os.Getenv("PG_PASSWORD"),
		port:       os.Getenv("PG_PORT"),
		database:   os.Getenv("PG_QUARKLOOP_PROJECT_DB"),
		connection: nil,
	}
}

func (db *projectDatabase) Connect() {
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

func (db *projectDatabase) Connection() *pgx.Conn {
	return db.connection
}
