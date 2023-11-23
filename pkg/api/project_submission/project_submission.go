package project_submission

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type Api interface {
	GetProjectSubmissionList(c *gin.Context)
	GetProjectSubmissionById(c *gin.Context)
	CreateProjectSubmission(c *gin.Context)
	UpdateProjectSubmissionById(c *gin.Context)
	DeleteProjectSubmissionById(c *gin.Context)
}

type ProjectSubmissionApi struct {
	projectSubmission project_submission.Service
}

func NewProjectSubmissionApi(service project_submission.Service) *ProjectSubmissionApi {
	return &ProjectSubmissionApi{
		projectSubmission: service,
	}
}
