package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/organization"
)

type CreateOrganizationRequest struct {
	model.Organization
}

type CreateOrganizationResponse struct {
	api.ApiResponse
	Data model.Organization `json:"data,omitempty"`
}

func (s *OrganizationApi) CreateOrganization(c *gin.Context) {
	req := &CreateOrganizationRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query database
	org, err := s.orgService.CreateOrganization(
		&organization.CreateOrganizationParams{
			Context:      c,
			Organization: req.Organization,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateOrganizationResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *org,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateOrganizationByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type UpdateOrganizationByIdRequest struct {
	model.Organization
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
		&organization.UpdateOrganizationByIdParams{
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
		&organization.DeleteOrganizationByIdParams{
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
