package organization_impl

import (
	"context"

	org "github.com/quarkloop/quarkloop/pkg/service/organization"
	"github.com/quarkloop/quarkloop/pkg/service/organization/store"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
)

type orgService struct {
	store        store.OrgStore
	quotaService quota.Service
}

func NewOrganizationService(ds store.OrgStore, quota quota.Service) org.Service {
	return &orgService{
		store:        ds,
		quotaService: quota,
	}
}

func (s *orgService) GetOrganizationList(ctx context.Context, p *org.GetOrganizationListParams) ([]org.Organization, error) {
	orgList, err := s.store.ListOrganizations(ctx)
	if err != nil {
		return nil, err
	}

	for i := range orgList {
		org := &orgList[i]
		org.GeneratePath()
	}
	return orgList, nil
}

func (s *orgService) GetOrganizationById(ctx context.Context, p *org.GetOrganizationByIdParams) (*org.Organization, error) {
	org, err := s.store.GetOrganizationById(ctx, p.OrgId)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) GetOrganization(ctx context.Context, p *org.GetOrganizationParams) (*org.Organization, error) {
	org, err := s.store.GetOrganization(ctx, &p.Organization)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) CreateOrganization(ctx context.Context, p *org.CreateOrganizationParams) (*org.Organization, error) {
	userId := ctx.Value("userId").(int)
	_, err := s.quotaService.CheckOrgQuotaReached(ctx, userId)
	if err != nil {
		return nil, err
	}

	org, err := s.store.CreateOrganization(ctx, &p.Organization)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) UpdateOrganizationById(ctx context.Context, p *org.UpdateOrganizationByIdParams) error {
	err := s.store.UpdateOrganizationById(ctx, p.OrgId, &p.Organization)
	return err
}

func (s *orgService) DeleteOrganizationById(ctx context.Context, p *org.DeleteOrganizationByIdParams) error {
	err := s.store.DeleteOrganizationById(ctx, p.OrgId)
	return err
}
