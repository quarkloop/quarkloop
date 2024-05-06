package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
)

// GET /orgs/:orgSid
//
// Get organization by id.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 404: StatusNotFound
// 500: StatusInternalServerError

type GetOrgByIdQuery struct {
	OrgSid string `uri:"orgSid" binding:"required,sid"`
	OrgId  int64  `uri:"-"`
}

type GetOrgByIdDTO struct {
	Data *model.Org `json:"data" binding:"required"`
}

func (s *orgApi) GetOrgById(ctx *gin.Context) {
	query := &GetOrgByIdQuery{}
	if err := ctx.ShouldBindUri(query); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	orgId, err, errStatus := GetOrgId(ctx, s.orgService, query.OrgSid)
	if err != nil {
		api.AbortWithStatusJSON(ctx, errStatus, err)
		return
	}
	query.OrgId = orgId

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

type GetOrgListDTO struct {
	Data []*model.Org `json:"data" binding:"required"`
}

func (s *orgApi) GetOrgList(ctx *gin.Context) {
	res := s.getOrgList(ctx)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgSid/workspaces
//
// Get organization workspace list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

type GetWorkspaceListQuery struct {
	OrgSid     string                `uri:"orgSid" binding:"required,sid"`
	OrgId      int64                 `uri:"-"`
	Visibility model.ScopeVisibility `uri:"-"`
}

type GetWorkspaceListDTO struct {
	Data []*model.Workspace `json:"data" binding:"required"`
}

func (s *orgApi) GetWorkspaceList(ctx *gin.Context) {
	query := &GetWorkspaceListQuery{}
	if err := ctx.ShouldBindUri(query); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	orgId, err, errStatus := GetOrgId(ctx, s.orgService, query.OrgSid)
	if err != nil {
		api.AbortWithStatusJSON(ctx, errStatus, err)
		return
	}
	query.OrgId = orgId

	res := s.getWorkspaceList(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgSid/members
//
// Get organization member list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 404: StatusNotFound
// 500: StatusInternalServerError

type GetMemberListQuery struct {
	OrgSid string `uri:"orgSid" binding:"required,sid"`
	OrgId  int64  `uri:"-"`
}

func (s *orgApi) GetMemberList(ctx *gin.Context) {
	query := &GetMemberListQuery{}
	if err := ctx.ShouldBindUri(query); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	orgId, err, errStatus := GetOrgId(ctx, s.orgService, query.OrgSid)
	if err != nil {
		api.AbortWithStatusJSON(ctx, errStatus, err)
		return
	}
	query.OrgId = orgId

	res := s.getMemberList(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}
