package org

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

type Service interface {
	// query
	GetOrgById(*gin.Context, *GetOrgByIdQuery) (*Org, error)
	GetOrgList(*gin.Context, *GetOrgListQuery) ([]*Org, error)
	GetWorkspaceList(*gin.Context, *GetWorkspaceListQuery) ([]*workspace.Workspace, error)
	GetProjectList(*gin.Context, *GetProjectListQuery) ([]*project.Project, error)
	GetUserAssignmentList(*gin.Context, *GetUserAssignmentListQuery) ([]*user.UserAssignment, error)

	// mutation
	CreateOrg(*gin.Context, *CreateOrgCommand) (*Org, error)
	UpdateOrgById(*gin.Context, *UpdateOrgByIdCommand) error
	DeleteOrgById(*gin.Context, *DeleteOrgByIdCommand) error
}
