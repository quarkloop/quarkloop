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

func (s *quotaService) CheckCreateOrgQuotaReached(ctx context.Context, userId int) error {
	q, err := s.store.GetQuotasByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if q.CheckQuotaReached() {
		return quota.ErrOrgQuotaReached
	}
	return nil
}

func (s *quotaService) CheckCreateOrgUserQuotaReached(ctx context.Context, orgId int) error {
	q, err := s.store.GetQuotasByOrgId(ctx, orgId)
	if err != nil {
		return err
	}

	for _, v := range q {
		if v.Feature == quota.OrgUserCount {
			if v.CheckQuotaReached() {
				return quota.ErrOrgUserQuotaReached
			}
			return nil
		}
	}

	return quota.ErrUnableToFindFeature
}

func (s *quotaService) CheckCreateWorkspaceQuotaReached(ctx context.Context, orgId int) error {
	q, err := s.store.GetQuotasByOrgId(ctx, orgId)
	if err != nil {
		return err
	}

	for _, v := range q {
		if v.Feature == quota.WorkspaceCount {
			if v.CheckQuotaReached() {
				return quota.ErrWorkspaceQuotaReached
			}
			return nil
		}
	}

	return quota.ErrUnableToFindFeature
}

func (s *quotaService) CheckCreateProjectQuotaReached(ctx context.Context, orgId int) error {
	q, err := s.store.GetQuotasByOrgId(ctx, orgId)
	if err != nil {
		return err
	}

	for _, v := range q {
		if v.Feature == quota.ProjectCount {
			if v.CheckQuotaReached() {
				return quota.ErrProjectQuotaReached
			}
			return nil
		}
	}

	return quota.ErrUnableToFindFeature
}
