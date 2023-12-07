package project_submission

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	GetAppSubmissionList(*GetAppSubmissionListParams) ([]model.AppSubmission, error)
	GetAppSubmissionById(*GetAppSubmissionByIdParams) (*model.AppSubmission, error)
	CreateAppSubmission(*CreateAppSubmissionParams) (*model.AppSubmission, error)
	UpdateAppSubmissionById(*UpdateAppSubmissionByIdParams) error
	DeleteAppSubmissionById(*DeleteAppSubmissionByIdParams) error
}
