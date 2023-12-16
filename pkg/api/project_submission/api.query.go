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

func (s *AppSubmissionApi) GetAppSubmissionList(c *gin.Context) {
	uriParams := &GetAppSubmissionListUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.projectSubmission.GetAppSubmissionList(
		&project_submission.GetAppSubmissionListParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, &projectList)
}

type GetAppSubmissionByIdUriParams struct {
	ProjectId    int    `uri:"projectId"`
	SubmissionId string `uri:"submissionId" binding:"required"`
}

func (s *AppSubmissionApi) GetAppSubmissionById(c *gin.Context) {
	uriParams := &GetAppSubmissionByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	project_submission, err := s.projectSubmission.GetAppSubmissionById(
		&project_submission.GetAppSubmissionByIdParams{
			Context:         c,
			ProjectId:       uriParams.ProjectId,
			AppSubmissionId: uriParams.SubmissionId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	c.JSON(http.StatusOK, project_submission)
}
