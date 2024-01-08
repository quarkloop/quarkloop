package user

import "context"

type Service interface {
	// query
	GetUser(context.Context, *GetUserQuery) (*User, error)
	GetUsername(context.Context, *GetUsernameQuery) (string, error)
	GetEmail(context.Context, *GetEmailQuery) (string, error)
	GetStatus(context.Context, *GetStatusQuery) (any, error)
	GetPreferences(context.Context, *GetPreferencesQuery) (any, error)
	GetSessions(context.Context, *GetSessionsQuery) ([]*Session, error)
	GetAccounts(context.Context, *GetAccountsQuery) ([]*Account, error)
	GetUserById(context.Context, *GetUserByIdQuery) (*User, error)
	GetUsernameByUserId(context.Context, *GetUsernameByUserIdQuery) (string, error)
	GetEmailByUserId(context.Context, *GetEmailByUserIdQuery) (string, error)
	GetStatusByUserId(context.Context, *GetStatusByUserIdQuery) (int, error)
	GetPreferencesByUserId(context.Context, *GetPreferencesByUserIdQuery) (any, error)
	GetSessionsByUserId(context.Context, *GetSessionsByUserIdQuery) ([]*Session, error)
	GetAccountsByUserId(context.Context, *GetAccountsByUserIdQuery) ([]*Account, error)
	GetUsers(context.Context, *GetUsersQuery) ([]*User, error)

	// mutation
	UpdateUser(context.Context, *UpdateUserCommand) error
	UpdateUsername(context.Context, *UpdateUsernameCommand) error
	UpdatePassword(context.Context, *UpdatePasswordCommand) error
	UpdatePreferences(context.Context, *UpdatePreferencesCommand) error
	UpdateUserById(context.Context, *UpdateUserByIdCommand) error
	UpdateUsernameByUserId(context.Context, *UpdateUsernameByUserIdCommand) error
	UpdatePasswordByUserId(context.Context, *UpdatePasswordByUserIdCommand) error
	UpdatePreferencesByUserId(context.Context, *UpdatePreferencesByUserIdCommand) error
	DeleteUserById(context.Context, *DeleteUserByIdCommand) error
	DeleteSessionById(context.Context, *DeleteSessionByIdCommand) error
	DeleteAccountById(context.Context, *DeleteAccountByIdCommand) error
}
