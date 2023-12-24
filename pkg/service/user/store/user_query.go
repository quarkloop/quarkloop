package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/service/user"
)

/// GetUserById

const getUserByIdQuery = `
SELECT 
	"id",
    "name",
    "email",
    "emailVerified",
    "password",
    "birthdate",
    "country",
    "image",
    "status",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
	"auth"."User"
WHERE 
	"id" = @id;
`

func (store *orgStore) GetUserById(ctx context.Context, userId int) (*user.User, error) {
	row := store.Conn.QueryRow(ctx, getUserByIdQuery, pgx.NamedArgs{"id": userId})

	var user user.User
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.EmailVerified,
		&user.Password,
		&user.Birthdate,
		&user.Country,
		&user.Image,
		&user.Status,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &user, nil
}

/// GetUserByEmail

const getUserByEmailQuery = `
SELECT 
	"id",
    "name",
    "email",
    "emailVerified",
    "password",
    "birthdate",
    "country",
    "image",
    "status",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
	"auth"."User"
WHERE 
	"email" = @email;
`

func (store *orgStore) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	row := store.Conn.QueryRow(ctx, getUserByEmailQuery, pgx.NamedArgs{"email": email})

	var user user.User
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.EmailVerified,
		&user.Password,
		&user.Birthdate,
		&user.Country,
		&user.Image,
		&user.Status,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &user, nil
}

/// GetUserAccountByUserId

const getUserAccountByUserIdQuery = `
SELECT 
	"id",
    "type",
    "provider",
    "providerAccountId",
    "refresh_token",
    "access_token",
    "expires_at",
    "token_type",
    "scope",
    "id_token",
    "session_state"
FROM 
	"auth"."Account"
WHERE 
	"userId" = @userId;
`

func (store *orgStore) GetUserAccountByUserId(ctx context.Context, userId int) (*user.UserAccount, error) {
	row := store.Conn.QueryRow(ctx, getUserAccountByUserIdQuery, pgx.NamedArgs{"userId": userId})

	var acc user.UserAccount
	err := row.Scan(
		&acc.Id,
		&acc.Type,
		&acc.Provider,
		&acc.ProviderAccountId,
		&acc.RefreshToken,
		&acc.AccessToken,
		&acc.ExpiresAt,
		&acc.TokenType,
		&acc.Scope,
		&acc.TokenId,
		&acc.SessionState,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &acc, nil
}

/// GetUserSessionByUserId

const getUserSessionByUserIdQuery = `
SELECT 
	"id",
    "sessionToken",
    "expires"
FROM 
	"auth"."Session"
WHERE 
	"userId" = @userId;
`

func (store *orgStore) GetUserSessionByUserId(ctx context.Context, userId int) (*user.UserSession, error) {
	row := store.Conn.QueryRow(ctx, getUserSessionByUserIdQuery, pgx.NamedArgs{"userId": userId})

	var session user.UserSession
	err := row.Scan(
		&session.Id,
		&session.SessionToken,
		&session.ExpiresAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &session, nil
}
