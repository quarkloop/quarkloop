package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type ProjectStore interface {
	// query
	GetProjectById(ctx context.Context, projectId int) (*project.Project, error)
	GetProjectList(ctx context.Context, visibility model.ScopeVisibility, userId int) ([]*project.Project, error)
	GetUserAssignmentList(ctx context.Context, orgId, workspaceId, projectId int) ([]*user.UserAssignment, error)

	// mutation
	CreateProject(ctx context.Context, orgId, workspaceId int, p *project.Project) (*project.Project, error)
	UpdateProjectById(ctx context.Context, projectId int, project *project.Project) error
	DeleteProjectById(ctx context.Context, projectId int) error
}

type projectStore struct {
	Conn *pgx.Conn
}

func NewProjectStore(conn *pgx.Conn) *projectStore {
	return &projectStore{
		Conn: conn,
	}
}
