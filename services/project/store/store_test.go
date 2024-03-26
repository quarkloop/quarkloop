package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/test"
	orgStore "github.com/quarkloop/quarkloop/services/org/store"
	projectErrors "github.com/quarkloop/quarkloop/services/project/errors"
	workspaceStore "github.com/quarkloop/quarkloop/services/workspace/store"
)

var (
	ctx         context.Context
	conn        *pgx.Conn
	orgId       int64
	workspaceId int64
)

const projectCount = 10

func init() {
	ctx, conn = test.InitTestSystemDB()
}

func TestMutationTruncateTables(t *testing.T) {
	t.Run("truncate tables", func(t *testing.T) {
		err := test.TruncateSystemDBTables(ctx, conn)
		require.NoError(t, err)
	})

	t.Run("get project list return empty after truncating tables", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)
		require.NoError(t, err)
		require.Zero(t, len(prList))
		require.Equal(t, 0, len(prList))
	})
}

func TestPrepare(t *testing.T) {
	{
		store := orgStore.NewOrgStore(conn)
		t.Run("create single org for whole test", func(t *testing.T) {
			cmd := &orgStore.CreateOrgCommand{
				ScopeId:     "quarkloop",
				Name:        "Quarkloop",
				Description: "Quarkloop Corporation",
				CreatedBy:   "admin",
				Visibility:  model.PublicVisibility,
			}
			org, err := store.CreateOrg(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, org)

			orgId = org.Id
		})
	}

	{
		store := workspaceStore.NewWorkspaceStore(conn)
		t.Run("create single workspace for whole test", func(t *testing.T) {
			cmd := &workspaceStore.CreateWorkspaceCommand{
				OrgId:       orgId,
				ScopeId:     "it",
				Name:        "IT department",
				Description: "Quarkloop Corporation IT department",
				CreatedBy:   "admin",
				Visibility:  model.PublicVisibility,
			}
			ws, err := store.CreateWorkspace(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, ws)

			workspaceId = ws.Id
		})
	}
}

func TestMutationCreateProject(t *testing.T) {
	store := NewProjectStore(conn)

	t.Run("create project with duplicate scopeId", func(t *testing.T) {
		var projectId int64 = 0
		cmd := &CreateProjectCommand{
			OrgId:       orgId,
			WorkspaceId: workspaceId,
			ScopeId:     "it",
			Name:        "IT",
			Description: "IT department",
			CreatedBy:   "admin",
			Visibility:  model.PublicVisibility,
		}
		{
			// first project
			pr, err := store.CreateProject(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, pr)
			require.NotEmpty(t, pr.ScopeId)
			require.Equal(t, cmd.ScopeId, pr.ScopeId)

			projectId = pr.Id
		}
		{
			// second project (duplicate)
			prDuplicate, err := store.CreateProject(ctx, cmd)

			require.Nil(t, prDuplicate)
			require.Error(t, err)
			require.Exactly(t, projectErrors.ErrProjectAlreadyExists, err)
			require.Equal(t, "project with same scopeId already exists", err.Error())
		}
		{
			// clean up
			cmd := &DeleteProjectByIdCommand{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: projectId}
			deleteErr := store.DeleteProjectById(ctx, cmd)
			require.NoError(t, deleteErr)

			query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: projectId}
			pr, err := store.GetProjectById(ctx, query)
			require.Nil(t, pr)
			require.Error(t, err)
			require.Exactly(t, projectErrors.ErrProjectNotFound, err)
			require.Equal(t, "project not found", err.Error())
		}
	})

	t.Run("create project without scopeId", func(t *testing.T) {
		for i := 0; i < projectCount; i++ {
			cmd := &CreateProjectCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			pr, err := store.CreateProject(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, pr)
			require.NotEmpty(t, pr.ScopeId)
			require.NotEmpty(t, pr.Name)
			require.NotEmpty(t, pr.Description)
			require.NotEmpty(t, pr.CreatedBy)
			require.NotZero(t, pr.Visibility)
			require.Equal(t, cmd.Name, pr.Name)
			require.Equal(t, cmd.Description, pr.Description)
			require.Equal(t, cmd.Visibility, pr.Visibility)
			require.Equal(t, cmd.CreatedBy, pr.CreatedBy)

			{
				// clean up
				cmd := &DeleteProjectByIdCommand{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
				deleteErr := store.DeleteProjectById(ctx, cmd)
				require.NoError(t, deleteErr)

				query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
				pr, err := store.GetProjectById(ctx, query)
				require.Nil(t, pr)
				require.Error(t, err)
				require.Exactly(t, projectErrors.ErrProjectNotFound, err)
				require.Equal(t, "project not found", err.Error())
			}
		}
	})

	t.Run("create project with scopeId", func(t *testing.T) {
		for i := 0; i < projectCount; i++ {
			cmd := &CreateProjectCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				ScopeId:     fmt.Sprintf("quarkloop_%d", i),
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			pr, err := store.CreateProject(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, pr)
			require.NotEmpty(t, pr.ScopeId)
			require.NotEmpty(t, pr.Name)
			require.NotEmpty(t, pr.Description)
			require.NotEmpty(t, pr.CreatedBy)
			require.NotZero(t, pr.Visibility)
			require.Equal(t, cmd.ScopeId, pr.ScopeId)
			require.Equal(t, cmd.Name, pr.Name)
			require.Equal(t, cmd.Description, pr.Description)
			require.Equal(t, cmd.Visibility, pr.Visibility)
			require.Equal(t, cmd.CreatedBy, pr.CreatedBy)
		}
	})

	t.Run("get project list return full", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)

		require.NoError(t, err)
		require.NotZero(t, len(prList))
		require.Equal(t, projectCount, len(prList))
	})
}

func TestQueryGetProjectAfterCreate(t *testing.T) {
	store := NewProjectStore(conn)

	t.Run("get project by id after creation", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)
		require.NoError(t, err)

		for idx, pr := range prList {
			query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
			project, err := store.GetProjectById(ctx, query)

			require.NoError(t, err)
			require.NotNil(t, project)
			require.NotEmpty(t, project.ScopeId)
			require.NotEmpty(t, project.Name)
			require.NotEmpty(t, project.Description)
			require.NotEmpty(t, project.CreatedBy)
			require.NotZero(t, project.Visibility)
			require.Equal(t, fmt.Sprintf("quarkloop_%d", idx), project.ScopeId)
			require.Equal(t, fmt.Sprintf("Quarkloop_%d", idx), project.Name)
			require.Equal(t, fmt.Sprintf("Quarkloop Corporation #%d", idx), project.Description)
			require.Equal(t, fmt.Sprintf("admin_%d", idx), project.CreatedBy)
			require.Equal(t, model.PublicVisibility, project.Visibility)
		}
	})

	t.Run("get project by wrong id", func(t *testing.T) {
		query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: 9999999}
		pr, err := store.GetProjectById(ctx, query)

		require.Nil(t, pr)
		require.Error(t, err)
		require.Exactly(t, projectErrors.ErrProjectNotFound, err)
		require.Equal(t, "project not found", err.Error())
	})

	t.Run("get project visibility by id after creation", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)
		require.NoError(t, err)

		for _, pr := range prList {
			query := &GetProjectVisibilityByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
			visibility, err := store.GetProjectVisibilityById(ctx, query)

			require.NoError(t, err)
			require.NotZero(t, visibility)
			require.Equal(t, model.PublicVisibility, visibility)
		}
	})
}

func TestMutationUpdateProject(t *testing.T) {
	store := NewProjectStore(conn)

	t.Run("update project with duplicate scope id", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)
		require.NoError(t, err)

		{
			// original scope id
			cmd := &UpdateProjectByIdCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				ProjectId:   prList[0].Id,
				ScopeId:     "quarkloop_updated_scopeid",
			}
			err := store.UpdateProjectById(ctx, cmd)

			require.NoError(t, err)
		}
		{
			// duplicate scope id
			cmd := &UpdateProjectByIdCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				ProjectId:   prList[len(prList)-1].Id,
				ScopeId:     "quarkloop_updated_scopeid",
			}
			err := store.UpdateProjectById(ctx, cmd)

			require.Error(t, err)
			require.Exactly(t, projectErrors.ErrProjectAlreadyExists, err)
			require.Equal(t, "project with same scopeId already exists", err.Error())
		}
	})

	t.Run("partial project update", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)
		require.NoError(t, err)

		// name
		for idx, pr := range prList {
			name := fmt.Sprintf("Quarkloop_Updated_%d", idx)
			cmd := &UpdateProjectByIdCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				ProjectId:   pr.Id,
				Name:        name,
			}
			err := store.UpdateProjectById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
				project, err := store.GetProjectById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, project)
				require.Equal(t, name, project.Name)
				require.NotEmpty(t, project.ScopeId)
				require.NotEmpty(t, project.Name)
				require.NotEmpty(t, project.Description)
				require.NotEmpty(t, project.CreatedBy)
				require.NotZero(t, project.Visibility)
			}
		}
		// description
		for idx, pr := range prList {
			description := fmt.Sprintf("Quarkloop_Description_Updated_%d", idx)
			cmd := &UpdateProjectByIdCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				ProjectId:   pr.Id,
				Description: description,
			}
			err := store.UpdateProjectById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
				project, err := store.GetProjectById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, project)
				require.Equal(t, description, project.Description)
				require.NotEmpty(t, project.ScopeId)
				require.NotEmpty(t, project.Name)
				require.NotEmpty(t, project.Description)
				require.NotEmpty(t, project.CreatedBy)
				require.NotZero(t, project.Visibility)
			}
		}
		// visibility
		for _, pr := range prList {
			visibility := model.PrivateVisibility
			cmd := &UpdateProjectByIdCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				ProjectId:   pr.Id,
				Visibility:  visibility,
			}
			err := store.UpdateProjectById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
				project, err := store.GetProjectById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, project)
				require.Equal(t, visibility, project.Visibility)
				require.NotEmpty(t, project.ScopeId)
				require.NotEmpty(t, project.Name)
				require.NotEmpty(t, project.Description)
				require.NotEmpty(t, project.CreatedBy)
				require.NotZero(t, project.Visibility)
			}
		}
		// updatedBy
		for idx, pr := range prList {
			updatedBy := fmt.Sprintf("Quarkloop_Admin2_Updated_%d", idx)
			cmd := &UpdateProjectByIdCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				ProjectId:   pr.Id,
				UpdatedBy:   updatedBy,
			}
			err := store.UpdateProjectById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
				project, err := store.GetProjectById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, project)
				require.Equal(t, updatedBy, *project.UpdatedBy)
				require.NotEmpty(t, project.ScopeId)
				require.NotEmpty(t, project.Name)
				require.NotEmpty(t, project.Description)
				require.NotEmpty(t, project.CreatedBy)
				require.NotZero(t, project.Visibility)
			}
		}
	})

	t.Run("update project with all fields", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)
		require.NoError(t, err)

		for idx, pr := range prList {
			cmd := &UpdateProjectByIdCommand{
				OrgId:       orgId,
				WorkspaceId: workspaceId,
				ProjectId:   pr.Id,
				ScopeId:     fmt.Sprintf("quarkloop_new_update_%d", idx),
				Name:        fmt.Sprintf("Quarkloop_New_Update_%d", idx),
				Description: fmt.Sprintf("Quarkloop Corporation Updated With #%d", idx),
				UpdatedBy:   fmt.Sprintf("admin_1_updated_%d", idx),
				Visibility:  model.PrivateVisibility,
			}
			err := store.UpdateProjectById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &GetProjectByIdQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
				project, err := store.GetProjectById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, project)
				require.Equal(t, cmd.ScopeId, project.ScopeId)
				require.Equal(t, cmd.Name, project.Name)
				require.Equal(t, cmd.Description, project.Description)
				require.Equal(t, cmd.Visibility, project.Visibility)
				require.Equal(t, cmd.UpdatedBy, *project.UpdatedBy)
				require.NotEmpty(t, project.ScopeId)
				require.NotEmpty(t, project.Name)
				require.NotEmpty(t, project.Description)
				require.NotZero(t, project.Visibility)
				require.NotNil(t, project.UpdatedBy)
			}
		}
	})
}

// func TestQueryWorkspaceRelations(t *testing.T) {
// 	store := NewProjectStore(conn)

// 	t.Run("get project's user assignment list", func(t *testing.T) {
// 		prList, err := test.GetFullProjectList(ctx, conn)
// 		require.NoError(t, err)

// 		for _, pr := range prList {
// 			query := &GetUserAssignmentListQuery{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
// 			list, err := store.GetUserAssignmentList(ctx, query)

// 			require.NoError(t, err)
// 			require.Empty(t, list)
// 			require.Equal(t, 0, len(list))
// 		}
// 	})
// }

func TestMutationDeleteProject(t *testing.T) {
	store := NewProjectStore(conn)

	t.Run("delete all workspaces by id", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)
		require.NoError(t, err)

		for _, pr := range prList {
			cmd := &DeleteProjectByIdCommand{OrgId: orgId, WorkspaceId: workspaceId, ProjectId: pr.Id}
			err := store.DeleteProjectById(ctx, cmd)
			require.NoError(t, err)
		}
	})

	t.Run("get project list should return empty", func(t *testing.T) {
		prList, err := test.GetFullProjectList(ctx, conn)

		require.NoError(t, err)
		require.Zero(t, len(prList))
		require.Equal(t, 0, len(prList))
	})
}

func TestCleanup(t *testing.T) {
	{
		store := workspaceStore.NewWorkspaceStore(conn)
		t.Run("delete workspace by id", func(t *testing.T) {
			cmd := &workspaceStore.DeleteWorkspaceByIdCommand{OrgId: orgId, WorkspaceId: workspaceId}
			err := store.DeleteWorkspaceById(ctx, cmd)
			require.NoError(t, err)
		})
	}
	{
		store := orgStore.NewOrgStore(conn)
		t.Run("delete org by id", func(t *testing.T) {
			err := store.DeleteOrgById(ctx, &orgStore.DeleteOrgByIdCommand{OrgId: orgId})
			require.NoError(t, err)
		})
	}
}
