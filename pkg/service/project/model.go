package project

import (
	"errors"

	"github.com/quarkloop/quarkloop/pkg/model"
)

var (
	ErrProjectNotFound      = errors.New("project not found")
	ErrProjectAlreadyExists = errors.New("project with same scopeId already exists")
)

// GetProjectById

type GetProjectByIdUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
	ProjectId   int32 `uri:"projectId" binding:"required"`
}

type GetProjectByIdQuery struct {
	OrgId       int32
	WorkspaceId int32
	ProjectId   int32
}

// GetProjectVisibilityById

type GetProjectVisibilityByIdQuery struct {
	OrgId       int32
	WorkspaceId int32
	ProjectId   int32
}

// GetProjectList

type GetProjectListQuery struct {
	ProjectIdList []int32
	Visibility    model.ScopeVisibility
}

//  CreateProject

type CreateProjectUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
}

type CreateProjectCommand struct {
	OrgId       int32
	WorkspaceId int32
	CreatedBy   string

	ScopeId     string                `json:"sid"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Visibility  model.ScopeVisibility `json:"visibility"`
}

// UpdateProjectById

type UpdateProjectByIdUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
	ProjectId   int32 `uri:"projectId" binding:"required"`
}

type UpdateProjectByIdCommand struct {
	OrgId       int32
	WorkspaceId int32
	ProjectId   int32
	UpdatedBy   string

	ScopeId     string                `json:"sid,omitempty"`
	Name        string                `json:"name,omitempty"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility,omitempty"`
}

// DeleteProjectById

type DeleteProjectByIdUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
	ProjectId   int32 `uri:"projectId" binding:"required"`
}

type DeleteProjectByIdCommand struct {
	OrgId       int32
	WorkspaceId int32
	ProjectId   int32
}

// GetMemberList

type GetMemberListUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
	ProjectId   int32 `uri:"projectId" binding:"required"`
}

type GetMemberListQuery struct {
	OrgId       int32
	WorkspaceId int32
	ProjectId   int32
}

// GetUserAssignmentList

type GetUserAssignmentListQuery struct {
	OrgId       int32
	WorkspaceId int32
	ProjectId   int32
}
