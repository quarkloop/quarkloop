package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetWorkspaceListResponse struct{}
type CreateWorkspaceRequest struct{}
type CreateWorkspaceResponse struct{}
type GetWorkspaceByIdRequest struct{}
type GetWorkspaceByIdResponse struct{}
type UpdateWorkspaceByIdRequest struct{}
type UpdateWorkspaceByIdResponse struct{}
type DeleteWorkspaceByIdRequest struct{}

func (s *Server) GetWorkspaceList(c *gin.Context) {
	osId := c.Param("osId")
	_ = osId

	// query database

	res := &GetWorkspaceListResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) CreateWorkspace(c *gin.Context) {
	req := &CreateWorkspaceRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	osId := c.Param("osId")
	_ = osId

	// query database

	res := &CreateWorkspaceResponse{}
	c.JSON(http.StatusCreated, res)
}

func (s *Server) GetWorkspaceById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	_ = osId + workspaceId

	// query database

	res := &GetWorkspaceByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) UpdateWorkspaceById(c *gin.Context) {
	req := &UpdateWorkspaceByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
		return
	}

	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	_ = osId + workspaceId

	// query database

	res := &UpdateWorkspaceByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteWorkspaceById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	_ = osId + workspaceId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
