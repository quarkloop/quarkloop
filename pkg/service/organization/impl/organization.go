package organization_impl

import (
	org "github.com/quarkloop/quarkloop/pkg/service/organization"
	"github.com/quarkloop/quarkloop/pkg/service/organization/store"
)

type orgService struct {
	UserService      interface{}
	WorkspaceService interface{}
	QuotaService     interface{}

	dataStore store.OrgStore
}

func NewOrganizationService(ds store.OrgStore) org.Service {
	return &orgService{
		dataStore: ds,
	}
}

func (s *orgService) GetOrganizationList(p *org.GetOrganizationListParams) ([]org.Organization, error) {
	orgList, err := s.dataStore.ListOrganizations(p.Context)
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
	org, err := s.dataStore.GetOrganizationById(p.Context, p.OrgId)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) GetOrganization(p *org.GetOrganizationParams) (*org.Organization, error) {
	org, err := s.dataStore.GetOrganization(p.Context, &p.Organization)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) CreateOrganization(p *org.CreateOrganizationParams) (*org.Organization, error) {
	org, err := s.dataStore.CreateOrganization(p.Context, &p.Organization)
	if err != nil {
		return nil, err
	}

	org.GeneratePath()
	return org, nil
}

func (s *orgService) UpdateOrganizationById(p *org.UpdateOrganizationByIdParams) error {
	err := s.dataStore.UpdateOrganizationById(p.Context, p.OrgId, &p.Organization)
	return err
}

func (s *orgService) DeleteOrganizationById(p *org.DeleteOrganizationByIdParams) error {
	err := s.dataStore.DeleteOrganizationById(p.Context, p.OrgId)
	return err
}
