package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// GetUserById

const getUserByIdQuery = `
SELECT 
    "id",
	"username",
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

type GetUserByIdQuery struct {
	UserId int64
}

func (store *userStore) GetUserById(ctx context.Context, query *GetUserByIdQuery) (*model.User, error) {
	row := store.Conn.QueryRow(ctx, getUserByIdQuery, pgx.NamedArgs{"id": query.UserId})

	var user model.User
	err := row.Scan(
		&user.Id,
		&user.Username,
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

/// GetUsernameByUserId

const getUsernameByUserIdQuery = `
SELECT 
	"username"
FROM 
    "auth"."User"
WHERE 
    "id" = @id;
`

type GetUsernameByUserIdQuery struct {
	UserId int64
}

func (store *userStore) GetUsernameByUserId(ctx context.Context, query *GetUsernameByUserIdQuery) (string, error) {
	row := store.Conn.QueryRow(ctx, getUserByIdQuery, pgx.NamedArgs{"id": query.UserId})

	var username string
	if err := row.Scan(&username); err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return "", err
	}

	return username, nil
}

/// GetEmailByUserId

const getEmailByUserIdQuery = `
SELECT 
	"email"
FROM 
    "auth"."User"
WHERE 
    "id" = @id;
`

type GetEmailByUserIdQuery struct {
	UserId int64
}

func (store *userStore) GetEmailByUserId(ctx context.Context, query *GetEmailByUserIdQuery) (string, error) {
	row := store.Conn.QueryRow(ctx, getUserByIdQuery, pgx.NamedArgs{"id": query.UserId})

	var email string
	if err := row.Scan(&email); err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return "", err
	}

	return email, nil
}

/// GetStatusByUserId

const getStatusByUserIdQuery = `
SELECT 
	"status"
FROM 
    "auth"."User"
WHERE 
    "id" = @id;
`

type GetStatusByUserIdQuery struct {
	UserId int64
}

func (store *userStore) GetStatusByUserId(ctx context.Context, query *GetStatusByUserIdQuery) (int, error) {
	row := store.Conn.QueryRow(ctx, getUserByIdQuery, pgx.NamedArgs{"id": query.UserId})

	var status int
	if err := row.Scan(&status); err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return 0, err
	}

	return status, nil
}

/// GetSessionsByUserId

const getSessionsByUserIdQuery = `
SELECT 
    "id",
    "sessionToken",
    "expires"
FROM 
    "auth"."Session"
WHERE 
    "userId" = @userId;
`

type GetSessionsByUserIdQuery struct {
	UserId int64
}

func (store *userStore) GetSessionsByUserId(ctx context.Context, query *GetSessionsByUserIdQuery) ([]*model.Session, error) {
	rows, err := store.Conn.Query(ctx, getSessionsByUserIdQuery, pgx.NamedArgs{"userId": query.UserId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var sessionList []*model.Session = []*model.Session{}
	for rows.Next() {
		var session model.Session
		err := rows.Scan(
			&session.Id,
			&session.SessionToken,
			&session.ExpiresAt,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		sessionList = append(sessionList, &session)
	}

	if err = rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows error: %v\n", rows.Err())
		return nil, err
	}

	return sessionList, nil
}

/// GetAccountsByUserId

const getAccountsByUserIdQuery = `
SELECT 
    "id",
    "type",
    "provider",
    "providerAccountId",
    "refresh_token",
    "access_token",
    to_timestamp("expires_at") AS expires_at,
    "token_type",
    "scope",
    "id_token",
    "session_state"
FROM 
    "auth"."Account"
WHERE 
    "userId" = @userId;
`

type GetAccountsByUserIdQuery struct {
	UserId int64
}

func (store *userStore) GetAccountsByUserId(ctx context.Context, query *GetAccountsByUserIdQuery) ([]*model.Account, error) {
	rows, err := store.Conn.Query(ctx, getAccountsByUserIdQuery, pgx.NamedArgs{"userId": query.UserId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var accountList []*model.Account = []*model.Account{}
	for rows.Next() {
		var acc model.Account
		err := rows.Scan(
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
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		accountList = append(accountList, &acc)
	}

	if err = rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows error: %v\n", rows.Err())
		return nil, err
	}

	return accountList, nil
}

/// GetUsers

const getUsersQuery = ``

type GetUsersQuery struct{}

func (store *userStore) GetUsers(ctx context.Context, query *GetUsersQuery) ([]*model.User, error) {
	panic("not implemented")
}
