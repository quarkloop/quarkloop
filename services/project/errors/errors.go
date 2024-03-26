package errors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrProjectNotFound      = errors.New("project not found")
	ErrProjectAlreadyExists = errors.New("project with same scopeId already exists")
)

func HandleError(err error) error {
	if err == nil {
		return nil
	}

	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case "22012":
			return errors.New("rows.Err()" + err.Error())
		case "23505":
			return ErrProjectAlreadyExists
		}
	}

	return err
}
