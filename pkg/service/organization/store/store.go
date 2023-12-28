package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
)

type OrgStore interface {
	ListOrganizations(ctx context.Context, visibility model.ScopeVisibility) ([]org.Organization, error)
	GetOrganizationById(ctx context.Context, orgId int) (*org.Organization, error)
	GetOrganization(ctx context.Context, org *org.Organization) (*org.Organization, error)
	CreateOrganization(ctx context.Context, org *org.Organization) (*org.Organization, error)
	UpdateOrganizationById(ctx context.Context, orgId int, org *org.Organization) error
	DeleteOrganizationById(ctx context.Context, orgId int) error
}

type orgStore struct {
	Conn *pgx.Conn
}

func NewOrgStore(conn *pgx.Conn) *orgStore {
	return &orgStore{
		Conn: conn,
	}
}
