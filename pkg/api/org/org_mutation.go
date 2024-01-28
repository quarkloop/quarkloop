package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/org"
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

func (s *orgApi) CreateOrg(ctx *gin.Context) {
	cmd := &org.CreateOrgCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.createOrg(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /orgs/:orgId
//
// Update organization by id.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *orgApi) UpdateOrgById(ctx *gin.Context) {
	cmd := &org.UpdateOrgByIdCommand{}

	if err := ctx.ShouldBindUri(cmd); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.updateOrgById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// DELETE /orgs/:orgId
//
// Delete organization by id.
//
// Response status:
// 204: StatusNoContent
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *orgApi) DeleteOrgById(ctx *gin.Context) {
	cmd := &org.DeleteOrgByIdCommand{}
	if err := ctx.ShouldBindUri(cmd); err != nil {
		api.AbortWithStatusJSON(ctx, http.StatusBadRequest, err)
		return
	}

	res := s.deleteOrgById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}
