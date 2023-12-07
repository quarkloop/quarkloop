package project_submission

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type Api interface {
	GetAppSubmissionList(c *gin.Context)
	GetAppSubmissionById(c *gin.Context)
	CreateAppSubmission(c *gin.Context)
	UpdateAppSubmissionById(c *gin.Context)
	DeleteAppSubmissionById(c *gin.Context)
}

type AppSubmissionApi struct {
	projectSubmission project_submission.Service
}

func NewAppSubmissionApi(service project_submission.Service) *AppSubmissionApi {
	return &AppSubmissionApi{
		projectSubmission: service,
	}
}
