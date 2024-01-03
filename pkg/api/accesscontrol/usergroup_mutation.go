package accesscontrol

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

// POST /orgs/:orgId/accesscontrol/groups
//
// Create user group.
//
// Response status:
// 201: StatusCreated
// 500: StatusInternalServerError

func (s *AccessControlApi) CreateUserGroup(ctx *gin.Context) {
	uriParams := &accesscontrol.CreateUserGroupUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &accesscontrol.CreateUserGroupCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.aclService.CreateUserGroup(ctx, &accesscontrol.CreateUserGroupCommand{
		OrgId:     uriParams.OrgId,
		UserGroup: cmd.UserGroup,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, org)
}

// PUT /orgs/:orgId/accesscontrol/groups/:groupId
//
// Update user group by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *AccessControlApi) UpdateUserGroupById(ctx *gin.Context) {
	uriParams := &accesscontrol.UpdateUserGroupByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &accesscontrol.UpdateUserGroupByIdCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.aclService.UpdateUserGroupById(ctx, &accesscontrol.UpdateUserGroupByIdCommand{
		OrgId:       uriParams.OrgId,
		UserGroupId: uriParams.UserGroupId,
		UserGroup:   cmd.UserGroup,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// DELETE /orgs/:orgId/accesscontrol/groups/:groupId
//
// Delete user group by id.
//
// Response status:
// 204: StatusNoContent
// 500: StatusInternalServerError

func (s *AccessControlApi) DeleteUserGroupById(ctx *gin.Context) {
	uriParams := &accesscontrol.DeleteUserGroupByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.aclService.DeleteUserGroupById(ctx, &accesscontrol.DeleteUserGroupByIdCommand{
		OrgId:       uriParams.OrgId,
		UserGroupId: uriParams.UserGroupId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
