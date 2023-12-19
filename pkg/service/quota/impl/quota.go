package quota_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/quota/store"
)

type quotaService struct {
	store store.QuotaStore
}

func NewQuotaService(ds store.QuotaStore) quota.Service {
	return &quotaService{
		store: ds,
	}
}

func (s *quotaService) GetQuotasByUserId(ctx context.Context, userId int) (quota.Quota, error) {
	q, err := s.store.GetQuotasByUserId(ctx, userId)
	if err != nil {
		return quota.Quota{}, err
	}

	return q, nil
}

func (s *quotaService) GetQuotasByOrgId(ctx context.Context, orgId int) ([]quota.Quota, error) {
	q, err := s.store.GetQuotasByOrgId(ctx, orgId)
	if err != nil {
		return []quota.Quota{}, err
	}

	return q, nil
}

func (s *quotaService) GetQuotasByWorkspaceId(ctx context.Context, workspaceId int) ([]quota.Quota, error) {
	q, err := s.store.GetQuotasByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return []quota.Quota{}, err
	}

	return q, nil
}

func (s *quotaService) CheckOrgQuotaReached(ctx context.Context, userId int) (bool, error) {
	q, err := s.store.GetQuotasByUserId(ctx, userId)
	if err != nil {
		return false, err
	}

	if q.CheckQuotaReached() {
		return true, nil
	}
	return false, quota.ErrOrgQuotaReached
}

func (s *quotaService) CheckOrgUserQuotaReached(ctx context.Context, orgId int) (bool, error) {
	q, err := s.store.GetQuotasByOrgId(ctx, orgId)
	if err != nil {
		return false, err
	}

	for _, v := range q {
		if v.Feature == quota.OrgUserQuota {
			if v.CheckQuotaReached() {
				return true, nil
			}
			return false, quota.ErrOrgUserQuotaReached
		}
	}

	return false, quota.ErrUnableToFindFeature
}

func (s *quotaService) CheckWorkspaceQuotaReached(ctx context.Context, workspaceId int) (bool, error) {
	q, err := s.store.GetQuotasByWorkspaceId(ctx, workspaceId)
	if err != nil {
		return false, err
	}

	for _, v := range q {
		if v.Feature == quota.WorkspaceQuota {
			if v.CheckQuotaReached() {
				return true, nil
			}
			return false, quota.ErrWorkspaceQuotaReached
		}
	}

	return false, quota.ErrUnableToFindFeature
}
