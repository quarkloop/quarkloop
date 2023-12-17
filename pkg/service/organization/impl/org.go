package organization_impl

import (
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
	"github.com/quarkloop/quarkloop/pkg/service/organization/store"
)

type orgService struct {
	store store.OrgStore

	UserService      interface{}
	WorkspaceService interface{}
	QuotaService     interface{}
}

func NewOrganizationService(ds store.OrgStore) org.Service {
	return &orgService{
		store: ds,
	}
}

func (s *orgService) GetOrganizationList(p *org.GetOrganizationListParams) ([]org.Organization, error) {
	orgList, err := s.store.ListOrganizations(p.Context)
	if err != nil {
		return nil, err
	}

	for i := range orgList {
		org := &orgList[i]
		org.GeneratePath()
	}
	return orgList, nil
}

func (s *orgService) GetOrganizationById(p *org.GetOrganizationByIdParams) (*org.Organization, error) {
	org, err := s.store.GetOrganizationById(p.Context, p.OrgId)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) GetOrganization(p *org.GetOrganizationParams) (*org.Organization, error) {
	org, err := s.store.GetOrganization(p.Context, &p.Organization)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) CreateOrganization(p *org.CreateOrganizationParams) (*org.Organization, error) {
	org, err := s.store.CreateOrganization(p.Context, &p.Organization)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) UpdateOrganizationById(p *org.UpdateOrganizationByIdParams) error {
	err := s.store.UpdateOrganizationById(p.Context, p.OrgId, &p.Organization)
	return err
}

func (s *orgService) DeleteOrganizationById(p *org.DeleteOrganizationByIdParams) error {
	err := s.store.DeleteOrganizationById(p.Context, p.OrgId)
	return err
}
