package repository

import "github.com/jackc/pgx/v5"

type Repository struct {
	AuthDbConn    *pgx.Conn
	SystemDbConn  *pgx.Conn
	ProjectDbConn *pgx.Conn
}
