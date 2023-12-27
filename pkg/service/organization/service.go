package org

import (
	"context"
)

type Service interface {
	GetOrganizationList(context.Context, *GetOrganizationListParams) ([]Organization, error)
	GetOrganizationById(context.Context, *GetOrganizationByIdParams) (*Organization, error)
	GetOrganization(context.Context, *GetOrganizationParams) (*Organization, error)
	CreateOrganization(context.Context, *CreateOrganizationParams) (*Organization, error)
	UpdateOrganizationById(context.Context, *UpdateOrganizationByIdParams) error
	DeleteOrganizationById(context.Context, *DeleteOrganizationByIdParams) error
}
