package org

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
)

type CreateOrganizationRequest struct {
	org.Organization
}

func (s *OrganizationApi) CreateOrganization(ctx *gin.Context) {
	req := &CreateOrganizationRequest{}
	if err := ctx.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query database
	org, err := s.orgService.CreateOrganization(ctx, &org.CreateOrganizationParams{
		Organization: req.Organization,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, org)
}

type UpdateOrganizationByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

type UpdateOrganizationByIdRequest struct {
	org.Organization
}

func (s *OrganizationApi) UpdateOrganizationById(ctx *gin.Context) {
	uriParams := &UpdateOrganizationByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	req := &UpdateOrganizationByIdRequest{}
	if err := ctx.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query database
	err := s.orgService.UpdateOrganizationById(ctx, &org.UpdateOrganizationByIdParams{
		OrgId:        uriParams.OrgId,
		Organization: req.Organization,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type DeleteOrganizationByIdUriParams struct {
	OrgId int `uri:"orgId" binding:"required"`
}

func (s *OrganizationApi) DeleteOrganizationById(ctx *gin.Context) {
	uriParams := &DeleteOrganizationByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query database
	err := s.orgService.DeleteOrganizationById(ctx, &org.DeleteOrganizationByIdParams{
		OrgId: uriParams.OrgId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
