package quota

import "errors"

type QuotaFeature string

var (
	OrgQuota       QuotaFeature = "org_quota"
	OrgUserQuota   QuotaFeature = "org_user_quota"
	WorkspaceQuota QuotaFeature = "workspace_quota"
	ProjectQuota   QuotaFeature = "project_quota"
)

var (
	OrgQuotaLimit       = 1
	OrgUserQuotaLimit   = 3
	WorkspaceQuotaLimit = 2
	ProjectQuotaLimit   = 2
)

var (
	ErrOrgQuotaReached       error = errors.New("org quota reached")
	ErrOrgUserQuotaReached   error = errors.New("org user quota reached")
	ErrWorkspaceQuotaReached error = errors.New("workspace quota reached")
	ErrUnableToFindFeature   error = errors.New("unable to find quota feature")
)

type Quota struct {
	Feature QuotaFeature `json:"feature,omitempty"`
	Limit   int          `json:"limit,omitempty"`
	Metric  int          `json:"metric,omitempty"`
}

func (q *Quota) CheckQuotaReached() bool {
	return q.Metric >= q.Limit
}
