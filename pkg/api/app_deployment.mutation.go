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

	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	_ = orgId + workspaceId + projectId

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

	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	deploymentId := c.Param("deploymentId")
	_ = orgId + workspaceId + projectId + deploymentId

	// query database

	res := &UpdateAppDeploymentByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *ServerApi) DeleteAppDeploymentById(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	deploymentId := c.Param("deploymentId")
	_ = orgId + workspaceId + projectId + deploymentId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
