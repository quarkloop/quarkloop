package workspace

import (
	"errors"

	"github.com/quarkloop/quarkloop/pkg/model"
)

var (
	ErrWorkspaceNotFound      = errors.New("workspace not found")
	ErrWorkspaceAlreadyExists = errors.New("workspace with same scopeId already exists")
)

// GetWorkspaceList

type GetWorkspaceListQuery struct {
	WorkspaceIdList []int32
	Visibility      model.ScopeVisibility
}

// GetWorkspaceById

type GetWorkspaceByIdUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
}

type GetWorkspaceByIdQuery struct {
	OrgId       int32
	WorkspaceId int32
}

// GetWorkspaceVisibilityById

type GetWorkspaceVisibilityByIdQuery struct {
	OrgId       int32
	WorkspaceId int32
}

// CreateWorkspace

type CreateWorkspaceUriParams struct {
	OrgId int32 `uri:"orgId" binding:"required"`
}

type CreateWorkspaceCommand struct {
	OrgId     int32
	CreatedBy string

	ScopeId     string                `json:"sid"`
	Name        string                `json:"name"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility"`
}

// UpdateWorkspaceById

type UpdateWorkspaceByIdUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
}

type UpdateWorkspaceByIdCommand struct {
	OrgId       int32
	WorkspaceId int32
	UpdatedBy   string

	ScopeId     string                `json:"sid,omitempty"`
	Name        string                `json:"name,omitempty"`
	Description string                `json:"description,omitempty"`
	Visibility  model.ScopeVisibility `json:"visibility,omitempty"`
}

// DeleteWorkspaceById

type DeleteWorkspaceByIdUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
}

type DeleteWorkspaceByIdCommand struct {
	OrgId       int32
	WorkspaceId int32
}

// GetProjectList

type GetProjectListUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
}

type GetProjectListQuery struct {
	OrgId       int32
	WorkspaceId int32
	Visibility  model.ScopeVisibility
}

// GetMemberList

type GetMemberListUriParams struct {
	OrgId       int32 `uri:"orgId" binding:"required"`
	WorkspaceId int32 `uri:"workspaceId" binding:"required"`
}

type GetMemberListQuery struct {
	OrgId       int32
	WorkspaceId int32
}

// GetUserAssignmentList

type GetUserAssignmentListQuery struct {
	OrgId       int32
	WorkspaceId int32
}
