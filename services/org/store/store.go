package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quarkloop/quarkloop/pkg/model"
)

type OrgStore interface {
	// query
	GetOrgId(context.Context, string) (int64, error)
	GetOrgById(context.Context, *GetOrgByIdQuery) (*model.Org, error)
	GetOrgVisibilityById(context.Context, *GetOrgVisibilityByIdQuery) (model.ScopeVisibility, error)
	GetOrgList(context.Context, *GetOrgListQuery) ([]*model.Org, error)
	GetWorkspaceList(context.Context, *GetWorkspaceListQuery) ([]*model.Workspace, error)
	// GetProjectList(context.Context, *GetProjectListQuery) ([]*model.Project, error)
	// GetUserAssignmentList(context.Context, *GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateOrg(context.Context, *CreateOrgCommand) (*model.Org, error)
	UpdateOrgById(context.Context, *UpdateOrgByIdCommand) error
	DeleteOrgById(context.Context, *DeleteOrgByIdCommand) error
}

type orgStore struct {
	Conn *pgxpool.Pool
}

func NewOrgStore(conn *pgxpool.Pool) *orgStore {
	return &orgStore{
		Conn: conn,
	}
}
