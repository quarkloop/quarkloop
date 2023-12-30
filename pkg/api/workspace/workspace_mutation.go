package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

// POST /orgs/:orgId/workspaces
//
// Create workspace.
//
// Response status:
// 201: StatusCreated
// 500: StatusInternalServerError

func (s *WorkspaceApi) CreateWorkspace(ctx *gin.Context) {
	cmd := &workspace.CreateWorkspaceCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	ws, err := s.workspaceService.CreateWorkspace(ctx, cmd)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, ws)
}

// PUT /orgs/:orgId/workspaces/:workspaceId
//
// Update workspace by id.
//
// Response status:
// 200: StatusOK
// 500: StatusInternalServerError

func (s *WorkspaceApi) UpdateWorkspaceById(ctx *gin.Context) {
	uriParams := &workspace.UpdateWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	cmd := &workspace.UpdateWorkspaceByIdCommand{}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.workspaceService.UpdateWorkspaceById(ctx, &workspace.UpdateWorkspaceByIdCommand{
		WorkspaceId: uriParams.WorkspaceId,
		Workspace:   cmd.Workspace,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// DELETE /orgs/:orgId/workspaces/:workspaceId
//
// Delete workspace by id.
//
// Response status:
// 204: StatusNoContent
// 500: StatusInternalServerError

func (s *WorkspaceApi) DeleteWorkspaceById(ctx *gin.Context) {
	uriParams := &workspace.DeleteWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.workspaceService.DeleteWorkspaceById(ctx, &workspace.DeleteWorkspaceByIdCommand{
		WorkspaceId: uriParams.WorkspaceId,
	})
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
