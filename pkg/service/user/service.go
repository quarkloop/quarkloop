package user

import "context"

type Service interface {
	// user
	GetUserById(context.Context, *GetUserByIdParams) (*User, error)
	GetUserByEmail(context.Context, *GetUserByIdParams) (*User, error)
	UpdateUserById(context.Context, *UpdateUserByIdParams) error
	DeleteUserById(context.Context, *DeleteUserByIdParams) error

	// account
	GetUserAccountByUserId(context.Context, *GetUserByIdParams) (*UserAccount, error)

	// session
	GetUserSessionByUserId(context.Context, *GetUserByIdParams) (*UserSession, error)
}
