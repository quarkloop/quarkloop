package model

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/quarkloop/quarkloop/service/v1/system"
)

type Org struct {
	// id
	Id      int32  `json:"id"`
	ScopeId string `json:"sid"`

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

func (o *Org) GeneratePath() {
	o.Path = fmt.Sprintf("/org/%s", o.ScopeId)
}

func (org *Org) Proto() *system.Org {
	var updatedAt *timestamppb.Timestamp = nil
	if org.UpdatedAt != nil {
		updatedAt = timestamppb.New(*org.UpdatedAt)
	}

	var updatedBy string = ""
	if org.UpdatedBy != nil {
		updatedBy = *org.UpdatedBy
	}

	o := &system.Org{
		Id:          org.Id,
		ScopeId:     org.ScopeId,
		Name:        org.Name,
		Description: org.Description,
		Visibility:  int32(org.Visibility),
		Path:        org.Path,
		CreatedAt:   timestamppb.New(org.CreatedAt),
		CreatedBy:   org.CreatedBy,
		UpdatedAt:   updatedAt,
		UpdatedBy:   updatedBy,
	}

	return o
}
