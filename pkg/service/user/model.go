package user

import (
	"fmt"
	"time"
)

type User struct {
	// id
	// string: type bigint from javascript is encoded into string
	Id int `json:"id,string,omitempty"`

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
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
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

// GetUserById

type GetUserByIdQuery struct {
	UserId int
}

// GetUserByEmail

type GetUserByEmailQuery struct {
	Email string
}

// UpdateUserById

type UpdateUserByIdCommand struct {
	UserId int
	User
}

// DeleteUserById

type DeleteUserByIdCommand struct {
	UserId int
}

// GetUserAccountByUserId

type GetUserAccountByUserIdQuery struct {
	UserId int
}

//  GetUserSessionByUserId

type GetUserSessionByUserIdQuery struct {
	UserId int
}
