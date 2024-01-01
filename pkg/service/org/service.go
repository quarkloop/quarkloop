package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type Service interface {
	// org
	GetOrgList(*gin.Context, *GetOrgListQuery) ([]*Org, error)
	GetOrgById(*gin.Context, *GetOrgByIdQuery) (*Org, error)
	// TODO: rewrite
	// GetOrg(context.Context, *GetOrgQuery) (*Org, error)
	CreateOrg(*gin.Context, *CreateOrgCommand) (*Org, error)
	UpdateOrgById(*gin.Context, *UpdateOrgByIdCommand) error
	DeleteOrgById(*gin.Context, *DeleteOrgByIdCommand) error

	// workspace
	GetWorkspaceList(*gin.Context, *GetWorkspaceListQuery) ([]*workspace.Workspace, error)

	// project
	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*project.Project, error)

	// user
	GetUserList(*gin.Context, *GetUserListQuery) ([]*user.User, error)
}
