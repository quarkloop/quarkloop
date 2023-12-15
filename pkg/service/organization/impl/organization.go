package organization_impl

import (
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/organization"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type orgService struct {
	UserService      interface{}
	WorkspaceService interface{}
	QuotaService     interface{}

	dataStore *repository.Repository
}

func NewOrganizationService(ds *repository.Repository) organization.Service {
	return &orgService{
		dataStore: ds,
	}
}

func (s *orgService) GetOrganizationList(p *organization.GetOrganizationListParams) ([]model.Organization, error) {
	orgList, err := s.dataStore.ListOrganizations(p.Context)
	return orgList, err
}

func (s *orgService) GetOrganizationById(p *organization.GetOrganizationByIdParams) (*model.Organization, error) {
	org, err := s.dataStore.GetOrganizationById(p.Context, p.OrgId)
	return org, err
}

func (s *orgService) GetOrganization(p *organization.GetOrganizationParams) (*model.Organization, error) {
	org, err := s.dataStore.GetOrganization(p.Context, &p.Organization)
	return org, err
}

func (s *orgService) CreateOrganization(p *organization.CreateOrganizationParams) (*model.Organization, error) {
	org, err := s.dataStore.CreateOrganization(p.Context, &p.Organization)
	return org, err
}

func (s *orgService) UpdateOrganizationById(p *organization.UpdateOrganizationByIdParams) error {
	err := s.dataStore.UpdateOrganizationById(p.Context, p.OrgId, &p.Organization)
	return err
}

func (s *orgService) DeleteOrganizationById(p *organization.DeleteOrganizationByIdParams) error {
	err := s.dataStore.DeleteOrganizationById(p.Context, p.OrgId)
	return err
}
