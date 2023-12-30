package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type ProjectStore interface {
	GetProjectList(ctx context.Context, visibility model.ScopeVisibility, orgId []int, workspaceId []int) ([]*project.Project, error)
	GetProjectById(ctx context.Context, projectId int) (*project.Project, error)
	GetProject(ctx context.Context, p *project.Project) (*project.Project, error)
	CreateProject(ctx context.Context, orgId int, workspaceId int, p *project.Project) (*project.Project, error)
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
