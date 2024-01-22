package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/user/store"
	"github.com/quarkloop/quarkloop/pkg/test"
)

var (
	ctx       context.Context
	conn      *pgx.Conn
	sessionId int32
	accountId int32
)

func init() {
	ctx, conn = test.InitTestAuthDB()
}

func TestMutationTruncateTables(t *testing.T) {
	t.Run("truncate tables", func(t *testing.T) {
		err := test.TruncateAuthDBTables(ctx, conn)
		require.NoError(t, err)
	})
}

func TestMutationOnCreateUser(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("create user", func(t *testing.T) {
		cmd := &user.CreateUserCommand{
			Username:  "username",
			Name:      "FirstName LastName",
			Email:     "email@gmail.com",
			Image:     "https://lh3.googleusercontent.com/a/A",
			Status:    1,
			CreatedBy: "admin",
		}
		u, err := store.CreateUser(ctx, cmd)
		require.NoError(t, err)
		require.NotNil(t, u)

		{
			// check the update
			query := &user.GetUserByIdQuery{UserId: u.Id}
			user, err := store.GetUserById(ctx, query)

			require.NoError(t, err)
			require.NotNil(t, user)
			require.NotEmpty(t, user.Username)
			require.NotEmpty(t, user.Name)
			require.NotEmpty(t, user.Email)
			require.Nil(t, user.Birthdate)
			require.Empty(t, *user.Country)
			require.NotEmpty(t, user.Image)
			require.NotEmpty(t, user.CreatedBy)
			require.Nil(t, user.UpdatedBy)
			require.NotZero(t, user.Status)
			require.Equal(t, cmd.Username, user.Username)
			require.Equal(t, cmd.Name, user.Name)
			require.Equal(t, cmd.Email, user.Email)
			require.Equal(t, cmd.Country, *user.Country)
			require.Equal(t, cmd.Image, user.Image)
		}
	})

	t.Run("update user", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for idx, u := range userList {
			birthdate := time.Now().UTC()
			cmd := &user.UpdateUserByIdCommand{
				UserId:    u.Id,
				Username:  fmt.Sprintf("%s #%d", u.Username, idx),
				Name:      fmt.Sprintf("%s #%d", u.Name, idx),
				Email:     fmt.Sprintf("%s #%d", u.Email, idx),
				Country:   fmt.Sprintf("%s #%d", *u.Country, idx),
				Image:     fmt.Sprintf("%s #%d", u.Image, idx),
				UpdatedBy: fmt.Sprintf("%s #%d", u.Image, idx),
				Birthdate: &birthdate,
				Status:    1,
			}
			err := store.UpdateUserById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &user.GetUserByIdQuery{UserId: u.Id}
				user, err := store.GetUserById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, user)
				require.NotEmpty(t, user.Username)
				require.NotEmpty(t, user.Name)
				require.NotEmpty(t, user.Email)
				require.NotNil(t, user.Birthdate)
				require.NotEmpty(t, user.Country)
				require.NotEmpty(t, user.Image)
				require.NotEmpty(t, user.CreatedBy)
				require.NotEmpty(t, user.UpdatedBy)
				require.NotZero(t, user.Status)
				require.Equal(t, cmd.Username, user.Username)
				require.Equal(t, cmd.Name, user.Name)
				require.Equal(t, cmd.Email, user.Email)
				require.WithinDuration(t, *cmd.Birthdate, *user.Birthdate, time.Millisecond)
				require.Equal(t, cmd.Country, *user.Country)
				require.Equal(t, cmd.Image, user.Image)
			}
		}
	})
}

func TestQueryOnUser(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("get user by wrong id", func(t *testing.T) {
		u, err := store.GetUserById(ctx, &user.GetUserByIdQuery{UserId: 99999})

		require.Error(t, err)
		require.Nil(t, u)
	})
}

func TestMutationOnCreateAccount(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("create account", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			cmd := &user.CreateAccountCommand{
				UserId:            u.Id,
				Type:              "oauth",
				Provider:          "google",
				ProviderAccountId: "1012545206303",
				TokenType:         "Bearer",
				Scope:             "https://www.googleapis.com/auth/userinfo",
				AccessToken:       "ya29.a0AfB_byBFphZcWM404QM7E2d3QswTIagqYchbY6kq5ffc1k9mtG5hYvQWLos",
				TokenId:           "eyJhbGciOiJSUzI1NiIsImtpZCI6IjkxNDEzY2Y0ZmEwY2I5MmEzYzNmNWEwNTQKV1Q",
			}
			acc, err := store.CreateAccount(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, acc)

			accountId = acc.Id
		}
	})
}

func TestMutationOnCreateSession(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("create session", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			ti, err := time.Parse(time.RFC3339, "2024-02-07T12:49:28.58Z")
			require.NoError(t, err)
			cmd := &user.CreateSessionCommand{
				UserId:       u.Id,
				ExpiresAt:    ti,
				SessionToken: "d74b5d8-8262-40da-9997-c609d1281a",
			}
			sess, err := store.CreateSession(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, sess)

			sessionId = sess.Id
		}
	})
}

func TestQueryOnSession(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("list sessions", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			list, err := store.GetSessionsByUserId(ctx, &user.GetSessionsByUserIdQuery{UserId: u.Id})

			require.NoError(t, err)
			require.NotEmpty(t, list)
		}
	})

	t.Run("get session by wrong id", func(t *testing.T) {
		list, err := store.GetSessionsByUserId(ctx, &user.GetSessionsByUserIdQuery{UserId: 99999})

		require.NoError(t, err)
		require.Empty(t, list)
	})
}

func TestQueryOnAccount(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("list accounts", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			list, err := store.GetAccountsByUserId(ctx, &user.GetAccountsByUserIdQuery{UserId: u.Id})

			require.NoError(t, err)
			require.NotEmpty(t, list)
		}
	})

	t.Run("get account by wrong id", func(t *testing.T) {
		list, err := store.GetAccountsByUserId(ctx, &user.GetAccountsByUserIdQuery{UserId: 99999})

		require.NoError(t, err)
		require.Empty(t, list)
	})
}

func TestMutationOnDeleteSession(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("delete session by wrong id", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			cmd := &user.DeleteSessionByIdCommand{UserId: u.Id, SessionId: 99999}
			err := store.DeleteSessionById(ctx, cmd)

			require.Error(t, err)
		}
	})

	t.Run("delete session by id", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			cmd := &user.DeleteSessionByIdCommand{UserId: u.Id, SessionId: sessionId}
			err := store.DeleteSessionById(ctx, cmd)

			require.NoError(t, err)
		}
	})
}

func TestMutationOnDeleteAccount(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("delete account by wrong id", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			cmd := &user.DeleteAccountByIdCommand{UserId: u.Id, AccountId: 99999}
			err := store.DeleteAccountById(ctx, cmd)

			require.Error(t, err)
		}
	})

	t.Run("delete account by id", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			cmd := &user.DeleteAccountByIdCommand{UserId: u.Id, AccountId: accountId}
			err := store.DeleteAccountById(ctx, cmd)

			require.NoError(t, err)
		}
	})
}

func TestMutationOnDeleteUser(t *testing.T) {
	store := store.NewUserStore(conn)

	t.Run("delete user by id", func(t *testing.T) {
		userList, err := test.GetFullUserList(ctx, conn)
		require.NoError(t, err)
		require.NotEmpty(t, userList)

		for _, u := range userList {
			cmd := &user.DeleteUserByIdCommand{UserId: u.Id}
			err := store.DeleteUserById(ctx, cmd)

			require.NoError(t, err)
		}
	})
}
