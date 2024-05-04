package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quarkloop/quarkloop/pkg/model"
)

type WorkspaceStore interface {
	// query
	GetWorkspaceId(ctx context.Context, query *GetWorkspaceIdQuery) (int64, int64, error)
	GetWorkspaceById(context.Context, *GetWorkspaceByIdQuery) (*model.Workspace, error)
	GetWorkspaceVisibilityById(context.Context, *GetWorkspaceVisibilityByIdQuery) (model.ScopeVisibility, error)
	GetWorkspaceList(context.Context, *GetWorkspaceListQuery) ([]*model.Workspace, error)
	// GetProjectList(context.Context, *GetProjectListQuery) ([]*model.Project, error)
	// GetUserAssignmentList(context.Context, *GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateWorkspace(context.Context, *CreateWorkspaceCommand) (*model.Workspace, error)
	UpdateWorkspaceById(context.Context, *UpdateWorkspaceByIdCommand) error
	DeleteWorkspaceById(context.Context, *DeleteWorkspaceByIdCommand) error
}

type workspaceStore struct {
	Conn *pgxpool.Pool
}

func NewWorkspaceStore(conn *pgxpool.Pool) *workspaceStore {
	return &workspaceStore{
		Conn: conn,
	}
}
