package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Database interface {
	Connect()
	GetConnection() *pgx.Conn

	GetDatabase() string
	GetHost() string
	GetPassword() string
	GetPort() string
	GetUser() string
}

type database struct {
	host       string
	user       string
	password   string
	port       string
	database   string
	connection *pgx.Conn
}

func NewAuthDatabase() *database {
	return &database{
		host:       os.Getenv("PGHOST"),
		user:       os.Getenv("PGUSER"),
		password:   os.Getenv("PGPASSWORD"),
		port:       os.Getenv("PGPORT"),
		database:   os.Getenv("PG_QUARKLOOP_AUTH_DB"),
		connection: nil,
	}
}

func NewAuthzDatabase() *database {
	return &database{
		host:       os.Getenv("PGHOST"),
		user:       os.Getenv("PGUSER"),
		password:   os.Getenv("PGPASSWORD"),
		port:       os.Getenv("PGPORT"),
		database:   os.Getenv("PG_QUARKLOOP_AUTHZ_DB"),
		connection: nil,
	}
}

func NewProjectDatabase() *database {
	return &database{
		host:       os.Getenv("PGHOST"),
		user:       os.Getenv("PGUSER"),
		password:   os.Getenv("PGPASSWORD"),
		port:       os.Getenv("PGPORT"),
		database:   os.Getenv("PG_QUARKLOOP_PROJECT_DB"),
		connection: nil,
	}
}

func NewSystemDatabase() *database {
	return &database{
		host:       os.Getenv("PGHOST"),
		user:       os.Getenv("PGUSER"),
		password:   os.Getenv("PGPASSWORD"),
		port:       os.Getenv("PGPORT"),
		database:   os.Getenv("PG_QUARKLOOP_SYSTEM_DB"),
		connection: nil,
	}
}

func NewEngineDatabase() *database {
	return &database{
		host:       os.Getenv("PGHOST"),
		user:       os.Getenv("PGUSER"),
		password:   os.Getenv("PGPASSWORD"),
		port:       os.Getenv("PGPORT"),
		database:   os.Getenv("PG_QUARKLOOP_ENGINE_DB"),
		connection: nil,
	}
}

func (db *database) Connect() {
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

func (db *database) GetDatabase() string {
	return db.database
}

func (db *database) GetHost() string {
	return db.host
}

func (db *database) GetPassword() string {
	return db.password
}

func (db *database) GetPort() string {
	return db.port
}

func (db *database) GetUser() string {
	return db.user
}

func (db *database) GetConnection() *pgx.Conn {
	return db.connection
}
