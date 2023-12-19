package quota

import "context"

type Service interface {
	// allocation quotas
	GetQuotasByUserId(context.Context, int) (Quota, error)
	GetQuotasByOrgId(context.Context, int) ([]Quota, error)
	GetQuotasByWorkspaceId(context.Context, int) ([]Quota, error)

	CheckOrgQuotaReached(context.Context, int) (bool, error)
	CheckOrgUserQuotaReached(context.Context, int) (bool, error)
	CheckWorkspaceQuotaReached(context.Context, int) (bool, error)

	// overrides
	// OverrideOrgUserQuotaLimit(context.Context, int) error
	// OverrideWorkspaceQuotaLimit(context.Context, int) error
	// OverrideProjectQuotaLimit(context.Context, int) error
}
