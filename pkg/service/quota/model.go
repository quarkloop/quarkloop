package quota

import "errors"

type QuotaFeature string

var (
	OrgCount       QuotaFeature = "org_count"
	OrgUserCount   QuotaFeature = "org_user_count"
	WorkspaceCount QuotaFeature = "workspace_count"
	ProjectCount   QuotaFeature = "project_count"
)

var (
	OrgQuotaLimit       = 1
	OrgUserQuotaLimit   = 3
	WorkspaceQuotaLimit = 2
	ProjectQuotaLimit   = 2
)

var (
	ErrOrgQuotaReached       = errors.New("org quota reached")
	ErrOrgUserQuotaReached   = errors.New("org user quota reached")
	ErrWorkspaceQuotaReached = errors.New("workspace quota reached")
	ErrProjectQuotaReached   = errors.New("project quota reached")
	ErrUnableToFindFeature   = errors.New("unable to find quota feature")
)

type Quota struct {
	Feature QuotaFeature `json:"feature"`
	Limit   int          `json:"limit"`
	Metric  int          `json:"metric"`
}

func (q *Quota) CheckQuotaReached() bool {
	return q.Metric >= q.Limit
}

// GetQuotasByUserId
type GetQuotasByUserIdQuery struct {
	UserId int
}

// GetQuotasByOrgId
type GetQuotasByOrgIdQuery struct {
	OrgId int
}

// CheckCreateOrgQuota
type CheckCreateOrgQuotaQuery struct {
	UserId int
}

// CheckCreateOrgUserQuota
type CheckCreateOrgUserQuotaQuery struct {
	OrgId int
}

// CheckCreateWorkspaceQuota
type CheckCreateWorkspaceQuotaQuery struct {
	OrgId int
}

// CheckCreateProjectQuota
type CheckCreateProjectQuotaQuery struct {
	OrgId int
}
