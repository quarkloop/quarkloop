package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
)

func (s *OrganizationApi) GetOrganizationList(c *gin.Context) {
	// query service
	orgList, err := s.orgService.GetOrganizationList(
		&org.GetOrganizationListParams{
			Context: c,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, &orgList)
}

type GetOrganizationByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

func (s *OrganizationApi) GetOrganizationById(c *gin.Context) {
	uriParams := &GetOrganizationByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	org, err := s.orgService.GetOrganizationById(
		&org.GetOrganizationByIdParams{
			Context: c,
			OrgId:   uriParams.OrgId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, org)
}

type GetOrganizationQueryParams struct {
	org.Organization
}

func (s *OrganizationApi) GetOrganization(c *gin.Context) {
	queryParams := &GetOrganizationQueryParams{}
	if err := c.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	org, err := s.orgService.GetOrganization(
		&org.GetOrganizationParams{
			Context:      c,
			Organization: queryParams.Organization,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, org)
}
