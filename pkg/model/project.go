package model

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
	"github.com/quarkloop/quarkloop/service/v1/system"
)

type Project struct {
	// id
	Id               int32  `json:"id"`
	ScopeId          string `json:"sid"`
	WorkspaceId      int32  `json:"workspaceId"`
	WorkspaceScopeId string `json:"workspaceScopeId"`
	OrgId            int32  `json:"orgId"`
	OrgScopeId       string `json:"orgScopeId"`

	// data
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Visibility  ScopeVisibility             `json:"visibility"`
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
	var updatedAt *timestamppb.Timestamp = nil
	if p.UpdatedAt != nil {
		updatedAt = timestamppb.New(*p.UpdatedAt)
	}

	var updatedBy string = ""
	if p.UpdatedBy != nil {
		updatedBy = *p.UpdatedBy
	}

	project := &system.Project{
		Id:               p.Id,
		ScopeId:          p.ScopeId,
		OrgId:            p.OrgId,
		OrgScopeId:       p.OrgScopeId,
		WorkspaceId:      p.WorkspaceId,
		WorkspaceScopeId: p.WorkspaceScopeId,
		Name:             p.Name,
		Description:      p.Description,
		Visibility:       int32(p.Visibility),
		Path:             p.Path,
		CreatedAt:        timestamppb.New(p.CreatedAt),
		CreatedBy:        p.CreatedBy,
		UpdatedAt:        updatedAt,
		UpdatedBy:        updatedBy,
	}

	return project
}
