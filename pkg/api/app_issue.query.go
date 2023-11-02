package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppIssueListResponse struct{}

func (s *ServerApi) GetAppIssueList(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &GetAppIssueListResponse{}
	c.JSON(http.StatusOK, res)
}

type GetAppIssueByIdResponse struct{}

func (s *ServerApi) GetAppIssueById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	issueId := c.Param("issueId")
	_ = osId + workspaceId + appId + issueId

	// query database

	res := &GetAppIssueByIdResponse{}
	c.JSON(http.StatusOK, res)
}
