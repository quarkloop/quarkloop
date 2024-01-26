package quota_impl

import (
	"context"
	"fmt"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/quota/store"
)

type quotaService struct {
	store      store.QuotaStore
	aclService accesscontrol.Service
}

func NewQuotaService(ds store.QuotaStore, aclService accesscontrol.Service) quota.Service {
	return &quotaService{
		store:      ds,
		aclService: aclService,
	}
}

func (s *quotaService) GetQuotasByUserId(ctx context.Context, query *quota.GetQuotasByUserIdQuery) (quota.Quota, error) {
	q, err := s.store.GetQuotasByUserId(ctx, query.UserId)
	if err != nil {
		return quota.Quota{}, err
	}

	return q, nil
}

func (s *quotaService) GetQuotasByOrgId(ctx context.Context, query *quota.GetQuotasByOrgIdQuery) ([]*quota.Quota, error) {
	q, err := s.store.GetQuotasByOrgId(ctx, query.OrgId)
	if err != nil {
		return []*quota.Quota{}, err
	}

	return q, nil
}

func (s *quotaService) CheckCreateOrgQuota(ctx context.Context, query *quota.CheckCreateOrgQuotaQuery) error {
	orgQuery := &accesscontrol.GetOrgListQuery{
		UserId:     query.UserId,
		Permission: "owner",
	}
	orgList, err := s.aclService.GetOrgList(ctx, orgQuery)
	if err != nil {
		return err
	}

	q := quota.Quota{
		Feature: quota.OrgCount,
		Limit:   quota.OrgQuotaLimit,
		Metric:  int32(len(orgList)),
	}

	if q.CheckQuotaReached() {
		return quota.ErrOrgQuotaReached
	}
	return nil
}

func (s *quotaService) CheckCreateOrgUserQuota(ctx context.Context, query *quota.CheckCreateOrgUserQuotaQuery) error {
	quotaList, err := s.store.GetQuotasByOrgId(ctx, query.OrgId)
	if err != nil {
		return err
	}

	for _, q := range quotaList {
		if q.Feature == quota.OrgUserCount {
			q.ApplyLimit()
			if q.CheckQuotaReached() {
				return quota.ErrWorkspaceQuotaReached
			}
			return nil
		}
	}

	return quota.ErrUnableToFindFeature
}

func (s *quotaService) CheckCreateWorkspaceQuota(ctx context.Context, query *quota.CheckCreateWorkspaceQuotaQuery) error {
	quotaList, err := s.store.GetQuotasByOrgId(ctx, query.OrgId)
	fmt.Printf("\nWorkspaceQuota Service => %+v => %+v\n\n", quotaList, err)
	if err != nil {
		return err
	}

	for _, q := range quotaList {
		if q.Feature == quota.WorkspaceCount {
			q.ApplyLimit()
			if q.CheckQuotaReached() {
				return quota.ErrWorkspaceQuotaReached
			}
			return nil
		}
	}

	return quota.ErrUnableToFindFeature
}

func (s *quotaService) CheckCreateProjectQuota(ctx context.Context, query *quota.CheckCreateProjectQuotaQuery) error {
	quotaList, err := s.store.GetQuotasByOrgId(ctx, query.OrgId)
	if err != nil {
		return err
	}

	for _, q := range quotaList {
		if q.Feature == quota.ProjectCount {
			q.ApplyLimit()
			if q.CheckQuotaReached() {
				return quota.ErrWorkspaceQuotaReached
			}
			return nil
		}
	}

	return quota.ErrUnableToFindFeature
}
