package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

type Api interface {
	// query
	GetOrgById(*gin.Context)
	GetOrgList(*gin.Context)
	GetWorkspaceList(*gin.Context)
	GetProjectList(*gin.Context)
	GetUserList(*gin.Context)

	// mutation
	CreateOrg(*gin.Context)
	UpdateOrgById(*gin.Context)
	DeleteOrgById(*gin.Context)
}

type OrgApi struct {
	orgService org.Service
	aclService accesscontrol.Service
}

func NewOrgApi(service org.Service, aclService accesscontrol.Service) *OrgApi {
	return &OrgApi{
		orgService: service,
		aclService: aclService,
	}
}
