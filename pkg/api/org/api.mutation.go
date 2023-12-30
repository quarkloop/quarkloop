package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
)

// POST /orgs
//
// Create organization.
//
// Response status:
// 201: StatusCreated
// 500: StatusInternalServerError

func (s *OrganizationApi) CreateOrganization(ctx *gin.Context) {
	cmd := &org.CreateOrganizationCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query database
	org, err := s.orgService.CreateOrganization(ctx, cmd)
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

func (s *OrganizationApi) UpdateOrganizationById(ctx *gin.Context) {
	uriParams := &org.UpdateOrganizationByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &org.UpdateOrganizationByIdCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query database
	err := s.orgService.UpdateOrganizationById(ctx, &org.UpdateOrganizationByIdCommand{
		OrgId:        uriParams.OrgId,
		Organization: cmd.Organization,
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

func (s *OrganizationApi) DeleteOrganizationById(ctx *gin.Context) {
	uriParams := &org.DeleteOrganizationByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query database
	err := s.orgService.DeleteOrganizationById(ctx, &org.DeleteOrganizationByIdCommand{
		OrgId: uriParams.OrgId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
