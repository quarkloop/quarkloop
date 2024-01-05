package org

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
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
			return ErrOrgAlreadyExists
		}
	}

	return err
}
