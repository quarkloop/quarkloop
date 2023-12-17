package project_submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type CreateAppSubmissionUriParams struct {
	ProjectId int `uri:"projectId" binding:"required"`
}

type CreateAppSubmissionRequest struct {
	// UserId     string                  `json:"userId" binding:"required"`
	// Submission model.AppSubmission `json:"submission" binding:"required"`
	model.AppSubmission
}

func (s *AppSubmissionApi) CreateAppSubmission(ctx *gin.Context) {
	uriParams := &CreateAppSubmissionUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	req := &CreateAppSubmissionRequest{}
	if err := ctx.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	ws, err := s.projectSubmission.CreateAppSubmission(ctx, &project_submission.CreateAppSubmissionParams{
		UserId:        "req.UserId",
		ProjectId:     uriParams.ProjectId,
		AppSubmission: &req.AppSubmission,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, ws)
}

type UpdateAppSubmissionByIdUriParams struct {
	SubmissionId string `uri:"submissionId" binding:"required"`
}

type UpdateAppSubmissionByIdRequest struct {
	model.AppSubmission
}

func (s *AppSubmissionApi) UpdateAppSubmissionById(ctx *gin.Context) {
	uriParams := &UpdateAppSubmissionByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	req := &UpdateAppSubmissionByIdRequest{}
	if err := ctx.BindJSON(req); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.projectSubmission.UpdateAppSubmissionById(ctx, &project_submission.UpdateAppSubmissionByIdParams{
		AppSubmissionId: uriParams.SubmissionId,
		AppSubmission:   &req.AppSubmission,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type DeleteAppSubmissionByIdUriParams struct {
	SubmissionId string `uri:"submissionId" binding:"required"`
}

func (s *AppSubmissionApi) DeleteAppSubmissionById(ctx *gin.Context) {
	uriParams := &DeleteAppSubmissionByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	err := s.projectSubmission.DeleteAppSubmissionById(ctx, &project_submission.DeleteAppSubmissionByIdParams{
		AppSubmissionId: uriParams.SubmissionId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
