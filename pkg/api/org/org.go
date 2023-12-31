package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

type Api interface {
	// org
	GetOrgList(*gin.Context)
	GetOrgById(*gin.Context)
	// TODO: rewrite
	// GetOrg(*gin.Context)
	CreateOrg(*gin.Context)
	UpdateOrgById(*gin.Context)
	DeleteOrgById(*gin.Context)

	// project
	GetProjectList(*gin.Context)

	// user
	GetUserList(*gin.Context)
}

type OrgApi struct {
	orgService org.Service
}

func NewOrgApi(service org.Service) *OrgApi {
	return &OrgApi{
		orgService: service,
	}
}
