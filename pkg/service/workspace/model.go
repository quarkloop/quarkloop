package workspace

import (
	"errors"
	"fmt"
	"time"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/service/system"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrWorkspaceNotFound      = errors.New("workspace not found")
	ErrWorkspaceAlreadyExists = errors.New("workspace with same scopeId already exists")
)

type Workspace struct {
	// id
	Id         int32  `json:"id"`
	ScopeId    string `json:"sid"`
	OrgId      int32  `json:"orgId"`
	OrgScopeId string `json:"orgScopeId"`

	// data
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Visibility  model.ScopeVisibility `json:"visibility"`
	Path        string                `json:"path"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

func (ws *Workspace) GeneratePath() {
	ws.Path = fmt.Sprintf("/org/%s/%s", ws.OrgScopeId, ws.ScopeId)
}

func (ws *Workspace) Proto() *system.Workspace {
	workspace := &system.Workspace{
		Id:          ws.Id,
		ScopeId:     ws.ScopeId,
		Name:        ws.Name,
		Description: ws.Description,
		Visibility:  int32(ws.Visibility),
		Path:        ws.Path,
		CreatedAt:   timestamppb.New(ws.CreatedAt),
		UpdatedAt:   timestamppb.New(*ws.UpdatedAt),
		CreatedBy:   ws.CreatedBy,
		UpdatedBy:   *ws.UpdatedBy,
	}

	return workspace
}

// GetWorkspaceList

type GetWorkspaceListQuery struct {
	UserId     int32
	Visibility model.ScopeVisibility
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
