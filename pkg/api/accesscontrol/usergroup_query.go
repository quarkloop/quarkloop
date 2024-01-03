package accesscontrol

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

// GET /orgs/:orgId/accesscontrol/groups
//
// Get user group list.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *AccessControlApi) GetUserGroupList(ctx *gin.Context) {
	uriParams := &accesscontrol.GetUserGroupListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	orgList, err := s.aclService.GetUserGroupList(ctx, &accesscontrol.GetUserGroupListQuery{
		OrgId: uriParams.OrgId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &orgList)
}

// GET /orgs/:orgId/accesscontrol/groups/:groupId
//
// Get user group by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *AccessControlApi) GetUserGroupById(ctx *gin.Context) {
	uriParams := &accesscontrol.GetUserGroupByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	org, err := s.aclService.GetUserGroupById(ctx, &accesscontrol.GetUserGroupByIdQuery{
		OrgId:       uriParams.OrgId,
		UserGroupId: uriParams.UserGroupId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, org)
}
