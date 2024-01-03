package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

// GET /orgs/:orgId/workspaces/:workspaceId
//
// Get workspace by id.
//
// Response status:
// 200: StatusOK
// 204: StatusNoContent
// 400: StatusBadRequest
// 404: StatusNotFound
// 500: StatusInternalServerError

func (s *WorkspaceApi) GetWorkspaceById(ctx *gin.Context) {
	uriParams := &workspace.GetWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	query := &workspace.GetWorkspaceByIdQuery{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
	}
	res := s.getWorkspaceById(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /workspaces
//
// Get global workspace list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *WorkspaceApi) GetWorkspaceList(ctx *gin.Context) {
	res := s.getWorkspaceList(ctx)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgId/workspaces/:workspaceId/projects
//
// Get workspace project list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *WorkspaceApi) GetProjectList(ctx *gin.Context) {
	uriParams := &workspace.GetProjectListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	query := &workspace.GetProjectListQuery{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
	}
	res := s.getProjectList(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}

// GET /orgs/:orgId/workspaces/:workspaceId/members
//
// Get workspace user list.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *WorkspaceApi) GetMemberList(ctx *gin.Context) {
	uriParams := &workspace.GetMemberListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	query := &workspace.GetMemberListQuery{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
	}
	res := s.getMemberList(ctx, query)
	ctx.JSON(res.Status(), res.Body())
}
