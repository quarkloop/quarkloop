package user

import (
	"time"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type UserAssignment struct {
	// id
	Id     int64 `json:"id"`
	UserId int64 `json:"userId"`

	// user assignment
	Role string `json:"role"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

type MemberDTO struct {
	User *model.User `json:"user"`
	Role string      `json:"role"`
}

type GetUserByIdUriParams struct {
	UserId int64 `uri:"userId_or_username" binding:"required"`
}

type GetUsernameByUserIdUriParams struct {
	UserId int64 `uri:"userId_or_username" binding:"required"`
}

type GetEmailByUserIdUriParams struct {
	UserId int64 `uri:"userId_or_username" binding:"required"`
}

type GetStatusByUserIdUriParams struct {
	UserId int64 `uri:"userId_or_username" binding:"required"`
}

type GetPreferencesByUserIdUriParams struct {
	UserId int64 `uri:"userId_or_username" binding:"required"`
}

type GetSessionsByUserIdUriParams struct {
	UserId int64 `uri:"userId_or_username" binding:"required"`
}

type GetAccountsByUserIdUriParams struct {
	UserId int64 `uri:"userId_or_username" binding:"required"`
}

/******************************* mutations*******************************/

type UpdateUserByIdUriParams struct {
	UserId int64 `uri:"userId" binding:"required"`
}

type UpdateUsernameByUserIdUriParams struct {
	UserId int64 `uri:"userId" binding:"required"`
}

type UpdatePasswordByUserIdUriParams struct {
	UserId int64 `uri:"userId" binding:"required"`
}

type UpdatePreferencesByUserIdUriParams struct {
	UserId int64 `uri:"userId" binding:"required"`
}

type DeleteUserByIdUriParams struct {
	UserId int64 `uri:"userId" binding:"required"`
}

type DeleteSessionByIdUriParams struct {
	UserId    int64 `uri:"userId" binding:"required"`
	SessionId int64 `uri:"sessionId" binding:"required"`
}

type CreateAccountCommand struct {
	UserId            int64
	TokenId           string `json:"tokenId"`
	ProviderAccountId string `json:"providerAccountId"`

	Type         string  `json:"type"`
	TokenType    string  `json:"tokenType"`
	Provider     string  `json:"provider"`
	Scope        string  `json:"scope"`
	RefreshToken *string `json:"refereshToken"`
	AccessToken  string  `json:"accessToken"`
}

type DeleteAccountByIdUriParams struct {
	UserId    int64 `uri:"userId" binding:"required"`
	AccountId int64 `uri:"accountId" binding:"required"`
}
