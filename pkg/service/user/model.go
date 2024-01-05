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
	Name          string     `json:"name,omitempty"`
	Email         string     `json:"email,omitempty"`
	EmailVerified *time.Time `json:"emailVerified,omitempty"`
	Password      *string    `json:"password,omitempty"`
	Birthdate     *time.Time `json:"birthdate,omitempty"`
	Country       *string    `json:"country,omitempty"`
	Image         string     `json:"image,omitempty"`
	Status        int        `json:"status,omitempty"`
	Path          string     `json:"path,omitempty"`

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

type UserAccount struct {
	// id
	Id                int    `json:"id,omitempty"`
	TokenId           string `json:"idToken,omitempty"`
	ProviderAccountId string `json:"providerAccountId,omitempty"`

	// account
	Type         string    `json:"type,omitempty"`
	TokenType    string    `json:"tokenType,omitempty"`
	Provider     string    `json:"provider,omitempty"`
	RefreshToken *string   `json:"refereshToken,omitempty"`
	AccessToken  string    `json:"accessToken,omitempty"`
	Scope        string    `json:"scope,omitempty"`
	SessionState *string   `json:"sessionState,omitempty"`
	ExpiresAt    time.Time `json:"expiresAt,omitempty"`
}

type UserSession struct {
	// id
	Id int `json:"id,omitempty"`

	// session
	SessionToken string    `json:"sessionToken,omitempty"`
	ExpiresAt    time.Time `json:"expires,omitempty"`
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

//////////////////////////////////////////////////

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

	Name          string    `json:"name,omitempty"`
	Email         string    `json:"email,omitempty"`
	EmailVerified time.Time `json:"emailVerified,omitempty"`
	Birthdate     time.Time `json:"birthdate,omitempty"`
	Country       string    `json:"country,omitempty"`
	Image         string    `json:"image,omitempty"`
	Status        int       `json:"status,omitempty"`
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

// DeleteSessionById

type DeleteSessionByIdUriParams struct {
	UserId    int `uri:"userId" binding:"required"`
	SessionId int `uri:"sessionId" binding:"required"`
}

type DeleteSessionByIdCommand struct {
	UserId    int
	SessionId int
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
