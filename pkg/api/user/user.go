package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type Api interface {
	GetUserById(*gin.Context)
	CreateUser(*gin.Context)
	UpdateUserById(*gin.Context)
	DeleteUserById(*gin.Context)
}

type UserApi struct {
	userService user.Service
	aclService  accesscontrol.Service
}

func NewUserApi(service user.Service, aclService accesscontrol.Service) *UserApi {
	return &UserApi{
		userService: service,
		aclService:  aclService,
	}
}
