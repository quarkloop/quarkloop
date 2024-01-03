package user

import "context"

type Service interface {
	// query
	GetUserById(context.Context, *GetUserByIdQuery) (*User, error)
	GetUserByEmail(context.Context, *GetUserByEmailQuery) (*User, error)
	GetUserAccountByUserId(context.Context, *GetUserAccountByUserIdQuery) (*UserAccount, error)
	GetUserSessionByUserId(context.Context, *GetUserSessionByUserIdQuery) (*UserSession, error)

	// mutation
	UpdateUserById(context.Context, *UpdateUserByIdCommand) error
	DeleteUserById(context.Context, *DeleteUserByIdCommand) error
}
