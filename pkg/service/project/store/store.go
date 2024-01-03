package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
)

type ProjectStore interface {
	// query
	GetProjectById(context.Context, *project.GetProjectByIdQuery) (*project.Project, error)
	GetProjectList(context.Context, *project.GetProjectListQuery) ([]*project.Project, error)
	GetUserAssignmentList(context.Context, *project.GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateProject(context.Context, *project.CreateProjectCommand) (*project.Project, error)
	UpdateProjectById(context.Context, *project.UpdateProjectByIdCommand) error
	DeleteProjectById(context.Context, *project.DeleteProjectByIdCommand) error
}

type projectStore struct {
	Conn *pgx.Conn
}

func NewProjectStore(conn *pgx.Conn) *projectStore {
	return &projectStore{
		Conn: conn,
	}
}
