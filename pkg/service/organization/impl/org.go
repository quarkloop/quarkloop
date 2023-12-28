package organization_impl

import (
	"context"
	"errors"

	"github.com/quarkloop/quarkloop/pkg/contextdata"
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

func (s *orgService) GetOrganizationList(ctx context.Context, params *org.GetOrganizationListParams) ([]org.Organization, error) {
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

func (s *orgService) GetOrganizationById(ctx context.Context, params *org.GetOrganizationByIdParams) (*org.Organization, error) {
	org, err := s.store.GetOrganizationById(ctx, params.OrgId)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

// func (s *orgService) GetOrganization(ctx context.Context, params *org.GetOrganizationParams) (*org.Organization, error) {
// 	org, err := s.store.GetOrganization(ctx, &params.Organization)
// 	if err != nil {
// 		return nil, err
// 	}

// 	org.GeneratePath()
// 	return org, nil
// }

func (s *orgService) CreateOrganization(ctx context.Context, params *org.CreateOrganizationParams) (*org.Organization, error) {
	if contextdata.IsUserAnonymous(ctx) {
		return nil, errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgCreate, &accesscontrol.EvaluateFilterParams{
		OrgId:  accesscontrol.GlobalOrgId, // TODO: move to contextdata
		UserId: user.GetId(),
	})
	if err != nil {
		return nil, err
	}

	userId := ctx.Value("userId").(int)

	// check quotas
	if err := s.quotaService.CheckCreateOrgQuotaReached(ctx, userId); err != nil {
		return nil, err
	}

	org, err := s.store.CreateOrganization(ctx, &params.Organization)
	if err != nil {
		return nil, err
	}
	org.GeneratePath()

	return org, nil
}

func (s *orgService) UpdateOrganizationById(ctx context.Context, params *org.UpdateOrganizationByIdParams) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgUpdate, &accesscontrol.EvaluateFilterParams{
		OrgId:  params.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return err
	}

	return s.store.UpdateOrganizationById(ctx, params.OrgId, &params.Organization)
}

func (s *orgService) DeleteOrganizationById(ctx context.Context, params *org.DeleteOrganizationByIdParams) error {
	if contextdata.IsUserAnonymous(ctx) {
		return errors.New("not authorized")
	}

	user := contextdata.GetUser(ctx)

	// check permissions
	err := s.aclService.Evaluate(ctx, accesscontrol.ActionOrgDelete, &accesscontrol.EvaluateFilterParams{
		OrgId:  params.OrgId,
		UserId: user.GetId(),
	})
	if err != nil {
		return err
	}

	return s.store.DeleteOrganizationById(ctx, params.OrgId)
}
