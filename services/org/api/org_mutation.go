package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/grpc/v1/system/org"
	"github.com/quarkloop/quarkloop/pkg/model"
)

// POST /orgs
//
// Create organization.
//
// Response status:
// 201: StatusCreated
// 400: StatusBadRequest
// 403: StatusForbidden
// 500: StatusInternalServerError

type CreateOrgCommand struct {
	Payload *model.OrgMutation `json:"payload" binding:"required"`
}

type CreateOrgDTO struct {
	Data *model.Org `json:"data" binding:"required"`
}

func (s *orgApi) CreateOrg(ctx *gin.Context) {
	cmd := &CreateOrgCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.createOrg(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /orgs/:orgSid
//
// Update organization by id.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

type UpdateOrgByIdUriParams struct {
	OrgSid string `uri:"orgSid" binding:"required,sid"`
}

type UpdateOrgByIdCommand struct {
	OrgId   int64              `json:"-"`
	Payload *model.OrgMutation `json:"payload" binding:"required"`
}

func (s *orgApi) UpdateOrgById(ctx *gin.Context) {
	uriParams := &UpdateOrgByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	reply, err := s.orgService.GetOrgId(ctx, &org.GetOrgIdQuery{OrgSid: uriParams.OrgSid})
	if err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	cmd := &UpdateOrgByIdCommand{OrgId: reply.OrgId}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.updateOrgById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// DELETE /orgs/:orgSid
//
// Delete organization by id.
//
// Response status:
// 204: StatusNoContent
// 400: StatusBadRequest
// 500: StatusInternalServerError

type DeleteOrgByIdUriParams struct {
	OrgSid string `uri:"orgSid" binding:"required,sid"`
}

type DeleteOrgByIdCommand struct {
	OrgId int64
}

func (s *orgApi) DeleteOrgById(ctx *gin.Context) {
	uriParams := &DeleteOrgByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	reply, err := s.orgService.GetOrgId(ctx, &org.GetOrgIdQuery{OrgSid: uriParams.OrgSid})
	if err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.deleteOrgById(ctx, &DeleteOrgByIdCommand{OrgId: reply.OrgId})
	ctx.JSON(res.Status(), res.Body())
}

// PUT /manage/:orgSid/visibility
//
// Change organization visibility.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

type ChangeOrgVisibilityUriParams struct {
	OrgSid string `uri:"orgSid" binding:"required,sid"`
}

type ChangeOrgVisibilityCommand struct {
	OrgId      int64                 `json:"-"`
	Visibility model.ScopeVisibility `json:"visibility" binding:"required"`
}

func (s *orgApi) ChangeOrgVisibility(ctx *gin.Context) {
	uriParams := &ChangeOrgVisibilityUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	reply, err := s.orgService.GetOrgId(ctx, &org.GetOrgIdQuery{OrgSid: uriParams.OrgSid})
	if err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	cmd := &ChangeOrgVisibilityCommand{OrgId: reply.OrgId}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.changeOrgVisibility(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}
