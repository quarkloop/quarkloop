package project_submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type GetAppSubmissionListUriParams struct {
	ProjectId int `uri:"projectId"`
}

func (s *AppSubmissionApi) GetAppSubmissionList(ctx *gin.Context) {
	uriParams := &GetAppSubmissionListUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	projectList, err := s.projectSubmission.GetAppSubmissionList(ctx, &project_submission.GetAppSubmissionListParams{
		ProjectId: uriParams.ProjectId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &projectList)
}

type GetAppSubmissionByIdUriParams struct {
	ProjectId    int    `uri:"projectId"`
	SubmissionId string `uri:"submissionId" binding:"required"`
}

func (s *AppSubmissionApi) GetAppSubmissionById(ctx *gin.Context) {
	uriParams := &GetAppSubmissionByIdUriParams{}
	if err := ctx.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(ctx, err)
		return
	}

	// query service
	project_submission, err := s.projectSubmission.GetAppSubmissionById(ctx, &project_submission.GetAppSubmissionByIdParams{
		ProjectId:       uriParams.ProjectId,
		AppSubmissionId: uriParams.SubmissionId,
	},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, project_submission)
}
