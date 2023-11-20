package organization

import "github.com/quarkloop/quarkloop/pkg/model"

type Service interface {
	GetOrganizationList(*GetOrganizationListParams) ([]model.Organization, error)
	GetOrganizationById(*GetOrganizationByIdParams) (*model.Organization, error)
	GetOrganization(*GetOrganizationParams) (*model.Organization, error)
	CreateOrganization(*CreateOrganizationParams) (*model.Organization, error)
	UpdateOrganizationById(*UpdateOrganizationByIdParams) error
	DeleteOrganizationById(*DeleteOrganizationByIdParams) error
}
