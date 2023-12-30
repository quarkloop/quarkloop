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
// 500: StatusInternalServerError

func (s *OrgApi) CreateOrg(ctx *gin.Context) {
	cmd := &org.CreateOrgCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.orgService.CreateOrg(ctx, cmd)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, org)
}

// PUT /orgs/:orgId
//
// Update organization by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *OrgApi) UpdateOrgById(ctx *gin.Context) {
	uriParams := &org.UpdateOrgByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &org.UpdateOrgByIdCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.orgService.UpdateOrgById(ctx, &org.UpdateOrgByIdCommand{
		OrgId: uriParams.OrgId,
		Org:   cmd.Org,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// DELETE /orgs/:orgId
//
// Delete organization by id.
//
// Response status:
// 204: StatusNoContent
// 500: StatusInternalServerError

func (s *OrgApi) DeleteOrgById(ctx *gin.Context) {
	uriParams := &org.DeleteOrgByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.orgService.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{
		OrgId: uriParams.OrgId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
