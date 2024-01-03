package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/service/org"
)

// GET /orgs/:orgId
//
// Get organization by id.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *OrgApi) GetOrgById(ctx *gin.Context) {
	uriParams := &org.GetOrgByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := s.getOrgById(ctx, uriParams.OrgId)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs
//
// Get global organization list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *OrgApi) GetOrgList(ctx *gin.Context) {
	res := s.getOrgList(ctx)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgId/workspaces
//
// Get organization workspace list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *OrgApi) GetWorkspaceList(ctx *gin.Context) {
	uriParams := &org.GetWorkspaceListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := s.getWorkspaceList(ctx, uriParams.OrgId)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgId/projects
//
// Get organization project list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *OrgApi) GetProjectList(ctx *gin.Context) {
	uriParams := &org.GetProjectListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := s.getProjectList(ctx, uriParams.OrgId)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgId/members
//
// Get organization member list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *OrgApi) GetMemberList(ctx *gin.Context) {
	uriParams := &org.GetMemberListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := s.getMemberList(ctx, &org.GetMemberListQuery{OrgId: uriParams.OrgId})
	ctx.JSON(res.Status(), res.Body())
}
