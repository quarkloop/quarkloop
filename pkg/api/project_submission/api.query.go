package project_submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type GetAppSubmissionListUriParams struct {
	ProjectId int `uri:"projectId"`
}

type GetAppSubmissionListResponse struct {
	api.ApiResponse
	Data []model.AppSubmission `json:"data"`
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

	res := &GetAppSubmissionListResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: projectList,
	}
	c.JSON(http.StatusOK, res)
}

type GetAppSubmissionByIdUriParams struct {
	ProjectId    int    `uri:"projectId"`
	SubmissionId string `uri:"submissionId" binding:"required"`
}

type GetAppSubmissionByIdResponse struct {
	api.ApiResponse
	Data model.AppSubmission `json:"data,omitempty"`
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

	res := &GetAppSubmissionByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *project_submission,
	}
	c.JSON(http.StatusOK, res)
}
