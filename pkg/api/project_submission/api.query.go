package project_submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type GetProjectSubmissionListUriParams struct {
	ProjectId string `uri:"projectId"`
}

type GetProjectSubmissionListResponse struct {
	api.ApiResponse
	Data []model.ProjectSubmission `json:"data"`
}

func (s *ProjectSubmissionApi) GetProjectSubmissionList(c *gin.Context) {
	uriParams := &GetProjectSubmissionListUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	projectList, err := s.projectSubmission.GetProjectSubmissionList(
		&project_submission.GetProjectSubmissionListParams{
			Context:   c,
			ProjectId: uriParams.ProjectId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetProjectSubmissionListResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: projectList,
	}
	c.JSON(http.StatusOK, res)
}

type GetProjectSubmissionByIdUriParams struct {
	ProjectId    string `uri:"projectId"`
	SubmissionId string `uri:"submissionId" binding:"required"`
}

type GetProjectSubmissionByIdResponse struct {
	api.ApiResponse
	Data model.ProjectSubmission `json:"data,omitempty"`
}

func (s *ProjectSubmissionApi) GetProjectSubmissionById(c *gin.Context) {
	uriParams := &GetProjectSubmissionByIdUriParams{}
	if err := c.ShouldBindUri(uriParams); err != nil {
		api.AbortWithBadRequestJSON(c, err)
		return
	}

	// query service
	project_submission, err := s.projectSubmission.GetProjectSubmissionById(
		&project_submission.GetProjectSubmissionByIdParams{
			Context:             c,
			ProjectId:           uriParams.ProjectId,
			ProjectSubmissionId: uriParams.SubmissionId,
		},
	)
	if err != nil {
		api.AbortWithInternalServerErrorJSON(c, err)
		return
	}

	res := &GetProjectSubmissionByIdResponse{
		ApiResponse: api.ApiResponse{
			Status:       http.StatusOK,
			StatusString: "OK",
		},
		Data: *project_submission,
	}
	c.JSON(http.StatusOK, res)
}
