package store_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	"github.com/quarkloop/quarkloop/pkg/service/org/store"
	"github.com/quarkloop/quarkloop/pkg/test"
)

var (
	ctx  context.Context
	conn *pgx.Conn
)

const orgCount = 10

func init() {
	ctx, conn = test.InitTestSystemDB()
}

func TestMutationTruncateTables(t *testing.T) {
	t.Run("truncate tables", func(t *testing.T) {
		err := test.TruncateSystemDBTables(ctx, conn)
		require.NoError(t, err)
	})

	t.Run("get org list return empty after truncating tables", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)
		require.Zero(t, len(orgList))
		require.Equal(t, 0, len(orgList))
	})
}

func TestMutationCreateOrg(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("create org with duplicate scopeId", func(t *testing.T) {
		var orgId int32 = 0
		cmd := &org.CreateOrgCommand{
			ScopeId:     "acme",
			Name:        "ACME",
			Description: "ACME Corporation",
			CreatedBy:   "admin",
			Visibility:  model.PublicVisibility,
		}
		{
			// first org
			org, err := store.CreateOrg(ctx, cmd)
			orgId = org.Id

			require.NoError(t, err)
			require.NotNil(t, org)
			require.NotEmpty(t, org.ScopeId)
			require.Equal(t, cmd.ScopeId, org.ScopeId)
		}
		{
			// second org (duplicate)
			orgDuplicate, err := store.CreateOrg(ctx, cmd)

			require.Nil(t, orgDuplicate)
			require.Error(t, err)
			require.Exactly(t, org.ErrOrgAlreadyExists, err)
			require.Equal(t, "org with same scopeId already exists", err.Error())
		}
		{
			// clean up
			deleteErr := store.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: orgId})
			require.NoError(t, deleteErr)

			o, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: orgId})
			require.Nil(t, o)
			require.Error(t, err)
			require.Exactly(t, org.ErrOrgNotFound, err)
			require.Equal(t, "org not found", err.Error())
		}
	})

	t.Run("create org without scopeId", func(t *testing.T) {
		for i := 0; i < orgCount; i++ {
			cmd := &org.CreateOrgCommand{
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			o, err := store.CreateOrg(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, o)
			require.NotEmpty(t, o.ScopeId)
			require.NotEmpty(t, o.Name)
			require.NotEmpty(t, o.Description)
			require.NotEmpty(t, o.CreatedBy)
			require.NotZero(t, o.Visibility)
			require.Equal(t, cmd.Name, o.Name)
			require.Equal(t, cmd.Description, o.Description)
			require.Equal(t, cmd.Visibility, o.Visibility)
			require.Equal(t, cmd.CreatedBy, o.CreatedBy)

			{
				// clean up
				deleteErr := store.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: o.Id})
				require.NoError(t, deleteErr)

				o, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
				require.Nil(t, o)
				require.Error(t, err)
				require.Exactly(t, org.ErrOrgNotFound, err)
				require.Equal(t, "org not found", err.Error())
			}
		}
	})

	t.Run("create org with scopeId", func(t *testing.T) {
		for i := 0; i < orgCount; i++ {
			cmd := &org.CreateOrgCommand{
				ScopeId:     fmt.Sprintf("quarkloop_%d", i),
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			org, err := store.CreateOrg(ctx, cmd)

			require.NoError(t, err)
			require.NotNil(t, org)
			require.NotEmpty(t, org.ScopeId)
			require.NotEmpty(t, org.Name)
			require.NotEmpty(t, org.Description)
			require.NotEmpty(t, org.CreatedBy)
			require.NotZero(t, org.Visibility)
			require.Equal(t, cmd.ScopeId, org.ScopeId)
			require.Equal(t, cmd.Name, org.Name)
			require.Equal(t, cmd.Description, org.Description)
			require.Equal(t, cmd.Visibility, org.Visibility)
			require.Equal(t, cmd.CreatedBy, org.CreatedBy)
		}
	})

	t.Run("get org list return full", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)

		require.NoError(t, err)
		require.NotZero(t, len(orgList))
		require.Equal(t, orgCount, len(orgList))
	})
}

func TestQueryGetOrgAfterCreate(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("get org by id after creation", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for idx, o := range orgList {
			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})

			require.NoError(t, err)
			require.NotNil(t, org)
			require.Equal(t, fmt.Sprintf("quarkloop_%d", idx), org.ScopeId)
			require.Equal(t, fmt.Sprintf("Quarkloop_%d", idx), org.Name)
			require.Equal(t, fmt.Sprintf("Quarkloop Corporation #%d", idx), org.Description)
			require.Equal(t, fmt.Sprintf("admin_%d", idx), org.CreatedBy)
			require.Equal(t, model.PublicVisibility, org.Visibility)
		}
	})

	t.Run("get org by wrong id", func(t *testing.T) {
		o, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: 9999999})

		require.Nil(t, o)
		require.Error(t, err)
		require.Exactly(t, org.ErrOrgNotFound, err)
		require.Equal(t, "org not found", err.Error())
	})

	t.Run("get org visibility by id after creation", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			visibility, err := store.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: o.Id})

			require.NoError(t, err)
			require.NotZero(t, visibility)
			require.Equal(t, model.PublicVisibility, visibility)
		}
	})
}

func TestMutationUpdateOrg(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("update org with duplicate scope id", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		{
			// original scope id
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:   orgList[0].Id,
				ScopeId: "quarkloop_updated_scopeid",
			}
			err := store.UpdateOrgById(ctx, cmd)

			require.NoError(t, err)
		}
		{
			// duplicate scope id
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:   orgList[len(orgList)-1].Id,
				ScopeId: "quarkloop_updated_scopeid",
			}
			err := store.UpdateOrgById(ctx, cmd)

			require.Error(t, err)
			require.Exactly(t, org.ErrOrgAlreadyExists, err)
			require.Equal(t, "org with same scopeId already exists", err.Error())
		}
	})

	t.Run("partial org update", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		// name
		for idx, o := range orgList {
			name := fmt.Sprintf("Quarkloop_Updated_%d", idx)
			cmd := &org.UpdateOrgByIdCommand{
				OrgId: o.Id,
				Name:  name,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})

				require.NoError(t, err)
				require.NotNil(t, org)
				require.Equal(t, name, org.Name)
				require.NotEmpty(t, org.ScopeId)
				require.NotEmpty(t, org.Name)
				require.NotEmpty(t, org.Description)
				require.NotEmpty(t, org.CreatedBy)
				require.NotZero(t, org.Visibility)
			}
		}
		// description
		for idx, o := range orgList {
			description := fmt.Sprintf("Quarkloop_Description_Updated_%d", idx)
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:       o.Id,
				Description: description,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})

				require.NoError(t, err)
				require.NotNil(t, org)
				require.Equal(t, description, org.Description)
				require.NotEmpty(t, org.ScopeId)
				require.NotEmpty(t, org.Name)
				require.NotEmpty(t, org.Description)
				require.NotEmpty(t, org.CreatedBy)
				require.NotZero(t, org.Visibility)
			}
		}
		// visibility
		for _, o := range orgList {
			visibility := model.PrivateVisibility
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:      o.Id,
				Visibility: visibility,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})

				require.NoError(t, err)
				require.NotNil(t, org)
				require.Equal(t, visibility, org.Visibility)
				require.NotEmpty(t, org.ScopeId)
				require.NotEmpty(t, org.Name)
				require.NotEmpty(t, org.Description)
				require.NotEmpty(t, org.CreatedBy)
				require.NotZero(t, org.Visibility)
			}
		}
		// updatedBy
		for idx, o := range orgList {
			updatedBy := fmt.Sprintf("Quarkloop_Admin2_Updated_%d", idx)
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:     o.Id,
				UpdatedBy: updatedBy,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})

				require.NoError(t, err)
				require.NotNil(t, org)
				require.Equal(t, updatedBy, *org.UpdatedBy)
				require.NotEmpty(t, org.ScopeId)
				require.NotEmpty(t, org.Name)
				require.NotEmpty(t, org.Description)
				require.NotEmpty(t, org.CreatedBy)
				require.NotZero(t, org.Visibility)
			}
		}
	})

	t.Run("update org with all fields", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for idx, o := range orgList {
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:       o.Id,
				ScopeId:     fmt.Sprintf("quarkloop_new_update_%d", idx),
				Name:        fmt.Sprintf("Quarkloop_New_Update_%d", idx),
				Description: fmt.Sprintf("Quarkloop Corporation Updated With #%d", idx),
				UpdatedBy:   fmt.Sprintf("admin_1_updated_%d", idx),
				Visibility:  model.PrivateVisibility,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			{
				// check the update
				org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})

				require.NoError(t, err)
				require.NotNil(t, org)
				require.Equal(t, cmd.ScopeId, org.ScopeId)
				require.Equal(t, cmd.Name, org.Name)
				require.Equal(t, cmd.Description, org.Description)
				require.Equal(t, cmd.Visibility, org.Visibility)
				require.Equal(t, cmd.UpdatedBy, *org.UpdatedBy)
				require.NotEmpty(t, org.ScopeId)
				require.NotEmpty(t, org.Name)
				require.NotEmpty(t, org.Description)
				require.NotZero(t, org.Visibility)
				require.NotNil(t, org.UpdatedBy)
			}
		}
	})
}

func TestQueryOrgRelations(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("get org's workspace list", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			query := &org.GetWorkspaceListQuery{
				OrgId:      o.Id,
				Visibility: model.AllVisibility,
			}
			// public and private
			{
				list, err := store.GetWorkspaceList(ctx, query)

				require.NoError(t, err)
				require.Empty(t, list)
				require.Equal(t, 0, len(list))
			}
			// public
			{
				query.Visibility = model.PublicVisibility
				list, err := store.GetWorkspaceList(ctx, query)

				require.NoError(t, err)
				require.Empty(t, list)
				require.Equal(t, 0, len(list))
			}
			// private
			{
				query.Visibility = model.PrivateVisibility
				list, err := store.GetWorkspaceList(ctx, query)

				require.NoError(t, err)
				require.Empty(t, list)
				require.Equal(t, 0, len(list))
			}
		}
	})

	t.Run("get org's project list", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			query := &org.GetProjectListQuery{
				OrgId:      o.Id,
				Visibility: model.AllVisibility,
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

	// t.Run("get org's user assignment list", func(t *testing.T) {
	// 	orgList, err := test.GetFullOrgList(ctx, conn)
	// 	require.NoError(t, err)

	// 	for _, o := range orgList {
	// 		query := &org.GetUserAssignmentListQuery{OrgId: o.Id}
	// 		list, err := store.GetUserAssignmentList(ctx, query)

	// 		require.NoError(t, err)
	// 		require.Empty(t, list)
	// 		require.Equal(t, 0, len(list))
	// 	}
	// })
}

func TestMutationDeleteOrg(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("delete all orgs by id", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			err := store.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: o.Id})
			require.NoError(t, err)
		}
	})

	t.Run("get org list should return empty", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)

		require.NoError(t, err)
		require.Zero(t, len(orgList))
		require.Equal(t, 0, len(orgList))
	})
}
