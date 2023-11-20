package organization

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetOrganizationListParams struct {
	Context context.Context
}

type GetOrganizationByIdParams struct {
	Context context.Context
	Id      string
}

type GetOrganizationParams struct {
	Context      context.Context
	Organization model.Organization
}

type CreateOrganizationParams struct {
	Context      context.Context
	Organization model.Organization
}

type UpdateOrganizationByIdParams struct {
	Context      context.Context
	OrgId        string
	Organization model.Organization
}

type DeleteOrganizationByIdParams struct {
	Context context.Context
	OrgId   string
}
