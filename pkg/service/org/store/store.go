package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type OrgStore interface {
	// org
	GetOrgList(ctx context.Context, visibility model.ScopeVisibility, userId int) ([]*org.Org, error)
	GetOrgById(ctx context.Context, orgId int) (*org.Org, error)
	GetOrg(ctx context.Context, org *org.Org) (*org.Org, error)
	CreateOrg(ctx context.Context, org *org.Org) (*org.Org, error)
	UpdateOrgById(ctx context.Context, orgId int, org *org.Org) error
	DeleteOrgById(ctx context.Context, orgId int) error

	// workspace
	GetWorkspaceList(ctx context.Context, visibility model.ScopeVisibility, orgId int) ([]*workspace.Workspace, error)

	// project
	GetProjectList(ctx context.Context, visibility model.ScopeVisibility, orgId int) ([]*project.Project, error)
}

type orgStore struct {
	Conn *pgx.Conn
}

func NewOrgStore(conn *pgx.Conn) *orgStore {
	return &orgStore{
		Conn: conn,
	}
}
