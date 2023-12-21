package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type OrgStore interface {
	// user
	GetUserById(ctx context.Context, userId int) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	UpdateUserById(ctx context.Context, userId int, user *user.User) error
	DeleteUserById(ctx context.Context, userId int) error

	// account
	GetUserAccountByUserId(ctx context.Context, userId int) (*user.UserAccount, error)

	// session
	GetUserSessionByUserId(ctx context.Context, userId int) (*user.UserSession, error)
}

type orgStore struct {
	Conn *pgx.Conn
}

func NewOrgStore(conn *pgx.Conn) *orgStore {
	return &orgStore{
		Conn: conn,
	}
}
