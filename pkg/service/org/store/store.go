package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
)

type OrgStore interface {
	// query
	GetOrgById(context.Context, *org.GetOrgByIdQuery) (*model.Org, error)
	GetOrgVisibilityById(context.Context, *org.GetOrgVisibilityByIdQuery) (model.ScopeVisibility, error)
	GetOrgList(context.Context, *org.GetOrgListQuery) ([]*model.Org, error)
	GetWorkspaceList(context.Context, *org.GetWorkspaceListQuery) ([]*model.Workspace, error)
	GetProjectList(context.Context, *org.GetProjectListQuery) ([]*model.Project, error)
	//GetUserAssignmentList(context.Context, *org.GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateOrg(context.Context, *org.CreateOrgCommand) (*model.Org, error)
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
