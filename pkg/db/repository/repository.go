package repository

import "github.com/jackc/pgx/v5"

type Repository struct {
	SystemDbConn *pgx.Conn
	AuthDbConn   *pgx.Conn
	AppDbConn    *pgx.Conn
}
