package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppDeploymentListResponse struct{}
type CreateAppDeploymentRequest struct{}
type CreateAppDeploymentResponse struct{}
type GetAppDeploymentByIdResponse struct{}
type UpdateAppDeploymentByIdRequest struct{}
type UpdateAppDeploymentByIdResponse struct{}

func (s *Server) GetAppDeploymentList(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &GetAppDeploymentListResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) CreateAppDeployment(c *gin.Context) {
	req := &CreateAppDeploymentRequest{}
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
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &CreateAppDeploymentResponse{}
	c.JSON(http.StatusCreated, res)
}

func (s *Server) GetAppDeploymentById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	deploymentId := c.Param("deploymentId")
	_ = osId + workspaceId + appId + deploymentId

	// query database

	res := &GetAppDeploymentByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) UpdateAppDeploymentById(c *gin.Context) {
	req := &UpdateAppDeploymentByIdRequest{}
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
	appId := c.Param("appId")
	deploymentId := c.Param("deploymentId")
	_ = osId + workspaceId + appId + deploymentId

	// query database

	res := &UpdateAppDeploymentByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteAppDeploymentById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	deploymentId := c.Param("deploymentId")
	_ = osId + workspaceId + appId + deploymentId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
