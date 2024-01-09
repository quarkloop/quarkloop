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
	roleId      int
	groupId     int
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

func TestMutationCreateGroup(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("create group", func(t *testing.T) {
		cmd := &accesscontrol.CreateUserGroupCommand{
			OrgId:     orgId,
			Name:      "administrators",
			CreatedBy: "admin",
			Users:     []int{1, 2, 3},
		}
		g, err := store.CreateUserGroup(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, g)

		groupId = g.Id
	})
}

func TestQueryGroup(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("get group list by org id", func(t *testing.T) {
		query := &accesscontrol.GetUserGroupListQuery{OrgId: orgId}
		groupList, err := store.GetUserGroupList(ctx, query)

		require.NoError(t, err)
		require.NotEmpty(t, groupList)
		require.NotZero(t, len(groupList))
	})

	t.Run("get group list by wrong org id", func(t *testing.T) {
		query := &accesscontrol.GetUserGroupListQuery{OrgId: 99999}
		groupList, err := store.GetUserGroupList(ctx, query)

		require.NoError(t, err)
		require.Empty(t, groupList)
		require.Zero(t, len(groupList))
	})

	t.Run("get group by id", func(t *testing.T) {
		query := &accesscontrol.GetUserGroupByIdQuery{OrgId: orgId, UserGroupId: groupId}
		group, err := store.GetUserGroupById(ctx, query)

		require.NoError(t, err)
		require.NotNil(t, group)
	})

	t.Run("get group by wrong id", func(t *testing.T) {
		query := &accesscontrol.GetUserGroupByIdQuery{OrgId: orgId, UserGroupId: 99999}
		group, err := store.GetUserGroupById(ctx, query)

		require.Error(t, err)
		require.Nil(t, group)
		require.Exactly(t, accesscontrol.ErrUserGroupNotFound, err)
		require.Equal(t, "user group not found", err.Error())
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

		roleId = r.Id
	})
}

func TestQueryRole(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("get role list by org id", func(t *testing.T) {
		query := &accesscontrol.GetUserRoleListQuery{OrgId: orgId}
		roleList, err := store.GetUserRoleList(ctx, query)

		require.NoError(t, err)
		require.NotEmpty(t, roleList)
		require.NotZero(t, len(roleList))
	})

	t.Run("get role list by wrong org id", func(t *testing.T) {
		query := &accesscontrol.GetUserRoleListQuery{OrgId: 99999}
		roleList, err := store.GetUserRoleList(ctx, query)

		require.NoError(t, err)
		require.Empty(t, roleList)
		require.Zero(t, len(roleList))
	})

	t.Run("get role by id", func(t *testing.T) {
		query := &accesscontrol.GetUserRoleByIdQuery{OrgId: orgId, UserRoleId: roleId}
		role, err := store.GetUserRoleById(ctx, query)

		require.NoError(t, err)
		require.NotNil(t, role)
	})

	t.Run("get role by wrong id", func(t *testing.T) {
		query := &accesscontrol.GetUserRoleByIdQuery{OrgId: orgId, UserRoleId: 99999}
		role, err := store.GetUserRoleById(ctx, query)

		require.Error(t, err)
		require.Nil(t, role)
		require.Exactly(t, accesscontrol.ErrRoleNotFound, err)
		require.Equal(t, "role not found", err.Error())
	})
}

func TestMutationDelete(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("delete group with its users by wrong group id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteUserGroupByIdCommand{
			OrgId:       orgId,
			UserGroupId: 99999,
		}
		err := store.DeleteUserGroupById(ctx, cmd)

		require.Error(t, err)
	})

	t.Run("delete group with its users by group id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteUserGroupByIdCommand{
			OrgId:       orgId,
			UserGroupId: groupId,
		}
		err := store.DeleteUserGroupById(ctx, cmd)

		require.NoError(t, err)
	})

	t.Run("delete role with its permissions by wrong role id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteUserRoleByIdCommand{
			OrgId:      orgId,
			UserRoleId: 99999,
		}
		err := store.DeleteUserRoleById(ctx, cmd)

		require.Error(t, err)
	})

	t.Run("delete role with its permissions by role id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteUserRoleByIdCommand{
			OrgId:      orgId,
			UserRoleId: roleId,
		}
		err := store.DeleteUserRoleById(ctx, cmd)

		require.NoError(t, err)
	})
}
