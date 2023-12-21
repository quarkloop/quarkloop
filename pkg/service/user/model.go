package user

import (
	"fmt"
	"time"
)

type User struct {
	// id
	Id int `json:"id" form:"id"`

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

func (o *User) GeneratePath() {
	o.Path = fmt.Sprintf("/users/%d", o.Id)
}

type UserAccount struct {
	// id
	Id                int    `json:"id" form:"id"`
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
	Id int `json:"id" form:"id"`

	// saession
	SessionToken string `json:"sessionToken,omitempty"`
	ExpiresAt    string `json:"expires,omitempty"`
}

type GetUserByIdParams struct {
	OrgId int
}

type UpdateUserByIdParams struct {
	OrgId int
	User  User
}

type DeleteUserByIdParams struct {
	OrgId int
}

type GetUserAccountByIdParams struct {
	OrgId int
}

type GetUserSessionByIdParams struct {
	OrgId int
}
