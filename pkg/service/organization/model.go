package org

import (
	"context"
	"fmt"
	"time"
)

type Organization struct {
	// id
	Id       int    `json:"id" form:"id"`
	ScopedId string `json:"sid"`

	// data
	Name        string `json:"name,omitempty" form:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AccessType  *int   `json:"accessType,omitempty" form:"accessType,omitempty"`
	Path        string `json:"path,omitempty"`

	// history
	CreatedAt time.Time  `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy *string    `json:"updatedBy,omitempty"`
}

func (o *Organization) GeneratePath() {
	o.Path = fmt.Sprintf("/org/%s", o.ScopedId)
}

type GetOrganizationListParams struct {
	Context context.Context
}

type GetOrganizationByIdParams struct {
	Context context.Context
	OrgId   int
}

type GetOrganizationParams struct {
	Context      context.Context
	Organization Organization
}

type CreateOrganizationParams struct {
	Context      context.Context
	Organization Organization
}

type UpdateOrganizationByIdParams struct {
	Context      context.Context
	OrgId        int
	Organization Organization
}

type DeleteOrganizationByIdParams struct {
	Context context.Context
	OrgId   int
}
