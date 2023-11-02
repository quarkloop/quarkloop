package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/db/model"
	"github.com/quarkloop/quarkloop/pkg/db/repository"
)

type CreateWorkspaceRequest struct {
	model.Workspace
}

type CreateWorkspaceResponse struct {
	ApiResponse
	Data model.Workspace `json:"data,omitempty"`
}

func (s *ServerApi) CreateWorkspace(c *gin.Context) {
	req := &CreateWorkspaceRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	osId := c.Param("osId")

	// query database
	ws, err := s.dataStore.CreateWorkspace(
		&repository.CreateWorkspaceParams{
			Context:   c,
			OsId:      osId,
			Workspace: req.Workspace,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateWorkspaceResponse{
		ApiResponse: ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateWorkspaceByIdRequest struct {
	model.Workspace
}

func (s *ServerApi) UpdateWorkspaceById(c *gin.Context) {
	req := &UpdateWorkspaceByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")

	// query database
	err := s.dataStore.UpdateWorkspaceById(
		&repository.UpdateWorkspaceByIdParams{
			Context:     c,
			OsId:        osId,
			WorkspaceId: workspaceId,
			Workspace:   req.Workspace,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (s *ServerApi) DeleteWorkspaceById(c *gin.Context) {
	//osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")

	// query database
	err := s.dataStore.DeleteWorkspaceById(
		&repository.DeleteWorkspaceByIdParams{
			Context: c,
			Id:      workspaceId,
		},
	)
	if err != nil {
		AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
