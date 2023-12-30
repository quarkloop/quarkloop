package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

// GET /workspaces
//
// Get workspace list.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *WorkspaceApi) GetWorkspaceList(ctx *gin.Context) {
	queryParams := &workspace.GetWorkspaceListQueryParams{}
	if err := ctx.ShouldBindQuery(queryParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	wsList, err := s.workspaceService.GetWorkspaceList(ctx, &workspace.GetWorkspaceListQuery{
		OrgId: queryParams.OrgId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &wsList)
}

// GET /workspaces/:workspaceId
//
// Get workspace by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *WorkspaceApi) GetWorkspaceById(ctx *gin.Context) {
	uriParams := &workspace.GetWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	ws, err := s.workspaceService.GetWorkspaceById(ctx, &workspace.GetWorkspaceByIdQuery{
		WorkspaceId: uriParams.WorkspaceId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ws)
}

// type GetWorkspaceQueryParams struct {
// 	OrgId int `form:"orgId" binding:"required"`
// 	workspace.Workspace
// }

// func (s *WorkspaceApi) GetWorkspace(ctx *gin.Context) {
// 	queryParams := &GetWorkspaceQueryParams{}
// 	if err := ctx.ShouldBindQuery(queryParams); err != nil {
// 		api.AbortWithBadRequestJSON(ctx, err)
// 		return
// 	}

// 	// query service
// 	ws, err := s.workspaceService.GetWorkspace(ctx, &workspace.GetWorkspaceQuery{
// 		OrgId:     queryParams.OrgId,
// 		Workspace: queryParams.Workspace,
// 	})
// 	if err != nil {
// 		api.AbortWithInternalServerErrorJSON(ctx, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, ws)
// }

// GET /orgs/:orgId/workspaces/:workspaceId/projects
//
// Get workspace project list.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *WorkspaceApi) GetProjectList(ctx *gin.Context) {
	uriParams := &workspace.GetProjectListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	projectList, err := s.workspaceService.GetProjectList(ctx, &workspace.GetProjectListQuery{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &projectList)
}
