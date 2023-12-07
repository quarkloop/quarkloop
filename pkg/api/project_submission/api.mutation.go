package project_submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type CreateAppSubmissionUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type CreateAppSubmissionRequest struct {
	// UserId     string                  `json:"userId" binding:"required"`
	// Submission model.AppSubmission `json:"submission" binding:"required"`
	model.AppSubmission
}

type CreateAppSubmissionResponse struct {
	api.ApiResponse
	Data model.AppSubmission `json:"data,omitempty"`
}

func (s *AppSubmissionApi) CreateAppSubmission(c *gin.Context) {
	uriParams := &CreateAppSubmissionUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &CreateAppSubmissionRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.projectSubmission.CreateAppSubmission(
		&project_submission.CreateAppSubmissionParams{
			Context:       c,
			UserId:        "req.UserId",
			ProjectId:     uriParams.ProjectId,
			AppSubmission: &req.AppSubmission,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateAppSubmissionResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateAppSubmissionByIdUriParams struct {
	SubmissionId string `uri:"submissionId" binding:"required"`
}

type UpdateAppSubmissionByIdRequest struct {
	model.AppSubmission
}

func (s *AppSubmissionApi) UpdateAppSubmissionById(c *gin.Context) {
	uriParams := &UpdateAppSubmissionByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateAppSubmissionByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectSubmission.UpdateAppSubmissionById(
		&project_submission.UpdateAppSubmissionByIdParams{
			Context:         c,
			AppSubmissionId: uriParams.SubmissionId,
			AppSubmission:   &req.AppSubmission,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteAppSubmissionByIdUriParams struct {
	SubmissionId string `uri:"submissionId" binding:"required"`
}

func (s *AppSubmissionApi) DeleteAppSubmissionById(c *gin.Context) {
	uriParams := &DeleteAppSubmissionByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectSubmission.DeleteAppSubmissionById(
		&project_submission.DeleteAppSubmissionByIdParams{
			Context:         c,
			AppSubmissionId: uriParams.SubmissionId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
