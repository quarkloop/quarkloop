package quota

import "context"

type Service interface {
	// allocation quotas
	GetQuotasByUserId(context.Context, *GetQuotasByUserIdQuery) (Quota, error)
	GetQuotasByOrgId(context.Context, *GetQuotasByOrgIdQuery) ([]Quota, error)

	CheckCreateOrgQuotaReached(context.Context, *CheckCreateOrgQuotaReachedQuery) error
	CheckCreateOrgUserQuotaReached(context.Context, *CheckCreateOrgUserQuotaReachedQuery) error
	CheckCreateWorkspaceQuotaReached(context.Context, *CheckCreateWorkspaceQuotaReachedQuery) error
	CheckCreateProjectQuotaReached(context.Context, *CheckCreateProjectQuotaReachedQuery) error

	// overrides
	// OverrideOrgUserQuotaLimit(context.Context, int) error
	// OverrideWorkspaceQuotaLimit(context.Context, int) error
	// OverrideProjectQuotaLimit(context.Context, int) error
}
