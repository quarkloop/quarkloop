package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

type Service interface {
	// org
	GetOrgList(*gin.Context) ([]*Org, error)
	GetOrgById(*gin.Context, *GetOrgByIdQuery) (*Org, error)
	// TODO: rewrite
	// GetOrg(context.Context, *GetOrgQuery) (*Org, error)
	CreateOrg(*gin.Context, *CreateOrgCommand) (*Org, error)
	UpdateOrgById(*gin.Context, *UpdateOrgByIdCommand) error
	DeleteOrgById(*gin.Context, *DeleteOrgByIdCommand) error

	// project
	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*project.Project, error)
}
