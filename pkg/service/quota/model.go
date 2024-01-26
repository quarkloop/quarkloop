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
	OrgQuotaLimit       int32 = 2
	OrgUserQuotaLimit   int32 = 3
	WorkspaceQuotaLimit int32 = 2
	ProjectQuotaLimit   int32 = 2
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
	Limit   int32        `json:"limit"`
	Metric  int32        `json:"metric"`
}

func (q *Quota) ApplyLimit() {
	switch q.Feature {
	case OrgCount:
		q.Limit = OrgQuotaLimit
	case OrgUserCount:
		q.Limit = OrgUserQuotaLimit
	case WorkspaceCount:
		q.Limit = WorkspaceQuotaLimit
	case ProjectCount:
		q.Limit = ProjectQuotaLimit
	default:
		panic("unknown quota feature found")
	}
}

func (q *Quota) CheckQuotaReached() bool {
	return q.Metric >= q.Limit
}

// GetQuotasByUserId
type GetQuotasByUserIdQuery struct {
	UserId int32
}

// GetQuotasByOrgId
type GetQuotasByOrgIdQuery struct {
	OrgId int32
}

// CheckCreateOrgQuota
type CheckCreateOrgQuotaQuery struct {
	UserId int32
}

// CheckCreateOrgUserQuota
type CheckCreateOrgUserQuotaQuery struct {
	OrgId int32
}

// CheckCreateWorkspaceQuota
type CheckCreateWorkspaceQuotaQuery struct {
	OrgId int32
}

// CheckCreateProjectQuota
type CheckCreateProjectQuotaQuery struct {
	OrgId int32
}
