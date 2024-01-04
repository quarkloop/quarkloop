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
	Password      string     `json:"password,omitempty"`
	Birthdate     *time.Time `json:"birthdate,omitempty"`
	Country       string     `json:"country,omitempty"`
	Image         string     `json:"image,omitempty"`
	Status        int        `json:"status,omitempty"`
	Path          string     `json:"path,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
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
	Type         string     `json:"type,omitempty"`
	TokenType    string     `json:"tokenType,omitempty"`
	Provider     string     `json:"provider,omitempty"`
	RefreshToken string     `json:"refereshToken,omitempty"`
	AccessToken  string     `json:"accessToken,omitempty"`
	Scope        string     `json:"scope,omitempty"`
	SessionState string     `json:"sessionState,omitempty"`
	ExpiresAt    *time.Time `json:"expiresAt,omitempty"`
}

type UserSession struct {
	// id
	Id int `json:"id,omitempty"`

	// session
	SessionToken string `json:"sessionToken,omitempty"`
	ExpiresAt    string `json:"expires,omitempty"`
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
	UserId int `uri:"userId" binding:"required"`
}

type GetUserByIdQuery struct {
	UserId int
}

// GetUsernameByUserId

type GetUsernameByUserIdQuery struct{}

// GetEmailByUserId

type GetEmailByUserIdQuery struct{}

// GetStatusByUserId

type GetStatusByUserIdQuery struct{}

// GetPreferencesByUserId

type GetPreferencesByUserIdQuery struct{}

// GetSessionsByUserId

type GetSessionsByUserIdQuery struct{}

// GetAccountsByUserId

type GetAccountsByUserIdQuery struct{}

// GetUsers

type GetUsersQuery struct{}

////////////////////

// UpdateUserById

type UpdateUserByIdUriParams struct {
	UserId int `uri:"userId" binding:"required"`
}

type UpdateUserByIdCommand struct {
	UserId int
	User
}

// DeleteUserById

type DeleteUserByIdUriParams struct {
	UserId int `uri:"userId" binding:"required"`
}

type DeleteUserByIdCommand struct {
	UserId int
}

/////////////////////////

// UpdateUser

type UpdateUserCommand struct{}

// UpdateUsername

type UpdateUsernameCommand struct{}

// UpdatePassword

type UpdatePasswordCommand struct{}

// UpdatePreferences

type UpdatePreferencesCommand struct{}

// UpdateUserById

type UpdateUserByIdCommand struct{}

// UpdateUsernameByUserId

type UpdateUsernameByUserIdCommand struct{}

// UpdatePasswordByUserId

type UpdatePasswordByUserIdCommand struct{}

// UpdatePreferencesByUserId

type UpdatePreferencesByUserIdCommand struct{}

//  DeleteUserById

type DeleteUserByIdCommand struct{}

// DeleteSessionById

type DeleteSessionByIdCommand struct{}

// DeleteAccountById

type DeleteAccountByIdCommand struct{}
