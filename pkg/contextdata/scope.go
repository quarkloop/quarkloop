package contextdata

import (
	"github.com/gin-gonic/gin"
)

type Scope interface {
	OrgId() int
	WorkspaceId() int
	ProjectId() int
}

func GetScope(ctx *gin.Context) Scope {
	val, exists := ctx.Get(scopeKey)
	if !exists || val == nil {
		panic("scope must be available")
	}

	scope, ok := val.(*scopeData)
	if !ok {
		panic("*scopeData type assertion failed")
	}

	return scope
}

func SetScope(ctx *gin.Context, params *ScopeParams) {
	scope := &scopeData{
		orgId:       params.OrgId,
		workspaceId: params.WorkspaceId,
		projectId:   params.ProjectId,
	}

	ctx.Set(userKey, scope)
}

type ScopeParams struct {
	OrgId, WorkspaceId, ProjectId int
}

type scopeData struct {
	orgId, workspaceId, projectId int
}

func (s *scopeData) OrgId() int {
	return s.orgId
}

func (s *scopeData) WorkspaceId() int {
	return s.workspaceId
}

func (s *scopeData) ProjectId() int {
	return s.projectId
}
