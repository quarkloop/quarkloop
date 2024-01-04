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

/// UpdateUserQuery

const updateUserQuery = `
`

func (store *orgStore) UpdateUser(ctx context.Context, cmd *user.UpdateUserCommand) error {

}

/// UpdateUsernameQuery

const updateUsernameQuery = `
`

func (store *orgStore) UpdateUsername(ctx context.Context, cmd *user.UpdateUsernameCommand) error {

}

/// UpdatePasswordQuery

const updatePasswordQuery = `
`

func (store *orgStore) UpdatePassword(ctx context.Context, cmd *user.UpdatePasswordCommand) error {

}

/// UpdatePreferencesQuery

const updatePreferencesQuery = `
`

func (store *orgStore) UpdatePreferences(ctx context.Context, cmd *user.UpdatePreferencesCommand) error {

}

/// UpdateUserByIdQuery

const updateUserByIdQuery = `
`

func (store *orgStore) UpdateUserById(ctx context.Context, cmd *user.UpdateUserByIdCommand) error {

}

/// UpdateUsernameByUserIdQuery

const updateUsernameByUserIdQuery = `
`

func (store *orgStore) UpdateUsernameByUserId(ctx context.Context, cmd *user.UpdateUsernameByUserIdCommand) error {

}

/// UpdatePasswordByUserIdQuery

const updatePasswordByUserIdQuery = `
`

func (store *orgStore) UpdatePasswordByUserId(ctx context.Context, cmd *user.UpdatePasswordByUserIdCommand) error {

}

/// UpdatePreferencesByUserIdQuery

const updatePreferencesByUserIdQuery = `
`

func (store *orgStore) UpdatePreferencesByUserId(ctx context.Context, cmd *user.UpdatePreferencesByUserIdCommand) error {

}

/// DeleteUserByIdQuery

const deleteUserByIdQuery = `
`

func (store *orgStore) DeleteUserById(ctx context.Context, cmd *user.DeleteUserByIdCommand) error {

}

/// DeleteSessionByIdQuery

const deleteSessionByIdQuery = `
`

func (store *orgStore) DeleteSessionById(ctx context.Context, cmd *user.DeleteSessionByIdCommand) error {

}

/// deleteAccountByIdQuery

const DeleteAccountByIdQuery = `
`

func (store *orgStore) DeleteAccountById(ctx context.Context, cmd *user.DeleteAccountByIdCommand) error {

}
