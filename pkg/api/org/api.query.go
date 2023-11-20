package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/organization"
)

type GetOrganizationListResponse struct {
	api.ApiResponse
	Data []model.Organization `json:"data"`
}

func (s *OrganizationApi) GetOrganizationList(c *gin.Context) {
	// query service
	orgList, err := s.orgService.GetOrganizationList(
		&organization.GetOrganizationListParams{
			Context: c,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetOrganizationListResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: orgList,
	}
	c.JSON(http.StatusOK, res)
}

type GetOrganizationByIdUriParams struct {
	OrgId string `uri:"orgId" binding:"required"`
}

type GetOrganizationByIdResponse struct {
	api.ApiResponse
	Data model.Organization `json:"data,omitempty"`
}

func (s *OrganizationApi) GetOrganizationById(c *gin.Context) {
	uriParams := &GetOrganizationByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	org, err := s.orgService.GetOrganizationById(
		&organization.GetOrganizationByIdParams{
			Context: c,
			Id:      uriParams.OrgId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetOrganizationByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *org,
	}
	c.JSON(http.StatusOK, res)
}

type GetOrganizationQueryParams struct {
	model.Organization
}

type GetOrganizationResponse struct {
	api.ApiResponse
	Data model.Organization `json:"data,omitempty"`
}

func (s *OrganizationApi) GetOrganization(c *gin.Context) {
	queryParams := &GetOrganizationQueryParams{}
	if err := c.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	org, err := s.orgService.GetOrganization(
		&organization.GetOrganizationParams{
			Context:      c,
			Organization: queryParams.Organization,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetOrganizationResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *org,
	}
	c.JSON(http.StatusOK, res)
}
