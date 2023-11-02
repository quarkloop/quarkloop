package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAppSubmissionRequest struct{}
type CreateAppSubmissionResponse struct{}

func (s *ServerApi) CreateAppSubmission(c *gin.Context) {
	req := &CreateAppSubmissionRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &CreateAppSubmissionResponse{}
	c.JSON(http.StatusCreated, res)
}

type UpdateAppSubmissionByIdRequest struct{}
type UpdateAppSubmissionByIdResponse struct{}

func (s *ServerApi) UpdateAppSubmissionById(c *gin.Context) {
	req := &UpdateAppSubmissionByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		AbortWithBadRequestJSON(c, err)
		return
	}

	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	submissionId := c.Param("submissionId")
	_ = osId + workspaceId + appId + submissionId

	// query database

	res := &UpdateAppSubmissionByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *ServerApi) DeleteAppSubmissionById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	submissionId := c.Param("submissionId")
	_ = osId + workspaceId + appId + submissionId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
