package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type OrgStore interface {
	// query
	GetOrgById(context.Context, *org.GetOrgByIdQuery) (*org.Org, error)
	GetOrgList(context.Context, *org.GetOrgListQuery) ([]*org.Org, error)
	GetWorkspaceList(context.Context, *org.GetWorkspaceListQuery) ([]*workspace.Workspace, error)
	GetProjectList(context.Context, *org.GetProjectListQuery) ([]*project.Project, error)
	GetUserAssignmentList(context.Context, *org.GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateOrg(context.Context, *org.CreateOrgCommand) (*org.Org, error)
	UpdateOrgById(context.Context, *org.UpdateOrgByIdCommand) error
	DeleteOrgById(context.Context, *org.DeleteOrgByIdCommand) error
}

type orgStore struct {
	Conn *pgx.Conn
}

func NewOrgStore(conn *pgx.Conn) *orgStore {
	return &orgStore{
		Conn: conn,
	}
}
