package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type CreateWorkspaceRequest struct {
	OrgId     int                 `json:"orgId" binding:"required"`
	Workspace workspace.Workspace `json:"workspace" binding:"required"`
}

func (s *WorkspaceApi) CreateWorkspace(ctx *gin.Context) {
	req := &CreateWorkspaceRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	ws, err := s.workspaceService.CreateWorkspace(ctx, &workspace.CreateWorkspaceParams{
		OrgId:     req.OrgId,
		Workspace: req.Workspace,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, ws)
}

type UpdateWorkspaceByIdUriParams struct {
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

type UpdateWorkspaceByIdRequest struct {
	workspace.Workspace
}

func (s *WorkspaceApi) UpdateWorkspaceById(ctx *gin.Context) {
	uriParams := &UpdateWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	req := &UpdateWorkspaceByIdRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.workspaceService.UpdateWorkspaceById(ctx, &workspace.UpdateWorkspaceByIdParams{
		WorkspaceId: uriParams.WorkspaceId,
		Workspace:   req.Workspace,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type DeleteWorkspaceByIdUriParams struct {
	WorkspaceId int `uri:"workspaceId" binding:"required"`
}

func (s *WorkspaceApi) DeleteWorkspaceById(ctx *gin.Context) {
	uriParams := &DeleteWorkspaceByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.workspaceService.DeleteWorkspaceById(ctx, &workspace.DeleteWorkspaceByIdParams{
		WorkspaceId: uriParams.WorkspaceId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
