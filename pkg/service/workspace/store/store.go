package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type WorkspaceStore interface {
	// query
	GetWorkspaceById(context.Context, *workspace.GetWorkspaceByIdQuery) (*model.Workspace, error)
	GetWorkspaceVisibilityById(context.Context, *workspace.GetWorkspaceVisibilityByIdQuery) (model.ScopeVisibility, error)
	GetWorkspaceList(context.Context, *workspace.GetWorkspaceListQuery) ([]*model.Workspace, error)
	GetProjectList(context.Context, *workspace.GetProjectListQuery) ([]*model.Project, error)
	GetUserAssignmentList(context.Context, *workspace.GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateWorkspace(context.Context, *workspace.CreateWorkspaceCommand) (*model.Workspace, error)
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
