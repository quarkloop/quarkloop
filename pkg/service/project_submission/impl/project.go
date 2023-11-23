package project_submission_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type projectSubmission struct {
	dataStore *repository.Repository
}

func NewProjectSubmissionService(ds *repository.Repository) project_submission.Service {
	return &projectSubmission{
		dataStore: ds,
	}
}

func (s *projectSubmission) GetProjectSubmissionList(p *project_submission.GetProjectSubmissionListParams) ([]model.ProjectSubmission, error) {
	projectList, err := s.dataStore.ListProjectSubmissions(p.Context, p.ProjectId)
	return projectList, err
}

func (s *projectSubmission) GetProjectSubmissionById(p *project_submission.GetProjectSubmissionByIdParams) (*model.ProjectSubmission, error) {
	project, err := s.dataStore.FindUniqueProjectSubmission(p.Context, p.ProjectId, p.ProjectSubmissionId)
	return project, err
}

func (s *projectSubmission) CreateProjectSubmission(p *project_submission.CreateProjectSubmissionParams) (*model.ProjectSubmission, error) {
	project, err := s.dataStore.CreateProjectSubmission(p.Context, p.UserId, p.ProjectId, p.ProjectSubmission)
	return project, err
}

func (s *projectSubmission) UpdateProjectSubmissionById(p *project_submission.UpdateProjectSubmissionByIdParams) error {
	err := s.dataStore.UpdateProjectSubmissionById(p.Context, p.ProjectSubmissionId, p.ProjectSubmission)
	return err
}

func (s *projectSubmission) DeleteProjectSubmissionById(p *project_submission.DeleteProjectSubmissionByIdParams) error {
	err := s.dataStore.DeleteProjectSubmissionById(p.Context, p.ProjectSubmissionId)
	return err
}
