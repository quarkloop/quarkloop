package contextdata

import "context"

type Scope interface {
	OrgId() int
	WorkspaceId() int
	ProjectId() int
}

func GetScope(ctx context.Context) Scope {
	scope, ok := ctx.Value(scopeKey).(*scopeData)
	if !ok || scope == nil {
		panic("scope must be available")
	}

	return scope
}

func SetScope(ctx context.Context, params *ScopeParams) context.Context {
	scope := &scopeData{
		orgId:       params.OrgId,
		workspaceId: params.WorkspaceId,
		projectId:   params.ProjectId,
	}
	return context.WithValue(ctx, scopeKey, scope)
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
