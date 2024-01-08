package user

import (
	"fmt"
	"time"
)

type User struct {
	// id
	Id       int    `json:"id,string"` // string: type bigint from javascript is encoded into string
	Username string `json:"username"`

	// user
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	EmailVerified *time.Time `json:"emailVerified"`
	Password      *string    `json:"password"`
	Birthdate     *time.Time `json:"birthdate"`
	Country       *string    `json:"country"`
	Image         string     `json:"image"`
	Status        int        `json:"status"`
	Path          string     `json:"path"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy *string    `json:"createdBy"` // just for user type we use pointer string, cuz we cannot fill it while creation
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

func (u *User) GetId() int {
	return u.Id
}

func (u *User) GeneratePath() {
	u.Path = fmt.Sprintf("/users/%d", u.GetId())
}

type Account struct {
	// id
	Id                int    `json:"id"`
	UserId            int    `json:"userId"`
	TokenId           string `json:"tokenId"`
	ProviderAccountId string `json:"providerAccountId"`

	// account
	Type         string    `json:"type"`
	TokenType    string    `json:"tokenType"`
	Provider     string    `json:"provider"`
	RefreshToken *string   `json:"refereshToken"`
	AccessToken  string    `json:"accessToken"`
	Scope        string    `json:"scope"`
	SessionState *string   `json:"sessionState"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

type Session struct {
	// id
	Id     int
	UserId int

	// session
	SessionToken string    `json:"sessionToken"`
	ExpiresAt    time.Time `json:"expires"`
}

type UserAssignment struct {
	// id
	Id     int `json:"id"`
	UserId int `json:"userId"`

	// user assignment
	Role string `json:"role"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type MemberDTO struct {
	// user
	User *User `json:"user"`

	// user assignment
	Role string `json:"role"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

// GetUser

type GetUserQuery struct{}

// GetUsername

type GetUsernameQuery struct{}

// GetEmail

type GetEmailQuery struct{}

// GetStatus

type GetStatusQuery struct{}

// GetPreferences

type GetPreferencesQuery struct{}

// GetSessions

type GetSessionsQuery struct{}

// GetAccounts

type GetAccountsQuery struct{}

// GetUserById

type GetUserByIdUriParams struct {
	UserId int `uri:"userId_or_username" binding:"required"`
}

type GetUserByIdQuery struct {
	UserId int
}

// GetUsernameByUserId

type GetUsernameByUserIdUriParams struct {
	UserId int `uri:"userId_or_username" binding:"required"`
}

type GetUsernameByUserIdQuery struct {
	UserId int
}

// GetEmailByUserId

type GetEmailByUserIdUriParams struct {
	UserId int `uri:"userId_or_username" binding:"required"`
}

type GetEmailByUserIdQuery struct {
	UserId int
}

// GetStatusByUserId

type GetStatusByUserIdUriParams struct {
	UserId int `uri:"userId_or_username" binding:"required"`
}

type GetStatusByUserIdQuery struct {
	UserId int
}

// GetPreferencesByUserId

type GetPreferencesByUserIdUriParams struct {
	UserId int `uri:"userId_or_username" binding:"required"`
}

type GetPreferencesByUserIdQuery struct {
	UserId int
}

// GetSessionsByUserId

type GetSessionsByUserIdUriParams struct {
	UserId int `uri:"userId_or_username" binding:"required"`
}

type GetSessionsByUserIdQuery struct {
	UserId int
}

// GetAccountsByUserId

type GetAccountsByUserIdUriParams struct {
	UserId int `uri:"userId_or_username" binding:"required"`
}

type GetAccountsByUserIdQuery struct {
	UserId int
}

// GetUsers

type GetUsersQuery struct{}

/******************************* mutations*******************************/

// CreateUser

type CreateUserCommand struct {
	CreatedBy string

	Username  string     `json:"username"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Birthdate *time.Time `json:"birthdate"`
	Country   string     `json:"country"`
	Image     string     `json:"image"`
	Status    int        `json:"status"`
}

// UpdateUser

type UpdateUserCommand struct{}

// UpdateUsername

type UpdateUsernameCommand struct{}

// UpdatePassword

type UpdatePasswordCommand struct{}

// UpdatePreferences

type UpdatePreferencesCommand struct{}

// UpdateUserById

type UpdateUserByIdUriParams struct {
	UserId int `uri:"userId" binding:"required"`
}

type UpdateUserByIdCommand struct {
	UserId    int
	UpdatedBy string

	Username      string     `json:"username"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	EmailVerified *time.Time `json:"emailVerified"`
	Birthdate     *time.Time `json:"birthdate"`
	Country       string     `json:"country"`
	Image         string     `json:"image"`
	Status        int        `json:"status"`
}

// UpdateUsernameByUserId

type UpdateUsernameByUserIdUriParams struct {
	UserId int `uri:"userId" binding:"required"`
}

type UpdateUsernameByUserIdCommand struct {
	UserId    int
	UpdatedBy string
	Username  string `json:"username" binding:"required"`
}

// UpdatePasswordByUserId

type UpdatePasswordByUserIdUriParams struct {
	UserId int `uri:"userId" binding:"required"`
}

type UpdatePasswordByUserIdCommand struct {
	UserId    int
	UpdatedBy string
	Password  string `json:"password" binding:"required"`
}

// UpdatePreferencesByUserId

type UpdatePreferencesByUserIdUriParams struct {
	UserId int `uri:"userId" binding:"required"`
}

type UpdatePreferencesByUserIdCommand struct {
	UserId int
}

//  DeleteUserById

type DeleteUserByIdUriParams struct {
	UserId int `uri:"userId" binding:"required"`
}

type DeleteUserByIdCommand struct {
	UserId int
}

// CreateSession

type CreateSessionCommand struct {
	UserId int

	SessionToken string    `json:"sessionToken"`
	ExpiresAt    time.Time `json:"expires"`
}

// DeleteSessionById

type DeleteSessionByIdUriParams struct {
	UserId    int `uri:"userId" binding:"required"`
	SessionId int `uri:"sessionId" binding:"required"`
}

type DeleteSessionByIdCommand struct {
	UserId    int
	SessionId int
}

// CreateAccountCommand

type CreateAccountCommand struct {
	UserId            int
	TokenId           string `json:"tokenId"`
	ProviderAccountId string `json:"providerAccountId"`

	Type         string  `json:"type"`
	TokenType    string  `json:"tokenType"`
	Provider     string  `json:"provider"`
	RefreshToken *string `json:"refereshToken"`
	AccessToken  string  `json:"accessToken"`
}

// DeleteAccountById

type DeleteAccountByIdUriParams struct {
	UserId    int `uri:"userId" binding:"required"`
	AccountId int `uri:"accountId" binding:"required"`
}

type DeleteAccountByIdCommand struct {
	UserId    int
	AccountId int
}
