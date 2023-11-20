package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/organization"
)

type Api interface {
	GetOrganizationList(c *gin.Context)
	GetOrganizationById(c *gin.Context)
	GetOrganization(c *gin.Context)
	CreateOrganization(c *gin.Context)
	UpdateOrganizationById(c *gin.Context)
	DeleteOrganizationById(c *gin.Context)
}

type OrganizationApi struct {
	orgService organization.Service
}

func NewOrganizationApi(service organization.Service) *OrganizationApi {
	return &OrganizationApi{
		orgService: service,
	}
}
