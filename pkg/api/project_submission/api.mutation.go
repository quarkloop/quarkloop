package project_submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type CreateProjectSubmissionUriParams struct {
	ProjectId string `uri:"projectId" binding:"required"`
}

type CreateProjectSubmissionRequest struct {
	// UserId     string                  `json:"userId" binding:"required"`
	// Submission model.ProjectSubmission `json:"submission" binding:"required"`
	model.ProjectSubmission
}

type CreateProjectSubmissionResponse struct {
	api.ApiResponse
	Data model.ProjectSubmission `json:"data,omitempty"`
}

func (s *ProjectSubmissionApi) CreateProjectSubmission(c *gin.Context) {
	uriParams := &CreateProjectSubmissionUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &CreateProjectSubmissionRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	ws, err := s.projectSubmission.CreateProjectSubmission(
		&project_submission.CreateProjectSubmissionParams{
			Context:           c,
			UserId:            "req.UserId",
			ProjectId:         uriParams.ProjectId,
			ProjectSubmission: &req.ProjectSubmission,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &CreateProjectSubmissionResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusCreated,
			StatusString: "Created",
		},
		Data: *ws,
	}
	c.JSON(http.StatusCreated, res)
}

type UpdateProjectSubmissionByIdUriParams struct {
	SubmissionId string `uri:"submissionId" binding:"required"`
}

type UpdateProjectSubmissionByIdRequest struct {
	model.ProjectSubmission
}

func (s *ProjectSubmissionApi) UpdateProjectSubmissionById(c *gin.Context) {
	uriParams := &UpdateProjectSubmissionByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	req := &UpdateProjectSubmissionByIdRequest{}
	if err := c.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectSubmission.UpdateProjectSubmissionById(
		&project_submission.UpdateProjectSubmissionByIdParams{
			Context:             c,
			ProjectSubmissionId: uriParams.SubmissionId,
			ProjectSubmission:   &req.ProjectSubmission,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type DeleteProjectSubmissionByIdUriParams struct {
	SubmissionId string `uri:"submissionId" binding:"required"`
}

func (s *ProjectSubmissionApi) DeleteProjectSubmissionById(c *gin.Context) {
	uriParams := &DeleteProjectSubmissionByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	err := s.projectSubmission.DeleteProjectSubmissionById(
		&project_submission.DeleteProjectSubmissionByIdParams{
			Context:             c,
			ProjectSubmissionId: uriParams.SubmissionId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
