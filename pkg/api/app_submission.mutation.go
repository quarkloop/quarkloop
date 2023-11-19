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

	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	_ = orgId + workspaceId + projectId

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

	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	submissionId := c.Param("submissionId")
	_ = orgId + workspaceId + projectId + submissionId

	// query database

	res := &UpdateAppSubmissionByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *ServerApi) DeleteAppSubmissionById(c *gin.Context) {
	orgId := c.Param("orgId")
	workspaceId := c.Param("workspaceId")
	projectId := c.Param("projectId")
	submissionId := c.Param("submissionId")
	_ = orgId + workspaceId + projectId + submissionId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
