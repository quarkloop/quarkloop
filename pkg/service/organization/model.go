package org

import (
	"errors"
	"fmt"
	"time"

	"github.com/quarkloop/quarkloop/pkg/model"
)

var (
	ErrOrgNotFound = errors.New("org not found")
)

type Organization struct {
	// id
	Id       int    `json:"id" form:"id"`
	ScopedId string `json:"sid"`

	// data
	Name        string                 `json:"name,omitempty" form:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Visibility  *model.ScopeVisibility `json:"visibility,omitempty" form:"visibility,omitempty"`
	Path        string                 `json:"path,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

func (o *Organization) GeneratePath() {
	o.Path = fmt.Sprintf("/org/%s", o.ScopedId)
}

type GetOrganizationByIdParams struct {
	OrgId int
}

type GetOrganizationParams struct {
	Organization Organization
}

type CreateOrganizationParams struct {
	Organization Organization
}

type UpdateOrganizationByIdParams struct {
	OrgId        int
	Organization Organization
}

type DeleteOrganizationByIdParams struct {
	OrgId int
}
