package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/user"
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

type OrgApi struct {
	orgService org.Service

	userService  user.Service
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewOrgApi(service org.Service, userService user.Service, aclService accesscontrol.Service, quotaService quota.Service) *OrgApi {
	return &OrgApi{
		orgService:   service,
		userService:  userService,
		aclService:   aclService,
		quotaService: quotaService,
	}
}
