package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
	userErrors "github.com/quarkloop/quarkloop/services/user/errors"
)

/// CreateUser

const createUserQuery = `
INSERT INTO "auth"."User" (
	"username",
	"name",
	"email",
	"birthdate",
	"country",
	"image",
	"status",
    "createdBy"
)
VALUES (
	@username,
	@name,
	@email,
	@birthdate,
	@country,
	@image,
	@status,
    @createdBy
)
RETURNING 
    "id",
    "username",
    "name",
    "email",
    "birthdate",
    "country",
    "image",
    "status",
    "createdAt",
    "createdBy";
`

type CreateUserCommand struct {
	CreatedBy string

	Username  string
	Name      string
	Email     string
	Birthdate *time.Time
	Country   string
	Image     string
	Status    int32
}

func (store *userStore) CreateUser(ctx context.Context, cmd *CreateUserCommand) (*model.User, error) {
	row := store.Conn.QueryRow(ctx, createUserQuery, pgx.NamedArgs{
		"username":  cmd.Username,
		"name":      cmd.Name,
		"email":     cmd.Email,
		"birthdate": cmd.Birthdate,
		"country":   cmd.Country,
		"image":     cmd.Image,
		"status":    cmd.Status,
		"createdBy": cmd.CreatedBy,
	})

	var u model.User
	err := row.Scan(
		&u.Id,
		&u.Username,
		&u.Name,
		&u.Email,
		&u.Birthdate,
		&u.Country,
		&u.Image,
		&u.Status,
		&u.CreatedAt,
		&u.CreatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, userErrors.HandleError(err)
	}

	return &u, nil
}

/// UpdateUserById

const updateUserByIdQuery = `
UPDATE
    "auth"."User"
SET
    "username"      = @username,
    "email"         = @email,
    "emailVerified" = @emailVerified,
    "name"          = @name,
    "birthdate"     = @birthdate,
    "country"       = @country,
    "image"         = @image,
    "status"        = @status,
    "updatedAt"     = @updatedAt,
    "updatedBy"     = @updatedBy
WHERE
    "id" = @id;
`

type UpdateUserByIdCommand struct {
	UserId    int64
	UpdatedBy string

	Username      string
	Name          string
	Email         string
	EmailVerified *time.Time
	Birthdate     *time.Time
	Country       string
	Image         string
	Status        int32
}

func (store *userStore) UpdateUserById(ctx context.Context, cmd *UpdateUserByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updateUserByIdQuery, pgx.NamedArgs{
		"id":            cmd.UserId,
		"username":      cmd.Username,
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

/// UpdateUsernameByUserId

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

type UpdateUsernameByUserIdCommand struct {
	UserId    int64
	UpdatedBy string
	Username  string
}

func (store *userStore) UpdateUsernameByUserId(ctx context.Context, cmd *UpdateUsernameByUserIdCommand) error {
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

/// UpdatePasswordByUserId

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

type UpdatePasswordByUserIdCommand struct {
	UserId    int64
	UpdatedBy string
	Password  string
}

func (store *userStore) UpdatePasswordByUserId(ctx context.Context, cmd *UpdatePasswordByUserIdCommand) error {
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

/// DeleteUserById

const deleteUserByIdQuery = `
DELETE FROM
    "auth"."User"
WHERE
    "id" = @id;
`

type DeleteUserByIdCommand struct {
	UserId int64
}

func (store *userStore) DeleteUserById(ctx context.Context, cmd *DeleteUserByIdCommand) error {
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

/// CreateSession

const createSessionQuery = `
INSERT INTO "auth"."Session" (
	"userId",
	"sessionToken",
	"expires"
)
VALUES (
	@userId,
	@sessionToken,
	@expires
)
RETURNING 
    "id",
    "userId",
    "sessionToken",
    "expires";
`

type CreateSessionCommand struct {
	UserId int64

	SessionToken string
	ExpiresAt    time.Time
}

func (store *userStore) CreateSession(ctx context.Context, cmd *CreateSessionCommand) (*model.Session, error) {
	row := store.Conn.QueryRow(ctx, createSessionQuery, pgx.NamedArgs{
		"userId":       cmd.UserId,
		"sessionToken": cmd.SessionToken,
		"expires":      cmd.ExpiresAt,
	})

	var sess model.Session
	err := row.Scan(
		&sess.Id,
		&sess.UserId,
		&sess.SessionToken,
		&sess.ExpiresAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, userErrors.HandleError(err)
	}

	return &sess, nil
}

/// DeleteSessionById

const deleteSessionByIdQuery = `
DELETE FROM
    "auth"."Session"
WHERE
    "id" = @id
AND
    "userId" = @userId;
`

type DeleteSessionByIdCommand struct {
	UserId    int64
	SessionId int64
}

func (store *userStore) DeleteSessionById(ctx context.Context, cmd *DeleteSessionByIdCommand) error {
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

/// CreateAccount

const createAccountQuery = `
INSERT INTO "auth"."Account" (
	"userId",
	"type",
	"provider",
	"providerAccountId",
	"id_token",	
	"refresh_token",
	"access_token"
)
VALUES (
	@userId,
	@type,
	@provider,
	@providerAccountId,
	@id_token,
	@refresh_token,
	@access_token
)
RETURNING 
    "id",
	"userId",
	"type",
	"provider",
	"providerAccountId",
	"id_token",	
	"refresh_token",
	"access_token";
`

type CreateAccountCommand struct {
	// id
	UserId            int64
	TokenId           *string
	ProviderAccountId string

	// account
	Type         string
	TokenType    *string
	Provider     string
	RefreshToken *string
	AccessToken  *string
	Scope        *string
	SessionState *string
	ExpiresAt    *time.Time
}

func (store *userStore) CreateAccount(ctx context.Context, cmd *CreateAccountCommand) (*model.Account, error) {
	row := store.Conn.QueryRow(ctx, createAccountQuery, pgx.NamedArgs{
		"userId":            cmd.UserId,
		"type":              cmd.Type,
		"provider":          cmd.Provider,
		"providerAccountId": cmd.ProviderAccountId,
		"id_token":          cmd.TokenId,
		"refresh_token":     cmd.RefreshToken,
		"access_token":      cmd.AccessToken,
	})

	var acc model.Account
	err := row.Scan(
		&acc.Id,
		&acc.UserId,
		&acc.Type,
		&acc.Provider,
		&acc.ProviderAccountId,
		&acc.TokenId,
		&acc.RefreshToken,
		&acc.AccessToken,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, userErrors.HandleError(err)
	}

	return &acc, nil
}

/// DeleteAccountById

const deleteAccountByIdQuery = `
DELETE FROM
    "auth"."Account"
WHERE
    "id" = @id
AND
    "userId" = @userId;
`

type DeleteAccountByIdCommand struct {
	UserId    int64
	AccountId int64
}

func (store *userStore) DeleteAccountById(ctx context.Context, cmd *DeleteAccountByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteAccountByIdQuery, pgx.NamedArgs{
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
