package project_submission

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	GetProjectSubmissionList(*GetProjectSubmissionListParams) ([]model.ProjectSubmission, error)
	GetProjectSubmissionById(*GetProjectSubmissionByIdParams) (*model.ProjectSubmission, error)
	CreateProjectSubmission(*CreateProjectSubmissionParams) (*model.ProjectSubmission, error)
	UpdateProjectSubmissionById(*UpdateProjectSubmissionByIdParams) error
	DeleteProjectSubmissionById(*DeleteProjectSubmissionByIdParams) error
}
