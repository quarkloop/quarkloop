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

func TestOrgServiceReturnEmpty(t *testing.T) {
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

func TestOrgServiceCreate(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("CreateOrg", func(t *testing.T) {
		for i := 0; i < orgCount; i++ {
			cmd := &org.CreateOrgCommand{
				Name:        fmt.Sprintf("Quarkloop_%d", i),
				ScopeId:     fmt.Sprintf("quarkloop_%d", i),
				Description: fmt.Sprintf("Quarkloop Corporation #%d", i),
				CreatedBy:   fmt.Sprintf("admin_%d", i),
				Visibility:  model.PublicVisibility,
			}
			org, err := store.CreateOrg(ctx, cmd)
			require.NoError(t, err)
			require.NotEmpty(t, org)
		}
	})

	t.Run("get org list return full", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)
		require.NotZero(t, len(orgList))
		require.Equal(t, orgCount, len(orgList))
	})
}

func TestOrgService(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("GetOrgById after creation", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for idx, o := range orgList {
			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, fmt.Sprintf("Quarkloop_%d", idx), org.Name)
			require.Equal(t, fmt.Sprintf("quarkloop_%d", idx), org.ScopeId)
			require.Equal(t, fmt.Sprintf("Quarkloop Corporation #%d", idx), org.Description)
			require.Equal(t, fmt.Sprintf("admin_%d", idx), org.CreatedBy)
			require.Equal(t, model.PublicVisibility, org.Visibility)
		}
	})

	t.Run("GetOrgVisibilityById after creation", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			visibility, err := store.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, visibility)
			require.Equal(t, model.PublicVisibility, visibility)
		}
	})

	t.Run("UpdateOrgById: ScopeId update fail with duplicate Organization_sid_key", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)
		{
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:   orgList[0].Id,
				ScopeId: "Quarkloop_Updated",
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)
		}
		{
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:   orgList[len(orgList)-1].Id,
				ScopeId: "Quarkloop_Updated",
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.Error(t, err)
			require.Exactly(t, org.ErrOrgAlreadyExists, err)
			require.Equal(t, "org with same scopeId already exists", err.Error())
		}
	})

	t.Run("UpdateOrgById: partial update", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for idx, o := range orgList {
			name := fmt.Sprintf("Quarkloop_Updated_%d", idx)

			cmd := &org.UpdateOrgByIdCommand{
				OrgId: o.Id,
				Name:  name,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, name, org.Name)

			require.NotEmpty(t, org.ScopeId)
			require.NotEmpty(t, org.Name)
			require.NotEmpty(t, org.Description)
			require.NotEmpty(t, org.Visibility)
			require.NotEmpty(t, org.CreatedBy)
		}
		for idx, o := range orgList {
			description := fmt.Sprintf("Quarkloop_Description_Updated_%d", idx)

			cmd := &org.UpdateOrgByIdCommand{
				OrgId:       o.Id,
				Description: description,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, description, org.Description)

			require.NotEmpty(t, org.ScopeId)
			require.NotEmpty(t, org.Name)
			require.NotEmpty(t, org.Description)
			require.NotEmpty(t, org.Visibility)
			require.NotEmpty(t, org.CreatedBy)
		}
		for _, o := range orgList {
			visibility := model.PrivateVisibility

			cmd := &org.UpdateOrgByIdCommand{
				OrgId:      o.Id,
				Visibility: visibility,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, visibility, org.Visibility)

			require.NotEmpty(t, org.ScopeId)
			require.NotEmpty(t, org.Name)
			require.NotEmpty(t, org.Description)
			require.NotEmpty(t, org.Visibility)
			require.NotEmpty(t, org.CreatedBy)
		}
		for idx, o := range orgList {
			updatedBy := fmt.Sprintf("Quarkloop_Admin2_Updated_%d", idx)

			cmd := &org.UpdateOrgByIdCommand{
				OrgId:     o.Id,
				UpdatedBy: updatedBy,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)
			require.Equal(t, updatedBy, *org.UpdatedBy)

			require.NotEmpty(t, org.ScopeId)
			require.NotEmpty(t, org.Name)
			require.NotEmpty(t, org.Description)
			require.NotEmpty(t, org.Visibility)
			require.NotEmpty(t, org.CreatedBy)
		}
	})

	t.Run("UpdateOrgById: full update", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for idx, o := range orgList {
			cmd := &org.UpdateOrgByIdCommand{
				OrgId:       o.Id,
				Name:        fmt.Sprintf("Quarkloop_Updated_%d", idx),
				ScopeId:     fmt.Sprintf("quarkloop_Updated_%d", idx),
				Description: fmt.Sprintf("Quarkloop Corporation Updated #%d", idx),
				UpdatedBy:   fmt.Sprintf("admin_Updated_%d", idx),
				Visibility:  model.PrivateVisibility,
			}
			err := store.UpdateOrgById(ctx, cmd)
			require.NoError(t, err)

			org, err := store.GetOrgById(ctx, &org.GetOrgByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, org)

			require.Equal(t, cmd.ScopeId, org.ScopeId)
			require.Equal(t, cmd.Name, org.Name)
			require.Equal(t, cmd.Description, org.Description)
			require.Equal(t, cmd.Visibility, org.Visibility)
			require.Equal(t, cmd.UpdatedBy, *org.UpdatedBy)

			require.NotEmpty(t, org.ScopeId)
			require.NotEmpty(t, org.Name)
			require.NotEmpty(t, org.Description)
			require.NotEmpty(t, org.Visibility)
			require.NotEmpty(t, org.CreatedBy)
		}
	})

	t.Run("GetOrgVisibilityById after update", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			visibility, err := store.GetOrgVisibilityById(ctx, &org.GetOrgVisibilityByIdQuery{OrgId: o.Id})
			require.NoError(t, err)
			require.NotEmpty(t, visibility)
			require.Equal(t, model.PrivateVisibility, visibility)
		}
	})
}

func TestOrgServiceRelations(t *testing.T) {
	store := store.NewOrgStore(conn)

	t.Run("GetWorkspaceList", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			query := &org.GetWorkspaceListQuery{
				OrgId:      o.Id,
				Visibility: model.AllVisibility,
			}
			list, err := store.GetWorkspaceList(ctx, query)
			require.NoError(t, err)
			require.Empty(t, list)
			require.Equal(t, 0, len(list))
		}
	})

	t.Run("GetProjectList", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			query := &org.GetProjectListQuery{
				OrgId:      o.Id,
				Visibility: model.AllVisibility,
			}
			list, err := store.GetProjectList(ctx, query)
			require.NoError(t, err)
			require.Empty(t, list)
			require.Equal(t, 0, len(list))
		}
	})

	t.Run("GetUserAssignmentList", func(t *testing.T) {
		orgList, err := test.GetFullOrgList(ctx, conn)
		require.NoError(t, err)

		for _, o := range orgList {
			query := &org.GetUserAssignmentListQuery{OrgId: o.Id}
			list, err := store.GetUserAssignmentList(ctx, query)
			require.NoError(t, err)
			require.Empty(t, list)
			require.Equal(t, 0, len(list))
		}
	})
}

// func TestOrgServiceDelete(t *testing.T) {
// 	store := store.NewOrgStore(conn)

// 	t.Run("DeleteOrgById", func(t *testing.T) {
// 		orgList, err := test.GetFullOrgList(ctx, conn)
// 		require.NoError(t, err)

// 		for _, o := range orgList {
// 			err := store.DeleteOrgById(ctx, &org.DeleteOrgByIdCommand{OrgId: o.Id})
// 			require.NoError(t, err)
// 		}
// 	})

// 	t.Run("get org list return empty after deleting tables", func(t *testing.T) {
// 		orgList, err := test.GetFullOrgList(ctx, conn)
// 		require.NoError(t, err)
// 		require.Zero(t, len(orgList))
// 		require.Equal(t, 0, len(orgList))
// 	})
// }
