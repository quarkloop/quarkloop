package workspace_impl

import (
	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/user"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/service/workspace/store"
)

type workspaceService struct {
	store store.WorkspaceStore
}

func NewWorkspaceService(ds store.WorkspaceStore) workspace.Service {
	return &workspaceService{
		store: ds,
	}
}

func (s *workspaceService) GetWorkspaceList(ctx *gin.Context, query *workspace.GetWorkspaceListQuery) ([]*workspace.Workspace, error) {
	workspaceList, err := s.store.GetWorkspaceList(ctx, query)
	if err != nil {
		return nil, err
	}

	for i := range workspaceList {
		workspace := workspaceList[i]
		workspace.GeneratePath()
	}
	return workspaceList, nil
}

func (s *workspaceService) GetWorkspaceById(ctx *gin.Context, query *workspace.GetWorkspaceByIdQuery) (*workspace.Workspace, error) {
	ws, err := s.store.GetWorkspaceById(ctx, query)
	if err != nil {
		return nil, err
	}

	ws.GeneratePath()
	return ws, nil
}

func (s *workspaceService) CreateWorkspace(ctx *gin.Context, cmd *workspace.CreateWorkspaceCommand) (*workspace.Workspace, error) {
	ws, err := s.store.CreateWorkspace(ctx, cmd)
	if err != nil {
		return nil, err
	}

	ws.GeneratePath()
	return ws, nil
}

func (s *workspaceService) UpdateWorkspaceById(ctx *gin.Context, cmd *workspace.UpdateWorkspaceByIdCommand) error {
	return s.store.UpdateWorkspaceById(ctx, cmd)
}

func (s *workspaceService) DeleteWorkspaceById(ctx *gin.Context, cmd *workspace.DeleteWorkspaceByIdCommand) error {
	return s.store.DeleteWorkspaceById(ctx, cmd)
}

func (s *workspaceService) GetProjectList(ctx *gin.Context, query *workspace.GetProjectListQuery) ([]*project.Project, error) {
	projectList, err := s.store.GetProjectList(ctx, query)
	if err != nil {
		return nil, err
	}

	for i := range projectList {
		p := projectList[i]
		p.GeneratePath()
	}

	return projectList, nil
}

func (s *workspaceService) GetUserAssignmentList(ctx *gin.Context, query *workspace.GetUserAssignmentListQuery) ([]*user.UserAssignment, error) {
	uaList, err := s.store.GetUserAssignmentList(ctx, query)
	if err != nil {
		return nil, err
	}

	return uaList, nil
}
