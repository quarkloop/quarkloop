package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type WorkspaceStore interface {
	// query
	GetWorkspaceById(context.Context, *workspace.GetWorkspaceByIdQuery) (*workspace.Workspace, error)
	GetWorkspaceList(context.Context, *workspace.GetWorkspaceListQuery) ([]*workspace.Workspace, error)
	GetProjectList(context.Context, *workspace.GetProjectListQuery) ([]*project.Project, error)
	GetUserAssignmentList(context.Context, *workspace.GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateWorkspace(context.Context, *workspace.CreateWorkspaceCommand) (*workspace.Workspace, error)
	UpdateWorkspaceById(context.Context, *workspace.UpdateWorkspaceByIdCommand) error
	DeleteWorkspaceById(context.Context, *workspace.DeleteWorkspaceByIdCommand) error
}

type workspaceStore struct {
	Conn *pgx.Conn
}

func NewWorkspaceStore(conn *pgx.Conn) *workspaceStore {
	return &workspaceStore{
		Conn: conn,
	}
}
