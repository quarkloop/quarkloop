package project_submission

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type Service interface {
	GetAppSubmissionList(context.Context, *GetAppSubmissionListParams) ([]model.AppSubmission, error)
	GetAppSubmissionById(context.Context, *GetAppSubmissionByIdParams) (*model.AppSubmission, error)
	CreateAppSubmission(context.Context, *CreateAppSubmissionParams) (*model.AppSubmission, error)
	UpdateAppSubmissionById(context.Context, *UpdateAppSubmissionByIdParams) error
	DeleteAppSubmissionById(context.Context, *DeleteAppSubmissionByIdParams) error
}
