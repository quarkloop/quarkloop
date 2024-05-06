package org

import (
	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetOrgListQuery struct {
	OrgIdList  []int32
	Visibility model.ScopeVisibility
}

type GetOrgVisibilityByIdQuery struct {
	OrgId int32
}

type GetUserAssignmentListQuery struct {
	OrgId int32 `uri:"orgId" binding:"required"`
}
