package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAppSubmissionListResponse struct{}
type CreateAppSubmissionRequest struct{}
type CreateAppSubmissionResponse struct{}
type GetAppSubmissionByIdResponse struct{}
type UpdateAppSubmissionByIdRequest struct{}
type UpdateAppSubmissionByIdResponse struct{}

func (s *Server) GetAppSubmissionList(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	_ = osId + workspaceId + appId

	// query database

	res := &GetAppSubmissionListResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) CreateAppSubmission(c *gin.Context) {
	req := &CreateAppSubmissionRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
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

func (s *Server) GetAppSubmissionById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	submissionId := c.Param("submissionId")
	_ = osId + workspaceId + appId + submissionId

	// query database

	res := &GetAppSubmissionByIdResponse{}
	c.JSON(http.StatusOK, res)
}

func (s *Server) UpdateAppSubmissionById(c *gin.Context) {
	req := &UpdateAppSubmissionByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AppResponsePayload{
			Status:       http.StatusBadRequest,
			StatusString: "BadRequest",
			Error:        err,
			ErrorString:  fmt.Sprintf("[BindJSON] %s", err.Error()),
		})
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

func (s *Server) DeleteAppSubmissionById(c *gin.Context) {
	osId := c.Param("osId")
	workspaceId := c.Param("workspaceId")
	appId := c.Param("appId")
	submissionId := c.Param("submissionId")
	_ = osId + workspaceId + appId + submissionId

	// query database

	c.JSON(http.StatusNoContent, nil)
}
