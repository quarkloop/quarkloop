package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

// POST /orgs/:orgId/workspaces
//
// Create workspace.
//
// Response status:
// 201: StatusCreated
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *WorkspaceApi) CreateWorkspace(ctx *gin.Context) {
	uriParams := &workspace.CreateWorkspaceUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cmd := &workspace.CreateWorkspaceCommand{OrgId: uriParams.OrgId}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := s.createWorkspace(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// PUT /orgs/:orgId/workspaces/:workspaceId
//
// Update workspace by id.
//
// Response status:
// 200: StatusOK
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *WorkspaceApi) UpdateWorkspaceById(ctx *gin.Context) {
	uriParams := &workspace.UpdateWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cmd := &workspace.UpdateWorkspaceByIdCommand{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
	}
	if err := ctx.ShouldBindJSON(cmd); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := s.updateWorkspaceById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}

// DELETE /orgs/:orgId/workspaces/:workspaceId
//
// Delete workspace by id.
//
// Response status:
// 204: StatusNoContent
// 400: StatusBadRequest
// 500: StatusInternalServerError

func (s *WorkspaceApi) DeleteWorkspaceById(ctx *gin.Context) {
	uriParams := &workspace.DeleteWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cmd := &workspace.DeleteWorkspaceByIdCommand{
		OrgId:       uriParams.OrgId,
		WorkspaceId: uriParams.WorkspaceId,
	}
	res := s.deleteWorkspaceById(ctx, cmd)
	ctx.JSON(res.Status(), res.Body())
}
