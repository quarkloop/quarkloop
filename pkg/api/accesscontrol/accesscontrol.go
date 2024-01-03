package accesscontrol

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

type Api interface {
	// UserGroup
	GetUserGroupList(*gin.Context)
	GetUserGroupById(*gin.Context)
	CreateUserGroup(*gin.Context)
	UpdateUserGroupById(*gin.Context)
	DeleteUserGroupById(*gin.Context)
}

type AccessControlApi struct {
	aclService accesscontrol.Service
}

func NewAccessControlApi(service accesscontrol.Service) *AccessControlApi {
	return &AccessControlApi{
		aclService: service,
	}
}
