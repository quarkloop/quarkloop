package model

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/quarkloop/quarkloop/service/v1/system"
)

type Workspace struct {
	// id
	Id         int32  `json:"id"`
	ScopeId    string `json:"sid"`
	OrgId      int32  `json:"orgId"`
	OrgScopeId string `json:"orgScopeId"`

	// data
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Visibility  ScopeVisibility `json:"visibility"`
	Path        string          `json:"path"`

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
	var updatedAt *timestamppb.Timestamp = nil
	if ws.UpdatedAt != nil {
		updatedAt = timestamppb.New(*ws.UpdatedAt)
	}

	var updatedBy string = ""
	if ws.UpdatedBy != nil {
		updatedBy = *ws.UpdatedBy
	}

	workspace := &system.Workspace{
		Id:          ws.Id,
		ScopeId:     ws.ScopeId,
		OrgId:       ws.OrgId,
		OrgScopeId:  ws.OrgScopeId,
		Name:        ws.Name,
		Description: ws.Description,
		Visibility:  int32(ws.Visibility),
		Path:        ws.Path,
		CreatedAt:   timestamppb.New(ws.CreatedAt),
		CreatedBy:   ws.CreatedBy,
		UpdatedAt:   updatedAt,
		UpdatedBy:   updatedBy,
	}

	return workspace
}
