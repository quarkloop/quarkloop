package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
)

func (s *OrganizationApi) GetOrganizationList(ctx *gin.Context) {
	// query service
	orgList, err := s.orgService.GetOrganizationList(ctx, &org.GetOrganizationListParams{})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &orgList)
}

type GetOrganizationByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

func (s *OrganizationApi) GetOrganizationById(ctx *gin.Context) {
	uriParams := &GetOrganizationByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.orgService.GetOrganizationById(ctx, &org.GetOrganizationByIdParams{
		OrgId: uriParams.OrgId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, org)
}

type GetOrganizationQueryParams struct {
	org.Organization
}

func (s *OrganizationApi) GetOrganization(ctx *gin.Context) {
	queryParams := &GetOrganizationQueryParams{}
	if err := ctx.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.orgService.GetOrganization(ctx, &org.GetOrganizationParams{
		Organization: queryParams.Organization,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, org)
}
