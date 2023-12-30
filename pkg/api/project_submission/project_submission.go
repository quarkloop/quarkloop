package project_submission

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
)

type Api interface {
	GetAppSubmissionList(*gin.Context)
	GetAppSubmissionById(*gin.Context)
	CreateAppSubmission(*gin.Context)
	UpdateAppSubmissionById(*gin.Context)
	DeleteAppSubmissionById(*gin.Context)
}

type AppSubmissionApi struct {
	projectSubmission project_submission.Service
}

func NewAppSubmissionApi(service project_submission.Service) *AppSubmissionApi {
	return &AppSubmissionApi{
		projectSubmission: service,
	}
}
