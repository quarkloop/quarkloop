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

func (s *OrgApi) GetOrgList(ctx *gin.Context) {
	// query service
	orgList, err := s.orgService.GetOrgList(ctx)
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

func (s *OrgApi) GetOrgById(ctx *gin.Context) {
	uriParams := &org.GetOrgByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.orgService.GetOrgById(ctx, &org.GetOrgByIdQuery{
		OrgId: uriParams.OrgId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, org)
}

// type GetOrgQueryParams struct {
// 	org.Org
// }

// func (s *OrgApi) GetOrg(ctx *gin.Context) {
// 	queryParams := &GetOrgQueryParams{}
// 	if err := ctx.ShouldBindQuery(queryParams); err != nil {
// 		api.AbortWithBadRequestJSON(ctx, err)
// 		return
// 	}

// 	// query service
// 	org, err := s.orgService.GetOrg(ctx, &org.GetOrgQuery{
// 		Org: queryParams.Org,
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

func (s *OrgApi) GetProjectList(ctx *gin.Context) {
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
