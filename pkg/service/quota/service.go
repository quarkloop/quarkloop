package quota

import "context"

type Service interface {
	// allocation quotas
	GetQuotasByUserId(context.Context, int) (Quota, error)
	GetQuotasByOrgId(context.Context, int) ([]Quota, error)

	CheckCreateOrgQuotaReached(ctx context.Context, userId int) error
	CheckCreateOrgUserQuotaReached(ctx context.Context, orgId int) error
	CheckCreateWorkspaceQuotaReached(ctx context.Context, orgId int) error
	CheckCreateProjectQuotaReached(ctx context.Context, orgId int) error

	// overrides
	// OverrideOrgUserQuotaLimit(context.Context, int) error
	// OverrideWorkspaceQuotaLimit(context.Context, int) error
	// OverrideProjectQuotaLimit(context.Context, int) error
}
