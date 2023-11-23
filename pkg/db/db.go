package db

import "github.com/jackc/pgx/v5"

type Database interface {
	Connect()
	Connection() *pgx.Conn
}
