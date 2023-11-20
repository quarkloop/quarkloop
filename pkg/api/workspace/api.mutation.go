package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type CreateWorkspaceRequest struct {
	OrgId     string          `json:"orgId" binding:"required"`
	Workspace model.Workspace `json:"workspace" binding:"required"`
}

type CreateWorkspaceResponse struct {
	api.ApiResponse
	Data model.Workspace `json:"data,omitempty"`
}

func (s *WorkspaceApi) CreateWorkspace(c *gin.Context) {
	req := &CreateWorkspaceRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.workspaceService.CreateWorkspace(
		&workspace.CreateWorkspaceParams{
			Context:   c,
			OrgId:     req.OrgId,
			Workspace: req.Workspace,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateWorkspaceResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateWorkspaceByIdUriParams struct {
	WorkspaceId string `uri:"workspaceId" binding:"required"`
}

type UpdateWorkspaceByIdRequest struct {
	model.Workspace
}

func (s *WorkspaceApi) UpdateWorkspaceById(c *gin.Context) {
	uriParams := &UpdateWorkspaceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateWorkspaceByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.workspaceService.UpdateWorkspaceById(
		&workspace.UpdateWorkspaceByIdParams{
			Context:     c,
			WorkspaceId: uriParams.WorkspaceId,
			Workspace:   req.Workspace,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteWorkspaceByIdUriParams struct {
	WorkspaceId string `uri:"workspaceId" binding:"required"`
}

func (s *WorkspaceApi) DeleteWorkspaceById(c *gin.Context) {
	uriParams := &DeleteWorkspaceByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.workspaceService.DeleteWorkspaceById(
		&workspace.DeleteWorkspaceByIdParams{
			Context:     c,
			WorkspaceId: uriParams.WorkspaceId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
