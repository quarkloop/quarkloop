package store_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol/store"
	"github.com/quarkloop/quarkloop/pkg/service/org"
	orgStore "github.com/quarkloop/quarkloop/pkg/service/org/store"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	workspaceStore "github.com/quarkloop/quarkloop/pkg/service/workspace/store"
	"github.com/quarkloop/quarkloop/pkg/test"
)

var (
	ctx         context.Context
	conn        *pgx.Conn
	orgId       int
	workspaceId int
)

func init() {
	ctx, conn = test.InitTestSystemDB()
}

func TestMutationTruncateTables(t *testing.T) {
	t.Run("truncate tables", func(t *testing.T) {
		err := test.TruncateSystemDBTables(ctx, conn)
		require.NoError(t, err)
	})
}

func TestPrepare(t *testing.T) {
	orgStore := orgStore.NewOrgStore(conn)
	workspaceStore := workspaceStore.NewWorkspaceStore(conn)

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

	t.Run("create single workspace for whole test", func(t *testing.T) {
		cmd := &workspace.CreateWorkspaceCommand{
			OrgId:       orgId,
			ScopeId:     "it",
			Name:        "IT department",
			Description: "Quarkloop Corporation IT department",
			CreatedBy:   "admin",
			Visibility:  model.PublicVisibility,
		}
		ws, err := workspaceStore.CreateWorkspace(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, ws)

		workspaceId = ws.Id
	})
}

func TestMutationCreateRole(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("create role", func(t *testing.T) {
		cmd := &accesscontrol.CreateUserRoleCommand{
			OrgId:     orgId,
			Name:      "administrator",
			CreatedBy: "admin",
			Permissions: []struct {
				Name string
			}{
				{Name: accesscontrol.ActionOrgCreate},
				{Name: accesscontrol.ActionOrgDelete},
				{Name: accesscontrol.ActionOrgList},
			},
		}
		r, err := store.CreateUserRole(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, r)
	})

	t.Run("get role list should return empty", func(t *testing.T) {

	})
}

// func TestMutationDeleteRole(t *testing.T) {
// 	store := store.NewAccessControlStore(conn)

// 	t.Run("delete all roles by id", func(t *testing.T) {

// 	})

// 	t.Run("get role list should return empty", func(t *testing.T) {

// 	})
// }
