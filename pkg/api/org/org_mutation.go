package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/service/org"
)

// POST /orgs
//
// Create organization.
//
// Response status:
// 201: StatusCreated
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *orgApi) CreateOrg(ctx *gin.Context) {
	cmd := &org.CreateOrgCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
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
	uriParams := &org.UpdateOrgByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	cmd := &org.UpdateOrgByIdCommand{OrgId: uriParams.OrgId}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
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
	uriParams := &org.DeleteOrgByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res := s.deleteOrgById(ctx, uriParams.OrgId)
	ctx.JSON(res.Status(), res.Body())
}
