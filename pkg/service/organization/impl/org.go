package organization_impl

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
	"github.com/quarkloop/quarkloop/pkg/service/organization/store"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
)

type orgService struct {
	store        store.OrgStore
	aclService   accesscontrol.Service
	quotaService quota.Service
}

func NewOrganizationService(ds store.OrgStore, aclService accesscontrol.Service, quotaService quota.Service) org.Service {
	return &orgService{
		store:        ds,
		aclService:   aclService,
		quotaService: quotaService,
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
	permission, err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgCreate, &accesscontrol.EvaluateFilterParams{
		OrgId:  accesscontrol.GlobalOrgId,
		UserId: 0,
	})
	if err != nil {
		return nil, err
	}
	if !permission {
		return nil, accesscontrol.ErrPermissionDenied
	}

	userId := ctx.Value("userId").(int)
	if err := s.quotaService.CheckCreateOrgQuotaReached(ctx, userId); err != nil {
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
	permission, err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgUpdate, &accesscontrol.EvaluateFilterParams{
		OrgId:  p.OrgId,
		UserId: 0,
	})
	if err != nil {
		return err
	}
	if !permission {
		return accesscontrol.ErrPermissionDenied
	}

	return s.store.UpdateOrganizationById(ctx, p.OrgId, &p.Organization)
}

func (s *orgService) DeleteOrganizationById(ctx context.Context, p *org.DeleteOrganizationByIdParams) error {
	permission, err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgDelete, &accesscontrol.EvaluateFilterParams{
		OrgId:  p.OrgId,
		UserId: 0,
	})
	if err != nil {
		return err
	}
	if !permission {
		return accesscontrol.ErrPermissionDenied
	}

	return s.store.DeleteOrganizationById(ctx, p.OrgId)
}
