package store_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	orgStore "github.com/quarkloop/quarkloop/pkg/service/org/store"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	"github.com/quarkloop/quarkloop/pkg/service/workspace/store"
	"github.com/quarkloop/quarkloop/pkg/test"
)

var (
	ctx   context.Context
	conn  *pgx.Conn
	orgId int32
)

const workspaceCount = 10

func init() {
	ctx, conn = test.InitTestSystemDB()
}

func TestMutationTruncateTables(t *testing.T) {
	t.Run("truncate tables", func(t *testing.T) {
		err := test.TruncateSystemDBTables(ctx, conn)
		require.NoError(t, err)
	})

	t.Run("get workspace list return empty after truncating tables", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)
		require.Zero(t, len(wsList))
		require.Equal(t, 0, len(wsList))
	})
}

func TestPrepare(t *testing.T) {
	orgStore := orgStore.NewOrgStore(conn)

	t.Run("create single org for whole test", func(t *testing.T) {
		cmd := &org.CreateOrgCommand{
			ScopeId:     "quarkloop",
			Name:        "Quarkloop",
			Description: "Quarkloop Corporation",
			CreatedBy:   "admin",
			Visibility:  model.PublicVisibility,
		}
		org, err := orgStore.CreateOrg(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, org)

		orgId = org.Id
	})
}

func TestMutationCreateWorkspace(t *testing.T) {
	store := store.NewWorkspaceStore(conn)

	t.Run("create workspace with duplicate scopeId", func(t *testing.T) {
		var workspaceId int32 = 0
		cmd := &workspace.CreateWorkspaceCommand{
			OrgId:       orgId,
			ScopeId:     "it",
			Name:        "IT",
			Description: "IT department",
			CreatedBy:   "admin",
			Visibility:  model.PublicVisibility,
		}
		{
			// first workspace
			ws, err := store.CreateWorkspace(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, ws)
			require.NotEmpty(t, ws.ScopeId)
			require.Equal(t, cmd.ScopeId, ws.ScopeId)

			workspaceId = ws.Id
		}
		{
			// second workspace (duplicate)
			wsDuplicate, err := store.CreateWorkspace(ctx, cmd)

			require.Nil(t, wsDuplicate)
			require.Error(t, err)
			require.Exactly(t, workspace.ErrWorkspaceAlreadyExists, err)
			require.Equal(t, "workspace with same scopeId already exists", err.Error())
		}
		{
			// clean up
			cmd := &workspace.DeleteWorkspaceByIdCommand{OrgId: orgId, WorkspaceId: workspaceId}
			deleteErr := store.DeleteWorkspaceById(ctx, cmd)
			require.NoError(t, deleteErr)

			query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: workspaceId}
			ws, err := store.GetWorkspaceById(ctx, query)
			require.Nil(t, ws)
			require.Error(t, err)
			require.Exactly(t, workspace.ErrWorkspaceNotFound, err)
			require.Equal(t, "workspace not found", err.Error())
		}
	})

	t.Run("create workspace without scopeId", func(t *testing.T) {
		for i := 0; i < workspaceCount; i++ {
			cmd := &workspace.CreateWorkspaceCommand{
				OrgId:       orgId,
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			ws, err := store.CreateWorkspace(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, ws)
			require.NotEmpty(t, ws.ScopeId)
			require.NotEmpty(t, ws.Name)
			require.NotEmpty(t, ws.Description)
			require.NotEmpty(t, ws.CreatedBy)
			require.NotZero(t, ws.Visibility)
			require.Equal(t, cmd.Name, ws.Name)
			require.Equal(t, cmd.Description, ws.Description)
			require.Equal(t, cmd.Visibility, ws.Visibility)
			require.Equal(t, cmd.CreatedBy, ws.CreatedBy)

			{
				// clean up
				cmd := &workspace.DeleteWorkspaceByIdCommand{OrgId: orgId, WorkspaceId: ws.Id}
				deleteErr := store.DeleteWorkspaceById(ctx, cmd)
				require.NoError(t, deleteErr)

				query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: ws.Id}
				ws, err := store.GetWorkspaceById(ctx, query)
				require.Nil(t, ws)
				require.Error(t, err)
				require.Exactly(t, workspace.ErrWorkspaceNotFound, err)
				require.Equal(t, "workspace not found", err.Error())
			}
		}
	})

	t.Run("create workspace with scopeId", func(t *testing.T) {
		for i := 0; i < workspaceCount; i++ {
			cmd := &workspace.CreateWorkspaceCommand{
				OrgId:       orgId,
				ScopeId:     fmt.Sprintf("quarkloop_%d", i),
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			ws, err := store.CreateWorkspace(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, ws)
			require.NotEmpty(t, ws.ScopeId)
			require.NotEmpty(t, ws.Name)
			require.NotEmpty(t, ws.Description)
			require.NotEmpty(t, ws.CreatedBy)
			require.NotZero(t, ws.Visibility)
			require.Equal(t, cmd.ScopeId, ws.ScopeId)
			require.Equal(t, cmd.Name, ws.Name)
			require.Equal(t, cmd.Description, ws.Description)
			require.Equal(t, cmd.Visibility, ws.Visibility)
			require.Equal(t, cmd.CreatedBy, ws.CreatedBy)
		}
	})

	t.Run("get workspace list return full", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)

		require.NoError(t, err)
		require.NotZero(t, len(wsList))
		require.Equal(t, workspaceCount, len(wsList))
	})
}

func TestQueryGetWorkspaceAfterCreate(t *testing.T) {
	store := store.NewWorkspaceStore(conn)

	t.Run("get workspace by id after creation", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)

		for idx, ws := range wsList {
			query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: ws.Id}
			workspace, err := store.GetWorkspaceById(ctx, query)

			require.NoError(t, err)
			require.NotNil(t, workspace)
			require.NotEmpty(t, workspace.ScopeId)
			require.NotEmpty(t, workspace.Name)
			require.NotEmpty(t, workspace.Description)
			require.NotEmpty(t, workspace.CreatedBy)
			require.NotZero(t, workspace.Visibility)
			require.Equal(t, fmt.Sprintf("quarkloop_%d", idx), workspace.ScopeId)
			require.Equal(t, fmt.Sprintf("Quarkloop_%d", idx), workspace.Name)
			require.Equal(t, fmt.Sprintf("Quarkloop Corporation #%d", idx), workspace.Description)
			require.Equal(t, fmt.Sprintf("admin_%d", idx), workspace.CreatedBy)
			require.Equal(t, model.PublicVisibility, workspace.Visibility)
		}
	})

	t.Run("get workspace by wrong id", func(t *testing.T) {
		query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: 9999999}
		ws, err := store.GetWorkspaceById(ctx, query)

		require.Nil(t, ws)
		require.Error(t, err)
		require.Exactly(t, workspace.ErrWorkspaceNotFound, err)
		require.Equal(t, "workspace not found", err.Error())
	})

	t.Run("get workspace visibility by id after creation", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)

		for _, ws := range wsList {
			query := &workspace.GetWorkspaceVisibilityByIdQuery{OrgId: orgId, WorkspaceId: ws.Id}
			visibility, err := store.GetWorkspaceVisibilityById(ctx, query)

			require.NoError(t, err)
			require.NotZero(t, visibility)
			require.Equal(t, model.PublicVisibility, visibility)
		}
	})
}

func TestMutationUpdateWorkspace(t *testing.T) {
	store := store.NewWorkspaceStore(conn)

	t.Run("update workspace with duplicate scope id", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)

		{
			// original scope id
			cmd := &workspace.UpdateWorkspaceByIdCommand{
				OrgId:       orgId,
				WorkspaceId: wsList[0].Id,
				ScopeId:     "quarkloop_updated_scopeid",
			}
			err := store.UpdateWorkspaceById(ctx, cmd)

			require.NoError(t, err)
		}
		{
			// duplicate scope id
			cmd := &workspace.UpdateWorkspaceByIdCommand{
				OrgId:       orgId,
				WorkspaceId: wsList[len(wsList)-1].Id,
				ScopeId:     "quarkloop_updated_scopeid",
			}
			err := store.UpdateWorkspaceById(ctx, cmd)

			require.Error(t, err)
			require.Exactly(t, workspace.ErrWorkspaceAlreadyExists, err)
			require.Equal(t, "workspace with same scopeId already exists", err.Error())
		}
	})

	t.Run("partial workspace update", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)

		// name
		for idx, ws := range wsList {
			name := fmt.Sprintf("Quarkloop_Updated_%d", idx)
			cmd := &workspace.UpdateWorkspaceByIdCommand{
				OrgId:       orgId,
				WorkspaceId: ws.Id,
				Name:        name,
			}
			err := store.UpdateWorkspaceById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: ws.Id}
				workspace, err := store.GetWorkspaceById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, workspace)
				require.Equal(t, name, workspace.Name)
				require.NotEmpty(t, workspace.ScopeId)
				require.NotEmpty(t, workspace.Name)
				require.NotEmpty(t, workspace.Description)
				require.NotEmpty(t, workspace.CreatedBy)
				require.NotZero(t, workspace.Visibility)
			}
		}
		// description
		for idx, ws := range wsList {
			description := fmt.Sprintf("Quarkloop_Description_Updated_%d", idx)
			cmd := &workspace.UpdateWorkspaceByIdCommand{
				OrgId:       orgId,
				WorkspaceId: ws.Id,
				Description: description,
			}
			err := store.UpdateWorkspaceById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: ws.Id}
				workspace, err := store.GetWorkspaceById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, workspace)
				require.Equal(t, description, workspace.Description)
				require.NotEmpty(t, workspace.ScopeId)
				require.NotEmpty(t, workspace.Name)
				require.NotEmpty(t, workspace.Description)
				require.NotEmpty(t, workspace.CreatedBy)
				require.NotZero(t, workspace.Visibility)
			}
		}
		// visibility
		for _, ws := range wsList {
			visibility := model.PrivateVisibility
			cmd := &workspace.UpdateWorkspaceByIdCommand{
				OrgId:       orgId,
				WorkspaceId: ws.Id,
				Visibility:  visibility,
			}
			err := store.UpdateWorkspaceById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: ws.Id}
				workspace, err := store.GetWorkspaceById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, workspace)
				require.Equal(t, visibility, workspace.Visibility)
				require.NotEmpty(t, workspace.ScopeId)
				require.NotEmpty(t, workspace.Name)
				require.NotEmpty(t, workspace.Description)
				require.NotEmpty(t, workspace.CreatedBy)
				require.NotZero(t, workspace.Visibility)
			}
		}
		// updatedBy
		for idx, ws := range wsList {
			updatedBy := fmt.Sprintf("Quarkloop_Admin2_Updated_%d", idx)
			cmd := &workspace.UpdateWorkspaceByIdCommand{
				OrgId:       orgId,
				WorkspaceId: ws.Id,
				UpdatedBy:   updatedBy,
			}
			err := store.UpdateWorkspaceById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: ws.Id}
				workspace, err := store.GetWorkspaceById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, workspace)
				require.Equal(t, updatedBy, *workspace.UpdatedBy)
				require.NotEmpty(t, workspace.ScopeId)
				require.NotEmpty(t, workspace.Name)
				require.NotEmpty(t, workspace.Description)
				require.NotEmpty(t, workspace.CreatedBy)
				require.NotZero(t, workspace.Visibility)
			}
		}
	})

	t.Run("update workspace with all fields", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)

		for idx, ws := range wsList {
			cmd := &workspace.UpdateWorkspaceByIdCommand{
				OrgId:       orgId,
				WorkspaceId: ws.Id,
				ScopeId:     fmt.Sprintf("quarkloop_new_update_%d", idx),
				Name:        fmt.Sprintf("Quarkloop_New_Update_%d", idx),
				Description: fmt.Sprintf("Quarkloop Corporation Updated With #%d", idx),
				UpdatedBy:   fmt.Sprintf("admin_1_updated_%d", idx),
				Visibility:  model.PrivateVisibility,
			}
			err := store.UpdateWorkspaceById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				query := &workspace.GetWorkspaceByIdQuery{OrgId: orgId, WorkspaceId: ws.Id}
				workspace, err := store.GetWorkspaceById(ctx, query)

				require.NoError(t, err)
				require.NotNil(t, workspace)
				require.Equal(t, cmd.ScopeId, workspace.ScopeId)
				require.Equal(t, cmd.Name, workspace.Name)
				require.Equal(t, cmd.Description, workspace.Description)
				require.Equal(t, cmd.Visibility, workspace.Visibility)
				require.Equal(t, cmd.UpdatedBy, *workspace.UpdatedBy)
				require.NotEmpty(t, workspace.ScopeId)
				require.NotEmpty(t, workspace.Name)
				require.NotEmpty(t, workspace.Description)
				require.NotZero(t, workspace.Visibility)
				require.NotNil(t, workspace.UpdatedBy)
			}
		}
	})
}

func TestQueryWorkspaceRelations(t *testing.T) {
	store := store.NewWorkspaceStore(conn)

	t.Run("get workspace's project list", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)

		for _, ws := range wsList {
			query := &workspace.GetProjectListQuery{
				WorkspaceId: ws.Id,
				Visibility:  model.AllVisibility,
			}
			// public and private
			{
				list, err := store.GetProjectList(ctx, query)

				require.NoError(t, err)
				require.Empty(t, list)
				require.Equal(t, 0, len(list))
			}
			// public
			{
				query.Visibility = model.PublicVisibility
				list, err := store.GetProjectList(ctx, query)

				require.NoError(t, err)
				require.Empty(t, list)
				require.Equal(t, 0, len(list))
			}
			// private
			{
				query.Visibility = model.PrivateVisibility
				list, err := store.GetProjectList(ctx, query)

				require.NoError(t, err)
				require.Empty(t, list)
				require.Equal(t, 0, len(list))
			}
		}
	})

	t.Run("get workspace's user assignment list", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)

		for _, ws := range wsList {
			query := &workspace.GetUserAssignmentListQuery{WorkspaceId: ws.Id}
			list, err := store.GetUserAssignmentList(ctx, query)

			require.NoError(t, err)
			require.Empty(t, list)
			require.Equal(t, 0, len(list))
		}
	})
}

func TestMutationDeleteWorkspace(t *testing.T) {
	store := store.NewWorkspaceStore(conn)

	t.Run("delete all workspaces by id", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)
		require.NoError(t, err)

		for _, ws := range wsList {
			cmd := &workspace.DeleteWorkspaceByIdCommand{OrgId: orgId, WorkspaceId: ws.Id}
			err := store.DeleteWorkspaceById(ctx, cmd)
			require.NoError(t, err)
		}
	})

	t.Run("get workspace list should return empty", func(t *testing.T) {
		wsList, err := test.GetFullWorkspaceList(ctx, conn)

		require.NoError(t, err)
		require.Zero(t, len(wsList))
		require.Equal(t, 0, len(wsList))
	})
}

func TestCleanup(t *testing.T) {
	orgStore := orgStore.NewOrgStore(conn)

	t.Run("delete org by id", func(t *testing.T) {
		err := orgStore.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: orgId})
		require.NoError(t, err)
	})
}
