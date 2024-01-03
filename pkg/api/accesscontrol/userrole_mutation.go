package accesscontrol

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

// POST /orgs/:orgId/accesscontrol/roles
//
// Create user role.
//
// Response status:
// 201: StatusCreated
// 500: StatusInternalServerError

func (s *AccessControlApi) CreateUserRole(ctx *gin.Context) {
	uriParams := &accesscontrol.CreateUserRoleUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &accesscontrol.CreateUserRoleCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.aclService.CreateUserRole(ctx, &accesscontrol.CreateUserRoleCommand{
		OrgId:    uriParams.OrgId,
		UserRole: cmd.UserRole,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, org)
}

// PUT /orgs/:orgId/accesscontrol/roles/:roleId
//
// Update user role by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *AccessControlApi) UpdateUserRoleById(ctx *gin.Context) {
	uriParams := &accesscontrol.UpdateUserRoleByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &accesscontrol.UpdateUserRoleByIdCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.aclService.UpdateUserRoleById(ctx, &accesscontrol.UpdateUserRoleByIdCommand{
		OrgId:      uriParams.OrgId,
		UserRoleId: uriParams.UserRoleId,
		UserRole:   cmd.UserRole,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// DELETE /orgs/:orgId/accesscontrol/roles/:roleId
//
// Delete user role by id.
//
// Response status:
// 204: StatusNoContent
// 500: StatusInternalServerError

func (s *AccessControlApi) DeleteUserRoleById(ctx *gin.Context) {
	uriParams := &accesscontrol.DeleteUserRoleByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.aclService.DeleteUserRoleById(ctx, &accesscontrol.DeleteUserRoleByIdCommand{
		OrgId:      uriParams.OrgId,
		UserRoleId: uriParams.UserRoleId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
