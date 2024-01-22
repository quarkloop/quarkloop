package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
)

type QuotaStore interface {
	// allocation quotas
	GetQuotasByUserId(context.Context, int32) (quota.Quota, error)
	GetQuotasByOrgId(context.Context, int32) ([]quota.Quota, error)

	// quota overrides
	UpdateQuotaLimits(context.Context, int, QuoataLimit) error
}

type quotaStore struct {
	Conn *pgx.Conn
}

func NewQuotaStore(conn *pgx.Conn) *quotaStore {
	return &quotaStore{
		Conn: conn,
	}
}

type QuoataLimit struct {
	OrgLimit                 int
	OrgUserLimit             int
	WorkspacePerOrgLimit     int
	ProjectPerWorkspaceLimit int
}
