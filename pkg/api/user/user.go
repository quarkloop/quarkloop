package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type Api interface {
	// query
	GetUser(*gin.Context)
	GetUsername(*gin.Context)
	GetEmail(*gin.Context)
	GetStatus(*gin.Context)
	GetPreferences(*gin.Context)
	GetSessions(*gin.Context)
	GetAccounts(*gin.Context)
	GetUserById(*gin.Context)
	GetUsernameByUserId(*gin.Context)
	GetEmailByUserId(*gin.Context)
	GetStatusByUserId(*gin.Context)
	GetPreferencesByUserId(*gin.Context)
	GetSessionsByUserId(*gin.Context)
	GetAccountsByUserId(*gin.Context)
	GetUsers(*gin.Context)

	// mutation
	UpdateUser(*gin.Context)
	UpdateUsername(*gin.Context)
	UpdatePassword(*gin.Context)
	UpdatePreferences(*gin.Context)
	UpdateUserById(*gin.Context)
	UpdateUsernameByUserId(*gin.Context)
	UpdatePasswordByUserId(*gin.Context)
	UpdatePreferencesByUserId(*gin.Context)
	DeleteUserById(*gin.Context)
	DeleteSessionById(*gin.Context)
	DeleteAccountById(*gin.Context)
}

type userApi struct {
	userService user.Service
	aclService  accesscontrol.Service
}

func NewUserApi(service user.Service, aclService accesscontrol.Service) *userApi {
	return &userApi{
		userService: service,
		aclService:  aclService,
	}
}
