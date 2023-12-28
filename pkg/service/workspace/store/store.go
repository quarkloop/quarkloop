package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type WorkspaceStore interface {
	ListWorkspaces(ctx context.Context, visibility model.ScopeVisibility, orgId []int) ([]workspace.Workspace, error)
	GetWorkspaceById(ctx context.Context, workspaceId int) (*workspace.Workspace, error)
	GetWorkspace(ctx context.Context, orgId int, workspace *workspace.Workspace) (*workspace.Workspace, error)
	CreateWorkspace(ctx context.Context, orgId int, workspace *workspace.Workspace) (*workspace.Workspace, error)
	UpdateWorkspaceById(ctx context.Context, workspaceId int, workspace *workspace.Workspace) error
	DeleteWorkspaceById(ctx context.Context, workspaceId int) error
}

type workspaceStore struct {
	Conn *pgx.Conn
}

func NewWorkspaceStore(conn *pgx.Conn) *workspaceStore {
	return &workspaceStore{
		Conn: conn,
	}
}
