package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/service/system"
)

type Api interface {
	// query
	GetOrgById(*gin.Context)
	GetOrgList(*gin.Context)
	GetWorkspaceList(*gin.Context)
	GetProjectList(*gin.Context)
	GetMemberList(*gin.Context)

	// mutation
	CreateOrg(*gin.Context)
	UpdateOrgById(*gin.Context)
	DeleteOrgById(*gin.Context)
}

type orgApi struct {
	orgService system.OrgServiceClient

	userService  user.Service
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewOrgApi(orgService system.OrgServiceClient, userService user.Service, aclService accesscontrol.Service, quotaService quota.Service) *orgApi {
	return &orgApi{
		orgService:   orgService,
		userService:  userService,
		aclService:   aclService,
		quotaService: quotaService,
	}
}

// func (api *orgApi) GetService() org.Service {
// 	return api.orgService
// }
