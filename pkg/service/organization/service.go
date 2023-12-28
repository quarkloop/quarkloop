package org

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetOrganizationList(*gin.Context) ([]Organization, error)
	GetOrganizationById(*gin.Context, *GetOrganizationByIdParams) (*Organization, error)
	// TODO: rewrite
	// GetOrganization(context.Context, *GetOrganizationParams) (*Organization, error)
	CreateOrganization(*gin.Context, *CreateOrganizationParams) (*Organization, error)
	UpdateOrganizationById(*gin.Context, *UpdateOrganizationByIdParams) error
	DeleteOrganizationById(*gin.Context, *DeleteOrganizationByIdParams) error
}
