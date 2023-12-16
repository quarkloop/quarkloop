package org

import (
	"github.com/gin-gonic/gin"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
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
	orgService org.Service
}

func NewOrganizationApi(service org.Service) *OrganizationApi {
	return &OrganizationApi{
		orgService: service,
	}
}
