package project

import (
	"errors"
	"fmt"
	"time"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
	"github.com/quarkloop/quarkloop/service/system"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrProjectNotFound      = errors.New("project not found")
	ErrProjectAlreadyExists = errors.New("project with same scopeId already exists")
)

type Project struct {
	// id
	Id               int32  `json:"id"`
	ScopeId          string `json:"sid"`
	WorkspaceId      int32  `json:"workspaceId"`
	WorkspaceScopeId string `json:"workspaceScopeId"`
	OrgId            int32  `json:"orgId"`
	OrgScopeId       string `json:"orgScopeId"`

	// project
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Visibility  model.ScopeVisibility       `json:"visibility"`
	Path        string                      `json:"path"`
	Branches    []*table_branch.TableBranch `json:"branches"`

	// history
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy *string    `json:"updatedBy"`
}

func (p *Project) GeneratePath() {
	p.Path = fmt.Sprintf("/org/%s/%s/%s", p.OrgScopeId, p.WorkspaceScopeId, p.ScopeId)
}

func (p *Project) Proto() *system.Project {
	project := &system.Project{
		Id:          p.Id,
		ScopeId:     p.ScopeId,
		Name:        p.Name,
		Description: p.Description,
		Visibility:  int32(p.Visibility),
		Path:        p.Path,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(*p.UpdatedAt),
		CreatedBy:   p.CreatedBy,
		UpdatedBy:   *p.UpdatedBy,
	}

	return project
}

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
	UserId     int32
	Visibility model.ScopeVisibility
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
