package accesscontrol

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

// GET /orgs/:orgId/accesscontrol/roles
//
// Get user role list.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *AccessControlApi) GetUserRoleList(ctx *gin.Context) {
	uriParams := &accesscontrol.GetUserRoleListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	orgList, err := s.aclService.GetUserRoleList(ctx, &accesscontrol.GetUserRoleListQuery{
		OrgId: uriParams.OrgId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &orgList)
}

// GET /orgs/:orgId/accesscontrol/roles/:roleId
//
// Get user role by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *AccessControlApi) GetUserRoleById(ctx *gin.Context) {
	uriParams := &accesscontrol.GetUserRoleByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.aclService.GetUserRoleById(ctx, &accesscontrol.GetUserRoleByIdQuery{
		OrgId:      uriParams.OrgId,
		UserRoleId: uriParams.UserRoleId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, org)
}
