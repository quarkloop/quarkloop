package project_submission_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type projectSubmission struct {
	dataStore *repository.Repository
}

func NewAppSubmissionService(ds *repository.Repository) project_submission.Service {
	return &projectSubmission{
		dataStore: ds,
	}
}

func (s *projectSubmission) GetAppSubmissionList(p *project_submission.GetAppSubmissionListParams) ([]model.AppSubmission, error) {
	projectList, err := s.dataStore.ListAppSubmissions(p.Context, p.ProjectId)
	return projectList, err
}

func (s *projectSubmission) GetAppSubmissionById(p *project_submission.GetAppSubmissionByIdParams) (*model.AppSubmission, error) {
	project, err := s.dataStore.GetAppSubmissionById(p.Context, p.ProjectId, p.AppSubmissionId)
	return project, err
}

func (s *projectSubmission) CreateAppSubmission(p *project_submission.CreateAppSubmissionParams) (*model.AppSubmission, error) {
	project, err := s.dataStore.CreateAppSubmission(p.Context, p.UserId, p.ProjectId, p.AppSubmission)
	return project, err
}

func (s *projectSubmission) UpdateAppSubmissionById(p *project_submission.UpdateAppSubmissionByIdParams) error {
	err := s.dataStore.UpdateAppSubmissionById(p.Context, p.AppSubmissionId, p.AppSubmission)
	return err
}

func (s *projectSubmission) DeleteAppSubmissionById(p *project_submission.DeleteAppSubmissionByIdParams) error {
	err := s.dataStore.DeleteAppSubmissionById(p.Context, p.AppSubmissionId)
	return err
}
