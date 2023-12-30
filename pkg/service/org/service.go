package org

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetOrganizationList(*gin.Context) ([]*Organization, error)
	GetOrganizationById(*gin.Context, *GetOrganizationByIdQuery) (*Organization, error)
	// TODO: rewrite
	// GetOrganization(context.Context, *GetOrganizationQuery) (*Organization, error)
	CreateOrganization(*gin.Context, *CreateOrganizationCommand) (*Organization, error)
	UpdateOrganizationById(*gin.Context, *UpdateOrganizationByIdCommand) error
	DeleteOrganizationById(*gin.Context, *DeleteOrganizationByIdCommand) error
}
