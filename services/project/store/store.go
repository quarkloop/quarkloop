package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quarkloop/quarkloop/pkg/model"
)

type ProjectStore interface {
	// query
	GetProjectId(ctx context.Context, query *GetProjectIdQuery) (int64, int64, int64, error)
	GetProjectById(context.Context, *GetProjectByIdQuery) (*model.Project, error)
	GetProjectVisibilityById(context.Context, *GetProjectVisibilityByIdQuery) (model.ScopeVisibility, error)
	GetProjectList(context.Context, *GetProjectListQuery) ([]*model.Project, error)
	//GetUserAssignmentList(context.Context, *GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateProject(context.Context, *CreateProjectCommand) (*model.Project, error)
	UpdateProjectById(context.Context, *UpdateProjectByIdCommand) error
	DeleteProjectById(context.Context, *DeleteProjectByIdCommand) error
}

type projectStore struct {
	Conn *pgxpool.Pool
}

func NewProjectStore(conn *pgxpool.Pool) *projectStore {
	return &projectStore{
		Conn: conn,
	}
}
