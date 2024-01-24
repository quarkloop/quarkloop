package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

// GET /orgs/:orgId
//
// Get organization by id.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 404: StatusNotFound
// 500: StatusInternalServerError

func (s *orgApi) GetOrgById(ctx *gin.Context) {
	query := &org.GetOrgByIdQuery{}
	if err := ctx.ShouldBindUri(query); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.getOrgById(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs
//
// Get global organization list.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *orgApi) GetOrgList(ctx *gin.Context) {
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

func (s *orgApi) GetWorkspaceList(ctx *gin.Context) {
	query := &org.GetWorkspaceListQuery{}
	if err := ctx.ShouldBindUri(query); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.getWorkspaceList(ctx, query)
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

func (s *orgApi) GetProjectList(ctx *gin.Context) {
	query := &org.GetProjectListQuery{}
	if err := ctx.ShouldBindUri(query); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.getProjectList(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgId/members
//
// Get organization member list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 404: StatusNotFound
// 500: StatusInternalServerError

func (s *orgApi) GetMemberList(ctx *gin.Context) {
	query := &org.GetMemberListQuery{}
	if err := ctx.ShouldBindUri(query); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.getMemberList(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}
