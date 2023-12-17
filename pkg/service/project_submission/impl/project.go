package project_submission_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project_submission"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type projectSubmission struct {
	store *repository.Repository
}

func NewAppSubmissionService(ds *repository.Repository) project_submission.Service {
	return &projectSubmission{
		store: ds,
	}
}

func (s *projectSubmission) GetAppSubmissionList(ctx context.Context, p *project_submission.GetAppSubmissionListParams) ([]model.AppSubmission, error) {
	projectList, err := s.store.ListAppSubmissions(ctx, p.ProjectId)
	return projectList, err
}

func (s *projectSubmission) GetAppSubmissionById(ctx context.Context, p *project_submission.GetAppSubmissionByIdParams) (*model.AppSubmission, error) {
	project, err := s.store.GetAppSubmissionById(ctx, p.ProjectId, p.AppSubmissionId)
	return project, err
}

func (s *projectSubmission) CreateAppSubmission(ctx context.Context, p *project_submission.CreateAppSubmissionParams) (*model.AppSubmission, error) {
	project, err := s.store.CreateAppSubmission(ctx, p.UserId, p.ProjectId, p.AppSubmission)
	return project, err
}

func (s *projectSubmission) UpdateAppSubmissionById(ctx context.Context, p *project_submission.UpdateAppSubmissionByIdParams) error {
	err := s.store.UpdateAppSubmissionById(ctx, p.AppSubmissionId, p.AppSubmission)
	return err
}

func (s *projectSubmission) DeleteAppSubmissionById(ctx context.Context, p *project_submission.DeleteAppSubmissionByIdParams) error {
	err := s.store.DeleteAppSubmissionById(ctx, p.AppSubmissionId)
	return err
}
