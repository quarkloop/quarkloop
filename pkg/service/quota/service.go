package quota

import "context"

type Service interface {
	// allocation quotas
	GetQuotasByUserId(context.Context, *GetQuotasByUserIdQuery) (Quota, error)
	GetQuotasByOrgId(context.Context, *GetQuotasByOrgIdQuery) ([]*Quota, error)

	CheckCreateOrgQuota(context.Context, *CheckCreateOrgQuotaQuery) error
	CheckCreateOrgUserQuota(context.Context, *CheckCreateOrgUserQuotaQuery) error
	CheckCreateWorkspaceQuota(context.Context, *CheckCreateWorkspaceQuotaQuery) error
	CheckCreateProjectQuota(context.Context, *CheckCreateProjectQuotaQuery) error

	// overrides
	// OverrideOrgUserQuotaLimit(context.Context, int) error
	// OverrideWorkspaceQuotaLimit(context.Context, int) error
	// OverrideProjectQuotaLimit(context.Context, int) error
}
