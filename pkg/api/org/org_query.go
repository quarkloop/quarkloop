package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

// GET /orgs
//
// Get organization list.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *OrganizationApi) GetOrganizationList(ctx *gin.Context) {
	// query service
	orgList, err := s.orgService.GetOrganizationList(ctx)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &orgList)
}

// GET /orgs/:orgId
//
// Get organization by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *OrganizationApi) GetOrganizationById(ctx *gin.Context) {
	uriParams := &org.GetOrganizationByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.orgService.GetOrganizationById(ctx, &org.GetOrganizationByIdQuery{
		OrgId: uriParams.OrgId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, org)
}

// type GetOrganizationQueryParams struct {
// 	org.Organization
// }

// func (s *OrganizationApi) GetOrganization(ctx *gin.Context) {
// 	queryParams := &GetOrganizationQueryParams{}
// 	if err := ctx.ShouldBindQuery(queryParams); err != nil {
// 		api.AbortWithBadRequestJSON(ctx, err)
// 		return
// 	}

// 	// query service
// 	org, err := s.orgService.GetOrganization(ctx, &org.GetOrganizationQuery{
// 		Organization: queryParams.Organization,
// 	})
// 	if err != nil {
// 		api.AbortWithInternalServerErrorJSON(ctx, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, org)
// }

// GET /orgs/:orgId/projects
//
// Get organization project list.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *OrganizationApi) GetProjectList(ctx *gin.Context) {
	uriParams := &org.GetProjectListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	projectList, err := s.orgService.GetProjectList(ctx, &org.GetProjectListQuery{
		OrgId: uriParams.OrgId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &projectList)
}
