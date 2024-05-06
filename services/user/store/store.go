package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quarkloop/quarkloop/pkg/model"
)

type UserStore interface {
	// query
	GetUserById(context.Context, *GetUserByIdQuery) (*model.User, error)
	GetUsernameByUserId(context.Context, *GetUsernameByUserIdQuery) (string, error)
	GetEmailByUserId(context.Context, *GetEmailByUserIdQuery) (string, error)
	GetStatusByUserId(context.Context, *GetStatusByUserIdQuery) (int, error)
	GetSessionsByUserId(context.Context, *GetSessionsByUserIdQuery) ([]*model.Session, error)
	GetAccountsByUserId(context.Context, *GetAccountsByUserIdQuery) ([]*model.Account, error)
	GetUsers(context.Context, *GetUsersQuery) ([]*model.User, error)

	// user mutation
	CreateUser(context.Context, *CreateUserCommand) (*model.User, error)
	UpdateUserById(context.Context, *UpdateUserByIdCommand) error
	UpdateUsernameByUserId(context.Context, *UpdateUsernameByUserIdCommand) error
	UpdatePasswordByUserId(context.Context, *UpdatePasswordByUserIdCommand) error
	DeleteUserById(context.Context, *DeleteUserByIdCommand) error

	// session mutation
	CreateSession(context.Context, *CreateSessionCommand) (*model.Session, error)
	DeleteSessionById(context.Context, *DeleteSessionByIdCommand) error

	// account mutation
	CreateAccount(context.Context, *CreateAccountCommand) (*model.Account, error)
	DeleteAccountById(context.Context, *DeleteAccountByIdCommand) error
}

type userStore struct {
	Conn *pgxpool.Pool
}

func NewUserStore(conn *pgxpool.Pool) UserStore {
	return &userStore{
		Conn: conn,
	}
}
