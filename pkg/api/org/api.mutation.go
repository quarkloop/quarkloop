package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
)

type CreateOrganizationRequest struct {
	org.Organization
}

func (s *OrganizationApi) CreateOrganization(c *gin.Context) {
	req := &CreateOrganizationRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	org, err := s.orgService.CreateOrganization(
		&org.CreateOrganizationParams{
			Context:      c,
			Organization: req.Organization,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusCreated, org)
}

type UpdateOrganizationByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type UpdateOrganizationByIdRequest struct {
	org.Organization
}

func (s *OrganizationApi) UpdateOrganizationById(c *gin.Context) {
	uriParams := &UpdateOrganizationByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateOrganizationByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	err := s.orgService.UpdateOrganizationById(
		&org.UpdateOrganizationByIdParams{
			Context:      c,
			OrgId:        uriParams.OrgId,
			Organization: req.Organization,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteOrganizationByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

func (s *OrganizationApi) DeleteOrganizationById(c *gin.Context) {
	uriParams := &DeleteOrganizationByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	err := s.orgService.DeleteOrganizationById(
		&org.DeleteOrganizationByIdParams{
			Context: c,
			OrgId:   uriParams.OrgId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
