package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

type Api interface {
	// org
	GetOrganizationList(*gin.Context)
	GetOrganizationById(*gin.Context)
	// TODO: rewrite
	// GetOrganization(*gin.Context)
	CreateOrganization(*gin.Context)
	UpdateOrganizationById(*gin.Context)
	DeleteOrganizationById(*gin.Context)

	// project
	GetProjectList(*gin.Context)
}

type OrganizationApi struct {
	orgService org.Service
}

func NewOrganizationApi(service org.Service) *OrganizationApi {
	return &OrganizationApi{
		orgService: service,
	}
}
