package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/service/user"
)

/// GetUser

const getUserQuery = `
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

func (store *userStore) GetUser(ctx context.Context, query *user.GetUserQuery) (*user.User, error) {
	panic("not implemented")
}

/// GetUsername

const getUsernameQuery = `
SELECT 
	"username"
FROM 
    "auth"."User"
WHERE 
    "id" = @id;
`

func (store *userStore) GetUsername(ctx context.Context, query *user.GetUsernameQuery) (string, error) {
	panic("not implemented")
}

/// GetEmail

const getEmailQuery = `
SELECT 
	"email"
FROM 
    "auth"."User"
WHERE 
    "id" = @id;
`

func (store *userStore) GetEmail(ctx context.Context, query *user.GetEmailQuery) (string, error) {
	panic("not implemented")
}

/// GetStatus

const getStatusQuery = `
SELECT 
	"status"
FROM 
    "auth"."User"
WHERE 
    "id" = @id;
`

func (store *userStore) GetStatus(ctx context.Context, query *user.GetStatusQuery) (any, error) {
	panic("not implemented")
}

/// GetPreferences

const getPreferencesQuery = ``

func (store *userStore) GetPreferences(ctx context.Context, query *user.GetPreferencesQuery) (any, error) {
	panic("not implemented")
}

/// GetSessions

const getSessionsQuery = `
SELECT 
    "id",
    "sessionToken",
    "expires"
FROM 
    "auth"."Session"
WHERE 
    "userId" = @userId;
`

func (store *userStore) GetSessions(ctx context.Context, query *user.GetSessionsQuery) ([]*user.UserSession, error) {
	panic("not implemented")
}

/// GetAccounts

const getAccountsQuery = `
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

func (store *userStore) GetAccounts(ctx context.Context, query *user.GetAccountsQuery) ([]*user.UserAccount, error) {
	panic("not implemented")
}

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

func (store *userStore) GetUserById(ctx context.Context, query *user.GetUserByIdQuery) (*user.User, error) {
	row := store.Conn.QueryRow(ctx, getUserByIdQuery, pgx.NamedArgs{"id": query.UserId})

	var user user.User
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

func (store *userStore) GetUsernameByUserId(ctx context.Context, query *user.GetUsernameByUserIdQuery) (string, error) {
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

func (store *userStore) GetEmailByUserId(ctx context.Context, query *user.GetEmailByUserIdQuery) (string, error) {
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

func (store *userStore) GetStatusByUserId(ctx context.Context, query *user.GetStatusByUserIdQuery) (int, error) {
	row := store.Conn.QueryRow(ctx, getUserByIdQuery, pgx.NamedArgs{"id": query.UserId})

	var status int
	if err := row.Scan(&status); err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return 0, err
	}

	return status, nil
}

/// GetPreferencesByUserId

const getPreferencesByUserIdQuery = ``

func (store *userStore) GetPreferencesByUserId(ctx context.Context, query *user.GetPreferencesByUserIdQuery) (any, error) {
	panic("not implemented")
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

func (store *userStore) GetSessionsByUserId(ctx context.Context, query *user.GetSessionsByUserIdQuery) ([]*user.UserSession, error) {
	rows, err := store.Conn.Query(ctx, getSessionsByUserIdQuery, pgx.NamedArgs{"userId": query.UserId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var sessionList []*user.UserSession = []*user.UserSession{}
	for rows.Next() {
		var session user.UserSession
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

func (store *userStore) GetAccountsByUserId(ctx context.Context, query *user.GetAccountsByUserIdQuery) ([]*user.UserAccount, error) {
	rows, err := store.Conn.Query(ctx, getAccountsByUserIdQuery, pgx.NamedArgs{"userId": query.UserId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var accountList []*user.UserAccount = []*user.UserAccount{}
	for rows.Next() {
		var acc user.UserAccount
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

	return accountList, nil
}

/// GetUsers

const getUsersQuery = ``

func (store *userStore) GetUsers(ctx context.Context, query *user.GetUsersQuery) ([]*user.User, error) {
	panic("not implemented")
}
