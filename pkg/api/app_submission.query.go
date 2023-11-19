package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppSubmissionListResponse struct{}

func (s *ServerApi) GetAppSubmissionList(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	_ = orgId + workspaceId + projectId

	// query database

	res := &GetAppSubmissionListResponse{}
	c.JSON(http.StatusOK, res)
}

type GetAppSubmissionByIdResponse struct{}

func (s *ServerApi) GetAppSubmissionById(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	submissionId := c.Param("submissionId")
	_ = orgId + workspaceId + projectId + submissionId

	// query database

	res := &GetAppSubmissionByIdResponse{}
	c.JSON(http.StatusOK, res)
}
