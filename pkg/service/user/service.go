package user

import "context"

type Service interface {
	// user
	GetUserById(context.Context, *GetUserByIdQuery) (*User, error)
	GetUserByEmail(context.Context, *GetUserByEmailQuery) (*User, error)
	UpdateUserById(context.Context, *UpdateUserByIdCommand) error
	DeleteUserById(context.Context, *DeleteUserByIdCommand) error

	// account
	GetUserAccountByUserId(context.Context, *GetUserAccountByUserIdQuery) (*UserAccount, error)

	// session
	GetUserSessionByUserId(context.Context, *GetUserSessionByUserIdQuery) (*UserSession, error)
}
