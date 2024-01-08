package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type UserStore interface {
	// query
	GetUser(context.Context, *user.GetUserQuery) (*user.User, error)
	GetUsername(context.Context, *user.GetUsernameQuery) (string, error)
	GetEmail(context.Context, *user.GetEmailQuery) (string, error)
	GetStatus(context.Context, *user.GetStatusQuery) (any, error)
	GetPreferences(context.Context, *user.GetPreferencesQuery) (any, error)
	GetSessions(context.Context, *user.GetSessionsQuery) ([]*user.Session, error)
	GetAccounts(context.Context, *user.GetAccountsQuery) ([]*user.Account, error)
	GetUserById(context.Context, *user.GetUserByIdQuery) (*user.User, error)
	GetUsernameByUserId(context.Context, *user.GetUsernameByUserIdQuery) (string, error)
	GetEmailByUserId(context.Context, *user.GetEmailByUserIdQuery) (string, error)
	GetStatusByUserId(context.Context, *user.GetStatusByUserIdQuery) (int, error)
	GetPreferencesByUserId(context.Context, *user.GetPreferencesByUserIdQuery) (any, error)
	GetSessionsByUserId(context.Context, *user.GetSessionsByUserIdQuery) ([]*user.Session, error)
	GetAccountsByUserId(context.Context, *user.GetAccountsByUserIdQuery) ([]*user.Account, error)
	GetUsers(context.Context, *user.GetUsersQuery) ([]*user.User, error)

	// mutation
	UpdateUser(context.Context, *user.UpdateUserCommand) error
	UpdateUsername(context.Context, *user.UpdateUsernameCommand) error
	UpdatePassword(context.Context, *user.UpdatePasswordCommand) error
	UpdatePreferences(context.Context, *user.UpdatePreferencesCommand) error
	UpdateUserById(context.Context, *user.UpdateUserByIdCommand) error
	UpdateUsernameByUserId(context.Context, *user.UpdateUsernameByUserIdCommand) error
	UpdatePasswordByUserId(context.Context, *user.UpdatePasswordByUserIdCommand) error
	UpdatePreferencesByUserId(context.Context, *user.UpdatePreferencesByUserIdCommand) error
	DeleteUserById(context.Context, *user.DeleteUserByIdCommand) error
	DeleteSessionById(context.Context, *user.DeleteSessionByIdCommand) error
	DeleteAccountById(context.Context, *user.DeleteAccountByIdCommand) error
}

type userStore struct {
	Conn *pgx.Conn
}

func NewUserStore(conn *pgx.Conn) *userStore {
	return &userStore{
		Conn: conn,
	}
}
