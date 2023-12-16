package org

type Service interface {
	GetOrganizationList(*GetOrganizationListParams) ([]Organization, error)
	GetOrganizationById(*GetOrganizationByIdParams) (*Organization, error)
	GetOrganization(*GetOrganizationParams) (*Organization, error)
	CreateOrganization(*CreateOrganizationParams) (*Organization, error)
	UpdateOrganizationById(*UpdateOrganizationByIdParams) error
	DeleteOrganizationById(*DeleteOrganizationByIdParams) error
}
