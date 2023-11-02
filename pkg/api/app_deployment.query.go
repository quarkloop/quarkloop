package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppDeploymentListResponse struct{}

func (s *ServerApi) GetAppDeploymentList(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &GetAppDeploymentListResponse{}
	c.JSON(http.StatusOK, res)
}

type GetAppDeploymentByIdResponse struct{}

func (s *ServerApi) GetAppDeploymentById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	deploymentId := c.Param("deploymentId")
	_ = osId + workspaceId + appId + deploymentId

	// query database

	res := &GetAppDeploymentByIdResponse{}
	c.JSON(http.StatusOK, res)
}
