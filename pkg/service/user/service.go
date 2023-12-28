package user

import "context"

type Service interface {
	// user
	GetUserById(context.Context, *GetUserByIdParams) (*User, error)
	GetUserByEmail(context.Context, *GetUserByEmailParams) (*User, error)
	UpdateUserById(context.Context, *UpdateUserByIdParams) error
	DeleteUserById(context.Context, *DeleteUserByIdParams) error

	// account
	GetUserAccountByUserId(context.Context, *GetUserAccountByUserIdParams) (*UserAccount, error)

	// session
	GetUserSessionByUserId(context.Context, *GetUserSessionByUserIdParams) (*UserSession, error)
}
