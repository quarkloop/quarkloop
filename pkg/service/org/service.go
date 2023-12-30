package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type Service interface {
	// org
	GetOrganizationList(*gin.Context) ([]*Organization, error)
	GetOrganizationById(*gin.Context, *GetOrganizationByIdQuery) (*Organization, error)
	// TODO: rewrite
	// GetOrganization(context.Context, *GetOrganizationQuery) (*Organization, error)
	CreateOrganization(*gin.Context, *CreateOrganizationCommand) (*Organization, error)
	UpdateOrganizationById(*gin.Context, *UpdateOrganizationByIdCommand) error
	DeleteOrganizationById(*gin.Context, *DeleteOrganizationByIdCommand) error

	// project
	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*project.Project, error)
}
