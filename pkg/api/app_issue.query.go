package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppIssueListResponse struct{}

func (s *ServerApi) GetAppIssueList(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	_ = orgId + workspaceId + projectId

	// query database

	res := &GetAppIssueListResponse{}
	c.JSON(http.StatusOK, res)
}

type GetAppIssueByIdResponse struct{}

func (s *ServerApi) GetAppIssueById(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	issueId := c.Param("issueId")
	_ = orgId + workspaceId + projectId + issueId

	// query database

	res := &GetAppIssueByIdResponse{}
	c.JSON(http.StatusOK, res)
}
