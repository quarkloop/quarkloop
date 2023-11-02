package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAppDeploymentRequest struct{}
type CreateAppDeploymentResponse struct{}

func (s *ServerApi) CreateAppDeployment(c *gin.Context) {
	req := &CreateAppDeploymentRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
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

type UpdateAppDeploymentByIdRequest struct{}
type UpdateAppDeploymentByIdResponse struct{}

func (s *ServerApi) UpdateAppDeploymentById(c *gin.Context) {
	req := &UpdateAppDeploymentByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
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

func (s *ServerApi) DeleteAppDeploymentById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	deploymentId := c.Param("deploymentId")
	_ = osId + workspaceId + appId + deploymentId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
