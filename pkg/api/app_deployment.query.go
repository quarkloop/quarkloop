package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppDeploymentListResponse struct{}

func (s *ServerApi) GetAppDeploymentList(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	_ = orgId + workspaceId + projectId

	// query database

	res := &GetAppDeploymentListResponse{}
	c.JSON(http.StatusOK, res)
}

type GetAppDeploymentByIdResponse struct{}

func (s *ServerApi) GetAppDeploymentById(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	deploymentId := c.Param("deploymentId")
	_ = orgId + workspaceId + projectId + deploymentId

	// query database

	res := &GetAppDeploymentByIdResponse{}
	c.JSON(http.StatusOK, res)
}
