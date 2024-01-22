package project

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quarkloop/quarkloop/pkg/api"
	"github.com/quarkloop/quarkloop/pkg/contextdata"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/project"
	"github.com/quarkloop/quarkloop/pkg/service/quota"
	"github.com/quarkloop/quarkloop/pkg/service/table_branch"
)

func (s *ProjectApi) createProject(ctx *gin.Context, cmd *project.CreateProjectCommand) api.Response {
	user := contextdata.GetUser(ctx)

	// check permissions
	access, err := s.evaluateCreatePermission(ctx, accesscontrol.ActionProjectCreate, cmd.OrgId, cmd.WorkspaceId)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}
	if !access {
		// unauthorized user (permission denied) => return project not found error
		return api.Error(http.StatusNotFound, project.ErrProjectNotFound) // TODO: change status code and error
	}

	// check quotas
	quotaQuery := &quota.CheckCreateProjectQuotaQuery{OrgId: cmd.OrgId}
	if err := s.quotaService.CheckCreateProjectQuota(ctx, quotaQuery); err != nil {
		return api.Error(http.StatusInternalServerError, err) // TODO: change status
	}

	p, err := s.projectService.CreateProject(ctx, cmd)
	if err != nil {
		return api.Error(http.StatusInternalServerError, err)
	}

	mainBranch, err := s.branchService.CreateTableBranch(ctx, &table_branch.CreateTableBranchParams{
		ProjectId: p.Id,
		Branch: &table_branch.TableBranch{
			Name:        "main",
			Type:        "main",
			Default:     true,
			Description: "main branch",
			CreatedBy:   user.Name, // TODO
		},
	})
	if err != nil {
		joinedErr := errors.Join(err)

		deleteErr := s.projectService.DeleteProjectById(ctx, &project.DeleteProjectByIdCommand{
			OrgId:       cmd.OrgId,
			WorkspaceId: cmd.WorkspaceId,
			ProjectId:   p.Id,
		})
		if deleteErr != nil {
			joinedErr = errors.Join(deleteErr)
		}

		return api.Error(http.StatusInternalServerError, joinedErr)
	}

	submissionBranch, err := s.branchService.CreateTableBranch(ctx, &table_branch.CreateTableBranchParams{
		ProjectId: p.Id,
		Branch: &table_branch.TableBranch{
			Name:        "submission",
			Type:        "submission",
			Default:     false,
			Description: "submission branch",
			CreatedBy:   "user", // TODO
		},
	})
	if err != nil {
		joinedErr := errors.Join(err)

		deleteErr := s.projectService.DeleteProjectById(ctx, &project.DeleteProjectByIdCommand{
			OrgId:       cmd.OrgId,
			WorkspaceId: cmd.WorkspaceId,
			ProjectId:   p.Id,
		})
		if deleteErr != nil {
			joinedErr = errors.Join(deleteErr)
		}
		deleteErr = s.branchService.DeleteTableBranchById(ctx, &table_branch.DeleteTableBranchByIdParams{
			ProjectId: p.Id,
			BranchId:  mainBranch.Id,
		})
		if deleteErr != nil {
			joinedErr = errors.Join(deleteErr)
		}

		return api.Error(http.StatusInternalServerError, joinedErr)
	}

	p.Branches = append(p.Branches, mainBranch, submissionBranch)

	return api.Success(http.StatusOK, p)
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

	if err := s.projectService.UpdateProjectById(ctx, cmd); err != nil {
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

	if err := s.projectService.DeleteProjectById(ctx, cmd); err != nil {
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
