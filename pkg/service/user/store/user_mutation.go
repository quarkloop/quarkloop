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

/// UpdateUserQuery

const updateUserQuery = `
`

func (store *userStore) UpdateUser(ctx context.Context, cmd *user.UpdateUserCommand) error {
	panic("not implemented")
}

/// UpdateUsernameQuery

const updateUsernameQuery = `
`

func (store *userStore) UpdateUsername(ctx context.Context, cmd *user.UpdateUsernameCommand) error {
	panic("not implemented")
}

/// UpdatePasswordQuery

const updatePasswordQuery = `
`

func (store *userStore) UpdatePassword(ctx context.Context, cmd *user.UpdatePasswordCommand) error {
	panic("not implemented")
}

/// UpdatePreferencesQuery

const updatePreferencesQuery = `
`

func (store *userStore) UpdatePreferences(ctx context.Context, cmd *user.UpdatePreferencesCommand) error {
	panic("not implemented")
}

/// UpdateUserByIdQuery

const updateUserByIdQuery = `
UPDATE
    "auth"."User"
SET
    "email"         = @email,
    "emailVerified" = @emailVerified,
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

func (store *userStore) UpdateUserById(ctx context.Context, cmd *user.UpdateUserByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updateUserByIdQuery, pgx.NamedArgs{
		"id":            cmd.UserId,
		"email":         cmd.Email,
		"emailVerified": cmd.EmailVerified,
		"name":          cmd.Name,
		"birthdate":     cmd.Birthdate,
		"country":       cmd.Country,
		"image":         cmd.Image,
		"status":        cmd.Status,
		"updatedBy":     cmd.UpdatedBy,
		"updatedAt":     time.Now(),
	})
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

/// UpdateUsernameByUserIdQuery

const updateUsernameByUserIdQuery = `
UPDATE
    "auth"."User"
SET
    "username"  = @username,
    "updatedAt" = @updatedAt,
    "updatedBy" = @updatedBy,
WHERE
    "id" = @id;
`

func (store *userStore) UpdateUsernameByUserId(ctx context.Context, cmd *user.UpdateUsernameByUserIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updateUsernameByUserIdQuery, pgx.NamedArgs{
		"id":        cmd.UserId,
		"username":  cmd.Username,
		"updatedBy": cmd.UpdatedBy,
		"updatedAt": time.Now(),
	})
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

/// UpdatePasswordByUserIdQuery

const updatePasswordByUserIdQuery = `
UPDATE
    "auth"."User"
SET
    "password"  = @password,
    "updatedAt" = @updatedAt,
    "updatedBy" = @updatedBy,
WHERE
    "id" = @id;
`

func (store *userStore) UpdatePasswordByUserId(ctx context.Context, cmd *user.UpdatePasswordByUserIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updatePasswordByUserIdQuery, pgx.NamedArgs{
		"id":        cmd.UserId,
		"password":  cmd.Password,
		"updatedBy": cmd.UpdatedBy,
		"updatedAt": time.Now(),
	})
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

/// UpdatePreferencesByUserIdQuery

const updatePreferencesByUserIdQuery = `
`

func (store *userStore) UpdatePreferencesByUserId(ctx context.Context, cmd *user.UpdatePreferencesByUserIdCommand) error {
	panic("not implemented")
}

/// DeleteUserByIdQuery

const deleteUserByIdQuery = `
DELETE FROM
    "auth"."User"
WHERE
    "id" = @id;
`

func (store *userStore) DeleteUserById(ctx context.Context, cmd *user.DeleteUserByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteUserByIdQuery, pgx.NamedArgs{"id": cmd.UserId})
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

/// DeleteSessionByIdQuery

const deleteSessionByIdQuery = `
DELETE FROM
    "auth"."Session"
WHERE
    "id" = @id
AND
    "userId" = @userId;
`

func (store *userStore) DeleteSessionById(ctx context.Context, cmd *user.DeleteSessionByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteSessionByIdQuery, pgx.NamedArgs{
		"id":     cmd.SessionId,
		"userId": cmd.UserId,
	})
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

/// deleteAccountByIdQuery

const DeleteAccountByIdQuery = `
DELETE FROM
    "auth"."Account"
WHERE
    "id" = @id
AND
    "userId" = @userId;
`

func (store *userStore) DeleteAccountById(ctx context.Context, cmd *user.DeleteAccountByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteSessionByIdQuery, pgx.NamedArgs{
		"id":     cmd.AccountId,
		"userId": cmd.UserId,
	})
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
