package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	grpc "github.com/quarkloop/quarkloop/service/v1/system/project"
)

func (s *ProjectApi) createProject(ctx *gin.Context, cmd *project.CreateProjectCommand) api.Response {
	// check permissions
	access, err := s.evaluateCreatePermission(ctx, accesscontrol.ActionProjectCreate, cmd.OrgId, cmd.WorkspaceId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return project not found error
		return api.Error(http.StatusForbidden, accesscontrol.ErrPermissionDenied) // TODO: change status code and error
	}

	// check quotas
	quotaQuery := &quota.CheckCreateProjectQuotaQuery{OrgId: cmd.OrgId}
	if err := s.quotaService.CheckCreateProjectQuota(ctx, quotaQuery); err != nil {
		return api.Error(http.StatusTooManyRequests, err)
	}

	p, err := s.projectService.CreateProject(ctx, &grpc.CreateProjectCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		CreatedBy:   cmd.CreatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  int32(cmd.Visibility),
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.AlreadyExists:
				return api.Error(http.StatusConflict, err)
			case codes.Internal:
				return api.Error(http.StatusInternalServerError, err)
			case codes.InvalidArgument:
				return api.Error(http.StatusBadRequest, err)
			}
		}
		return api.Error(http.StatusInternalServerError, err)
	}

	// mainBranch, err := s.branchService.CreateTableBranch(ctx, &table_branch.CreateTableBranchParams{
	// 	ProjectId: p.Project.Id,
	// 	Branch: &table_branch.TableBranch{
	// 		Name:        "main",
	// 		Type:        "main",
	// 		Default:     true,
	// 		Description: "main branch",
	// 		CreatedBy:   "admin", // TODO
	// 	},
	// })
	// if err != nil {
	// 	joinedErr := errors.Join(err)

	// 	_, deleteErr := s.projectService.DeleteProjectById(ctx, &grpc.DeleteProjectByIdCommand{
	// 		OrgId:       cmd.OrgId,
	// 		WorkspaceId: cmd.WorkspaceId,
	// 		ProjectId:   p.Project.Id,
	// 	})
	// 	if deleteErr != nil {
	// 		joinedErr = errors.Join(deleteErr)
	// 	}

	// 	return api.Error(http.StatusInternalServerError, joinedErr)
	// }

	// submissionBranch, err := s.branchService.CreateTableBranch(ctx, &table_branch.CreateTableBranchParams{
	// 	ProjectId: p.Project.Id,
	// 	Branch: &table_branch.TableBranch{
	// 		Name:        "submission",
	// 		Type:        "submission",
	// 		Default:     false,
	// 		Description: "submission branch",
	// 		CreatedBy:   "user", // TODO
	// 	},
	// })
	// if err != nil {
	// 	joinedErr := errors.Join(err)

	// 	_, deleteErr := s.projectService.DeleteProjectById(ctx, &grpc.DeleteProjectByIdCommand{
	// 		OrgId:       cmd.OrgId,
	// 		WorkspaceId: cmd.WorkspaceId,
	// 		ProjectId:   p.Project.Id,
	// 	})
	// 	if deleteErr != nil {
	// 		joinedErr = errors.Join(deleteErr)
	// 	}
	// 	deleteErr = s.branchService.DeleteTableBranchById(ctx, &table_branch.DeleteTableBranchByIdParams{
	// 		ProjectId: p.Project.Id,
	// 		BranchId:  mainBranch.Id,
	// 	})
	// 	if deleteErr != nil {
	// 		joinedErr = errors.Join(deleteErr)
	// 	}

	// 	return api.Error(http.StatusInternalServerError, joinedErr)
	// }

	//p.Branches = append(p.Project.Branches, mainBranch, submissionBranch)

	relationCmd := &accesscontrol.MakeParentResourceCommand{
		ParentResource:   "workspace",
		ParentResourceId: cmd.WorkspaceId,
		ChildResource:    "project",
		ChildResourceId:  p.Project.Id,
	}
	err = s.aclService.MakeParentResource(ctx, relationCmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusCreated, p.GetProject())
}

func (s *ProjectApi) updateProjectById(ctx *gin.Context, cmd *project.UpdateProjectByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, accesscontrol.ActionProjectUpdate, cmd.OrgId, cmd.WorkspaceId, cmd.ProjectId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return project not found error
		return api.Error(http.StatusNotFound, project.ErrProjectNotFound) // TODO: change status code and error
	}

	_, err = s.projectService.UpdateProjectById(ctx, &grpc.UpdateProjectByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		ProjectId:   cmd.ProjectId,
		UpdatedBy:   cmd.UpdatedBy,
		ScopeId:     cmd.ScopeId,
		Name:        cmd.Name,
		Description: cmd.Description,
		Visibility:  int32(cmd.Visibility),
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusOK, nil)
}

func (s *ProjectApi) deleteProjectById(ctx *gin.Context, cmd *project.DeleteProjectByIdCommand) api.Response {
	// check permissions
	access, err := s.evaluatePermission(ctx, accesscontrol.ActionProjectDelete, cmd.OrgId, cmd.WorkspaceId, cmd.ProjectId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return project not found error
		return api.Error(http.StatusNotFound, project.ErrProjectNotFound) // TODO: change status code and error
	}

	_, err = s.projectService.DeleteProjectById(ctx, &grpc.DeleteProjectByIdCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		ProjectId:   cmd.ProjectId,
	})
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	aclCommand := &accesscontrol.RevokeUserAccessCommand{
		OrgId:       cmd.OrgId,
		WorkspaceId: cmd.WorkspaceId,
		ProjectId:   cmd.ProjectId,
	}
	err = s.aclService.RevokeUserAccess(ctx, aclCommand)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	return api.Success(http.StatusNoContent, nil)
}

func (s *ProjectApi) evaluateCreatePermission(ctx *gin.Context, permission string, orgId, workspaceId int32) (bool, error) {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateQuery{
		Permission:  permission,
		UserId:      user.GetId(),
		OrgId:       orgId,
		WorkspaceId: workspaceId,
	}

	return s.aclService.EvaluateUserAccess(ctx, query)
}

func (s *ProjectApi) evaluatePermission(ctx *gin.Context, permission string, orgId, workspaceId, projectId int32) (bool, error) {
	user := contextdata.GetUser(ctx)
	query := &accesscontrol.EvaluateQuery{
		Permission:  permission,
		UserId:      user.GetId(),
		OrgId:       orgId,
		WorkspaceId: workspaceId,
		ProjectId:   projectId,
	}

	return s.aclService.EvaluateUserAccess(ctx, query)
}
