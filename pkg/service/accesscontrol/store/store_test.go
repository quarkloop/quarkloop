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
	"github.com/quarkloop/quarkloop/pkg/service/project"
	projectStore "github.com/quarkloop/quarkloop/pkg/service/project/store"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
	workspaceStore "github.com/quarkloop/quarkloop/pkg/service/workspace/store"
	"github.com/quarkloop/quarkloop/pkg/test"
)

var (
	ctx               context.Context
	conn              *pgx.Conn
	orgId             int
	workspaceId       int
	projectId         int
	roleId            int
	userId            int = 1
	orgMemberId       int
	workspaceMemberId int
	projectMemberId   int
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
	projectStore := projectStore.NewProjectStore(conn)

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

	t.Run("create single project for whole test", func(t *testing.T) {
		cmd := &project.CreateProjectCommand{
			OrgId:       orgId,
			WorkspaceId: workspaceId,
			ScopeId:     "expenses",
			Name:        "Expenses",
			Description: "IT department Expenses",
			CreatedBy:   "admin",
			Visibility:  model.PublicVisibility,
		}
		pr, err := projectStore.CreateProject(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, pr)

		projectId = pr.Id
	})
}

// func TestMutationCreateGroup(t *testing.T) {
// 	store := store.NewAccessControlStore(conn)

// 	t.Run("create group", func(t *testing.T) {
// 		cmd := &accesscontrol.CreateUserGroupCommand{
// 			OrgId:     orgId,
// 			Name:      "administrators",
// 			CreatedBy: "admin",
// 			Users:     []int{1, 2, 3},
// 		}
// 		g, err := store.CreateUserGroup(ctx, cmd)

// 		require.NoError(t, err)
// 		require.NotNil(t, g)

// 		groupId = g.Id
// 	})
// }

// func TestQueryGroup(t *testing.T) {
// 	store := store.NewAccessControlStore(conn)

// 	t.Run("get group list by org id", func(t *testing.T) {
// 		query := &accesscontrol.GetUserGroupListQuery{OrgId: orgId}
// 		groupList, err := store.GetUserGroupList(ctx, query)

// 		require.NoError(t, err)
// 		require.NotEmpty(t, groupList)
// 		require.NotZero(t, len(groupList))
// 	})

// 	t.Run("get group list by wrong org id", func(t *testing.T) {
// 		query := &accesscontrol.GetUserGroupListQuery{OrgId: 99999}
// 		groupList, err := store.GetUserGroupList(ctx, query)

// 		require.NoError(t, err)
// 		require.Empty(t, groupList)
// 		require.Zero(t, len(groupList))
// 	})

// 	t.Run("get group by id", func(t *testing.T) {
// 		query := &accesscontrol.GetUserGroupByIdQuery{OrgId: orgId, UserGroupId: groupId}
// 		group, err := store.GetUserGroupById(ctx, query)

// 		require.NoError(t, err)
// 		require.NotNil(t, group)
// 	})

// 	t.Run("get group by wrong id", func(t *testing.T) {
// 		query := &accesscontrol.GetUserGroupByIdQuery{OrgId: orgId, UserGroupId: 99999}
// 		group, err := store.GetUserGroupById(ctx, query)

// 		require.Error(t, err)
// 		require.Nil(t, group)
// 		require.Exactly(t, accesscontrol.ErrUserGroupNotFound, err)
// 		require.Equal(t, "user group not found", err.Error())
// 	})
// }

func TestMutationCreateRole(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("create role", func(t *testing.T) {
		cmd := &accesscontrol.CreateRoleCommand{Name: "administrator", CreatedBy: "admin"}
		r, err := store.CreateRole(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, r)

		roleId = r.Id
	})
}

func TestQueryRole(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("get role list", func(t *testing.T) {
		roleList, err := store.GetRoleList(ctx)

		require.NoError(t, err)
		require.NotEmpty(t, roleList)
		require.NotZero(t, len(roleList))
	})

	t.Run("get role by id", func(t *testing.T) {
		query := &accesscontrol.GetRoleByIdQuery{RoleId: roleId}
		role, err := store.GetRoleById(ctx, query)

		require.NoError(t, err)
		require.NotNil(t, role)
	})

	t.Run("get role by wrong id", func(t *testing.T) {
		query := &accesscontrol.GetRoleByIdQuery{RoleId: 99999}
		role, err := store.GetRoleById(ctx, query)

		require.Error(t, err)
		require.Nil(t, role)
		require.Exactly(t, accesscontrol.ErrRoleNotFound, err)
		require.Equal(t, "role not found", err.Error())
	})
}

func TestMutationCreateOrgMember(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("create org member", func(t *testing.T) {
		cmd := &accesscontrol.CreateOrgMemberCommand{
			UserId:    userId,
			RoleId:    roleId,
			OrgId:     orgId,
			CreatedBy: "admin",
		}
		member, err := store.CreateOrgMember(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, member)

		orgMemberId = member.Id
	})

	t.Run("create workspace member", func(t *testing.T) {
		cmd := &accesscontrol.CreateWorkspaceMemberCommand{
			UserId:      userId,
			RoleId:      roleId,
			WorkspaceId: workspaceId,
			CreatedBy:   "admin",
		}
		member, err := store.CreateWorkspaceMember(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, member)

		workspaceMemberId = member.Id
	})

	t.Run("create project member", func(t *testing.T) {
		cmd := &accesscontrol.CreateProjectMemberCommand{
			UserId:    userId,
			RoleId:    roleId,
			ProjectId: projectId,
			CreatedBy: "admin",
		}
		member, err := store.CreateProjectMember(ctx, cmd)

		require.NoError(t, err)
		require.NotNil(t, member)

		projectMemberId = member.Id
	})
}

func TestQueryMemberList(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("get org member list by org id", func(t *testing.T) {
		query := &accesscontrol.GetOrgMemberListQuery{OrgId: orgId}
		memberList, err := store.GetOrgMemberList(ctx, query)

		require.NoError(t, err)
		require.NotEmpty(t, memberList)
		require.NotZero(t, len(memberList))
	})

	t.Run("get org member list by wrong org id", func(t *testing.T) {
		query := &accesscontrol.GetOrgMemberListQuery{OrgId: 99999}
		memberList, err := store.GetOrgMemberList(ctx, query)

		require.NoError(t, err)
		require.Empty(t, memberList)
		require.Zero(t, len(memberList))
	})

	t.Run("get workspace member list by workspace id", func(t *testing.T) {
		query := &accesscontrol.GetWorkspaceMemberListQuery{WorkspaceId: workspaceId}
		memberList, err := store.GetWorkspaceMemberList(ctx, query)

		require.NoError(t, err)
		require.NotEmpty(t, memberList)
		require.NotZero(t, len(memberList))
	})

	t.Run("get workspace member list by wrong workspace id", func(t *testing.T) {
		query := &accesscontrol.GetWorkspaceMemberListQuery{WorkspaceId: 99999}
		memberList, err := store.GetWorkspaceMemberList(ctx, query)

		require.NoError(t, err)
		require.Empty(t, memberList)
		require.Zero(t, len(memberList))
	})

	t.Run("get project member list by project id", func(t *testing.T) {
		query := &accesscontrol.GetProjectMemberListQuery{ProjectId: projectId}
		memberList, err := store.GetProjectMemberList(ctx, query)

		require.NoError(t, err)
		require.NotEmpty(t, memberList)
		require.NotZero(t, len(memberList))
	})

	t.Run("get project member list by wrong project id", func(t *testing.T) {
		query := &accesscontrol.GetProjectMemberListQuery{ProjectId: 99999}
		memberList, err := store.GetProjectMemberList(ctx, query)

		require.NoError(t, err)
		require.Empty(t, memberList)
		require.Zero(t, len(memberList))
	})
}

func TestQueryMemberById(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("get org member by wrong org id", func(t *testing.T) {
		query := &accesscontrol.GetOrgMemberByIdQuery{OrgId: 99999, MemberId: orgMemberId}
		member, err := store.GetOrgMemberById(ctx, query)

		require.Error(t, err)
		require.Nil(t, member)
	})

	t.Run("get org member by id", func(t *testing.T) {
		query := &accesscontrol.GetOrgMemberByIdQuery{OrgId: orgId, MemberId: orgMemberId}
		member, err := store.GetOrgMemberById(ctx, query)

		require.NoError(t, err)
		require.NotNil(t, member)
	})

	t.Run("get workspace member by wrong workspace id", func(t *testing.T) {
		query := &accesscontrol.GetWorkspaceMemberByIdQuery{WorkspaceId: 99999, MemberId: workspaceMemberId}
		member, err := store.GetWorkspaceMemberById(ctx, query)

		require.Error(t, err)
		require.Nil(t, member)
	})

	t.Run("get workspace member by id", func(t *testing.T) {
		query := &accesscontrol.GetWorkspaceMemberByIdQuery{WorkspaceId: workspaceId, MemberId: workspaceMemberId}
		member, err := store.GetWorkspaceMemberById(ctx, query)

		require.NoError(t, err)
		require.NotNil(t, member)
	})

	t.Run("get project member by wrong project id", func(t *testing.T) {
		query := &accesscontrol.GetProjectMemberByIdQuery{ProjectId: 99999, MemberId: projectMemberId}
		member, err := store.GetProjectMemberById(ctx, query)

		require.Error(t, err)
		require.Nil(t, member)
	})

	t.Run("get project member by id", func(t *testing.T) {
		query := &accesscontrol.GetProjectMemberByIdQuery{ProjectId: projectId, MemberId: projectMemberId}
		member, err := store.GetProjectMemberById(ctx, query)

		require.NoError(t, err)
		require.NotNil(t, member)
	})
}

// func TestQueryEvaluateUserAssignment(t *testing.T) {
// 	store := store.NewAccessControlStore(conn)

// 	t.Run("evaluate assignment by orgId", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission: accesscontrol.ActionOrgCreate,
// 			UserId:     userId,
// 			OrgId:      orgId,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.True(t, permission)
// 	})

// 	t.Run("evaluate assignment by wrong orgId", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission: accesscontrol.ActionOrgCreate,
// 			UserId:     userId,
// 			OrgId:      99999,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.False(t, permission)
// 	})

// 	t.Run("evaluate assignment by orgId and wrong permission", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission: accesscontrol.ActionProjectUserDelete,
// 			UserId:     userId,
// 			OrgId:      orgId,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.False(t, permission)
// 	})

// 	t.Run("evaluate assignment by workspaceId", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission:  accesscontrol.ActionOrgCreate,
// 			UserId:      userId,
// 			OrgId:       orgId,
// 			WorkspaceId: workspaceId,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.True(t, permission)
// 	})

// 	t.Run("evaluate assignment by wrong workspaceId", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission:  accesscontrol.ActionOrgCreate,
// 			UserId:      userId,
// 			OrgId:       orgId,
// 			WorkspaceId: 99999,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.False(t, permission)
// 	})

// 	t.Run("evaluate assignment by workspaceId and wrong permission", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission:  accesscontrol.ActionProjectQuotaCreate,
// 			UserId:      userId,
// 			OrgId:       orgId,
// 			WorkspaceId: workspaceId,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.False(t, permission)
// 	})

// 	t.Run("evaluate assignment by projectId", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission:  accesscontrol.ActionOrgCreate,
// 			UserId:      userId,
// 			OrgId:       orgId,
// 			WorkspaceId: workspaceId,
// 			ProjectId:   projectId,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.True(t, permission)
// 	})

// 	t.Run("evaluate assignment by wrong projectId", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission:  accesscontrol.ActionOrgCreate,
// 			UserId:      userId,
// 			OrgId:       orgId,
// 			WorkspaceId: workspaceId,
// 			ProjectId:   99999,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.False(t, permission)
// 	})

// 	t.Run("evaluate assignment by projectId and wrong permission", func(t *testing.T) {
// 		query := &accesscontrol.EvaluateQuery{
// 			Permission:  accesscontrol.ActionWorkspaceUserUpdate,
// 			UserId:      userId,
// 			OrgId:       orgId,
// 			WorkspaceId: workspaceId,
// 			ProjectId:   projectId,
// 		}
// 		permission, err := store.EvaluateUserAccess(ctx, query)

// 		require.NoError(t, err)
// 		require.False(t, permission)
// 	})
// }

func TestMutationDeleteUserAssignment(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	t.Run("delete org member by wrong id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteOrgMemberByIdCommand{
			OrgId:    orgId,
			MemberId: 99999,
		}
		err := store.DeleteOrgMemberById(ctx, cmd)

		require.Error(t, err)
	})

	t.Run("delete org member by id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteOrgMemberByIdCommand{
			OrgId:    orgId,
			MemberId: orgMemberId,
		}
		err := store.DeleteOrgMemberById(ctx, cmd)

		require.NoError(t, err)
	})

	t.Run("delete workspace member by wrong id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteWorkspaceMemberByIdCommand{
			WorkspaceId: workspaceId,
			MemberId:    99999,
		}
		err := store.DeleteWorkspaceMemberById(ctx, cmd)

		require.Error(t, err)
	})

	t.Run("delete workspace member by id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteWorkspaceMemberByIdCommand{
			WorkspaceId: workspaceId,
			MemberId:    workspaceMemberId,
		}
		err := store.DeleteWorkspaceMemberById(ctx, cmd)

		require.NoError(t, err)
	})

	t.Run("delete project member by wrong id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteProjectMemberByIdCommand{
			ProjectId: projectId,
			MemberId:  99999,
		}
		err := store.DeleteProjectMemberById(ctx, cmd)

		require.Error(t, err)
	})

	t.Run("delete project member by id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteProjectMemberByIdCommand{
			ProjectId: projectId,
			MemberId:  projectMemberId,
		}
		err := store.DeleteProjectMemberById(ctx, cmd)

		require.NoError(t, err)
	})
}

func TestMutationDelete(t *testing.T) {
	store := store.NewAccessControlStore(conn)

	// t.Run("delete group with its users by wrong group id", func(t *testing.T) {
	// 	cmd := &accesscontrol.DeleteUserGroupByIdCommand{
	// 		OrgId:       orgId,
	// 		UserGroupId: 99999,
	// 	}
	// 	err := store.DeleteUserGroupById(ctx, cmd)

	// 	require.Error(t, err)
	// })

	// t.Run("delete group with its users by group id", func(t *testing.T) {
	// 	cmd := &accesscontrol.DeleteUserGroupByIdCommand{
	// 		OrgId:       orgId,
	// 		UserGroupId: groupId,
	// 	}
	// 	err := store.DeleteUserGroupById(ctx, cmd)

	// 	require.NoError(t, err)
	// })

	t.Run("delete role by wrong id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteRoleByIdCommand{RoleId: 99999}
		err := store.DeleteRoleById(ctx, cmd)

		require.Error(t, err)
	})

	t.Run("delete role by id", func(t *testing.T) {
		cmd := &accesscontrol.DeleteRoleByIdCommand{RoleId: roleId}
		err := store.DeleteRoleById(ctx, cmd)

		require.NoError(t, err)
	})
}
