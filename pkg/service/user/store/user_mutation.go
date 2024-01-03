package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

/// UpdateUserById

const updateUserByIdMutation = `
UPDATE
    "auth"."User"
SET
    "sid"           = @scopeId,
    "email"         = @email,
    "emailVerified" = @emailVerified,
    "password"      = @password,
    "name"          = @name,
    "birthdate"     = @birthdate,
    "country"       = @country,
    "image"         = @image,
    "status"        = @status,
    "updatedAt"     = @updatedAt,
    "updatedBy"     = @updatedBy,
WHERE
    "id" = @id;
`

func (store *orgStore) UpdateUserById(ctx context.Context, userId int, user *user.User) error {
	commandTag, err := store.Conn.Exec(ctx, updateUserByIdMutation, pgx.NamedArgs{
		"id":            userId,
		"scopeId":       user.ScopeId,
		"email":         user.Email,
		"emailVerified": user.EmailVerified,
		"password":      user.Password,
		"name":          user.Name,
		"birthdate":     user.Birthdate,
		"country":       user.Country,
		"image":         user.Image,
		"status":        user.Status,
		"updatedAt":     time.Now(),
		"updatedBy":     user.UpdatedBy,
	},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to update")
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}

/// DeleteUserById

const deleteUserByIdMutation = `
DELETE FROM
    "auth"."User"
WHERE
    "id" = @id;
`

func (store *orgStore) DeleteUserById(ctx context.Context, userId int) error {
	commandTag, err := store.Conn.Exec(ctx, deleteUserByIdMutation, pgx.NamedArgs{"id": userId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to delete")
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}
