package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppSubmissionListResponse struct{}

func (s *ServerApi) GetAppSubmissionList(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &GetAppSubmissionListResponse{}
	c.JSON(http.StatusOK, res)
}

type GetAppSubmissionByIdResponse struct{}

func (s *ServerApi) GetAppSubmissionById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	submissionId := c.Param("submissionId")
	_ = osId + workspaceId + appId + submissionId

	// query database

	res := &GetAppSubmissionByIdResponse{}
	c.JSON(http.StatusOK, res)
}
