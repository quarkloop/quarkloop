package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

type AccessControlStore interface {
	// user access query and mutation
	EvaluateUserAccess(context.Context, *accesscontrol.EvaluateQuery) error
	GrantUserAccess(context.Context, *accesscontrol.GrantUserAccessCommand) (bool, error)
	RevokeUserAccess(context.Context, *accesscontrol.RevokeUserAccessCommand) error

	// member query
	GetOrgMemberList(context.Context, *accesscontrol.GetOrgMemberListQuery) ([]*accesscontrol.OrgMember, error)
	GetOrgMemberById(context.Context, *accesscontrol.GetOrgMemberByIdQuery) (*accesscontrol.OrgMember, error)
	GetOrgMemberByUserId(ctx context.Context, query *accesscontrol.GetOrgMemberByUserIdQuery) (*accesscontrol.OrgMember, error)
	GetWorkspaceMemberList(context.Context, *accesscontrol.GetWorkspaceMemberListQuery) ([]*accesscontrol.WorkspaceMember, error)
	GetWorkspaceMemberById(context.Context, *accesscontrol.GetWorkspaceMemberByIdQuery) (*accesscontrol.WorkspaceMember, error)
	GetWorkspaceMemberByUserId(ctx context.Context, query *accesscontrol.GetWorkspaceMemberByUserIdQuery) (*accesscontrol.WorkspaceMember, error)
	GetProjectMemberList(context.Context, *accesscontrol.GetProjectMemberListQuery) ([]*accesscontrol.ProjectMember, error)
	GetProjectMemberById(context.Context, *accesscontrol.GetProjectMemberByIdQuery) (*accesscontrol.ProjectMember, error)
	GetProjectMemberByUserId(context.Context, *accesscontrol.GetProjectMemberByUserIdQuery) (*accesscontrol.ProjectMember, error)

	// member mutation
	CreateOrgMember(context.Context, *accesscontrol.CreateOrgMemberCommand) (*accesscontrol.OrgMember, error)
	UpdateOrgMemberById(context.Context, *accesscontrol.UpdateOrgMemberByIdCommand) error
	DeleteOrgMemberById(context.Context, *accesscontrol.DeleteOrgMemberByIdCommand) error
	CreateWorkspaceMember(context.Context, *accesscontrol.CreateWorkspaceMemberCommand) (*accesscontrol.WorkspaceMember, error)
	UpdateWorkspaceMemberById(context.Context, *accesscontrol.UpdateWorkspaceMemberByIdCommand) error
	DeleteWorkspaceMemberById(context.Context, *accesscontrol.DeleteWorkspaceMemberByIdCommand) error
	CreateProjectMember(context.Context, *accesscontrol.CreateProjectMemberCommand) (*accesscontrol.ProjectMember, error)
	UpdateProjectMemberById(context.Context, *accesscontrol.UpdateProjectMemberByIdCommand) error
	DeleteProjectMemberById(context.Context, *accesscontrol.DeleteProjectMemberByIdCommand) error

	// user role query and mutation
	GetRoleList(context.Context) ([]*accesscontrol.Role, error)
	GetRoleById(context.Context, *accesscontrol.GetRoleByIdQuery) (*accesscontrol.Role, error)
	CreateRole(context.Context, *accesscontrol.CreateRoleCommand) (*accesscontrol.Role, error)
	DeleteRoleById(context.Context, *accesscontrol.DeleteRoleByIdCommand) error
}

type accessControlStore struct {
	Conn *pgx.Conn
}

func NewAccessControlStore(conn *pgx.Conn) *accessControlStore {
	return &accessControlStore{
		Conn: conn,
	}
}

// type AccessControlStore interface {
// 	// access control query
// 	Evaluate(context.Context, *accesscontrol.EvaluateQuery) (bool, error)
// 	GetUserAssignmentList(context.Context, *accesscontrol.GetUserAssignmentListQuery) ([]accesscontrol.UserAssignment, error)
// 	GetUserAssignmentById(context.Context, *accesscontrol.GetUserAssignmentByIdQuery) (*accesscontrol.UserAssignment, error)

// 	// access control mutation
// 	CreateUserAssignment(context.Context, *accesscontrol.CreateUserAssignmentCommand) (*accesscontrol.UserAssignment, error)
// 	UpdateUserAssignmentById(context.Context, *accesscontrol.UpdateUserAssignmentByIdCommand) error
// 	DeleteUserAssignmentById(context.Context, *accesscontrol.DeleteUserAssignmentByIdCommand) error

// 	// user group query
// 	GetUserGroupList(context.Context, *accesscontrol.GetUserGroupListQuery) ([]*accesscontrol.UserGroup, error)
// 	GetUserGroupById(context.Context, *accesscontrol.GetUserGroupByIdQuery) (*accesscontrol.UserGroup, error)

// 	// user group mutation
// 	CreateUserGroup(context.Context, *accesscontrol.CreateUserGroupCommand) (*accesscontrol.UserGroup, error)
// 	UpdateUserGroupById(context.Context, *accesscontrol.UpdateUserGroupByIdCommand) error
// 	DeleteUserGroupById(context.Context, *accesscontrol.DeleteUserGroupByIdCommand) error

// 	// role query
// 	GetRoleList(context.Context, *accesscontrol.GetRoleListQuery) ([]*accesscontrol.Role, error)
// 	GetRoleById(context.Context, *accesscontrol.GetRoleByIdQuery) (*accesscontrol.Role, error)

// 	// role mutation
// 	CreateRole(context.Context, *accesscontrol.CreateRoleCommand) (*accesscontrol.Role, error)
// 	//UpdateRoleById(context.Context, *accesscontrol.UpdateRoleByIdCommand) error
// 	DeleteRoleById(context.Context, *accesscontrol.DeleteRoleByIdCommand) error
// }

const testQuery = `
SELECT 
	COUNT(*) FILTER(WHERE "orgId" = 2) as org_exists,
	COUNT(*) FILTER(WHERE "orgId" = 2 AND "workspaceId" = COALESCE(2, "workspaceId")) as ws_exists,
	COUNT(*) FILTER(WHERE "orgId" = 2 AND "workspaceId" = 2 AND "projectId" = COALESCE(2, "projectId")) as project_exists
FROM
    system."UserAssignment";
`
const test2Query = `
SELECT 
	COUNT(*) FILTER(WHERE ua."orgId" = 2) as org_exists,
	COUNT(*) FILTER(WHERE ua."orgId" = 2 AND ua."workspaceId" = NULLIF(2, 0)) as ws_exists,
	COUNT(*) FILTER(WHERE ua."orgId" = 2 AND ua."workspaceId" = 2 AND ua."projectId" = NULLIF(2, 0)) as project_exists,
	ug."name" AS user_group,
	ur."name" AS role,
	rp."name" AS permission
FROM
    "system"."UserAssignment" AS ua
LEFT JOIN "system"."UserGroup" AS ug ON ug.id = ua."userGroupId"	
LEFT JOIN "system"."Role" AS ur ON ur.id = ua."userRoleId"	
LEFT JOIN "system"."Permission" AS rp ON rp."roleId" = ur.id
WHERE 
	ua."orgId" = 2 
	AND (
		ua."userId" = NULLIF(0, 0) OR ua."userGroupId" = NULLIF(1, 0)
	)
GROUP BY ua.id, ug.id, ur.id, rp.id;
`
const q3 = `
--EXPLAIN ANALYZE
SELECT 
    --COUNT(*) FILTER(WHERE ua."orgId" = 2) AS org_exists,
    --COUNT(*) FILTER(WHERE ua."orgId" = 2 AND ua."workspaceId" = NULLIF(2, 0)) AS ws_exists,
    --COUNT(*) FILTER(WHERE ua."orgId" = 2 AND ua."workspaceId" = 2 AND ua."projectId" = NULLIF(2, 0)) AS project_exists,
    --COUNT(*) FILTER(WHERE ua."userId" = NULLIF(0, 0) OR ua."userGroupId" = NULLIF(1, 0)) AS user_group_1,
	ua."userId" AS user_id,
    ug."name" AS user_group,
    ur."name" AS role,
    --rp."name" AS permission,
	COALESCE(json_agg(rp."name")::jsonb, '[]'::jsonb) AS permissions
FROM 
    "system"."UserAssignment" AS ua
LEFT JOIN "system"."UserGroup" AS ug ON ug.id = ua."userGroupId" AND ug."userId" = ua."userId"
LEFT JOIN "system"."Role" AS ur ON ur.id = ua."userRoleId"
LEFT JOIN "system"."Permission" AS rp ON rp."roleId" = ur.id
WHERE 
    ua."orgId" = 2
    AND (ua."workspaceId" = 2 OR 0 = 0)
    AND (ua."projectId" = 2 OR 0 = 0)
    AND (ua."userId" = 2 OR 0 = 0)
    AND (ua."userGroupId" = 2 OR 0 = 0)
GROUP BY 
    ua.id,
    ug.id,
    ur.id
    --rp.id;
`

const q4 = `
--EXPLAIN ANALYZE
WITH ugg AS (SELECT * from "system"."UserGroup" WHERE "orgId" = 2 AND "users" @> '[]')
SELECT 
    --COUNT(*) FILTER(WHERE ua."orgId" = 2) AS org_exists,
    --COUNT(*) FILTER(WHERE ua."orgId" = 2 AND ua."workspaceId" = NULLIF(2, 0)) AS ws_exists,
    --COUNT(*) FILTER(WHERE ua."orgId" = 2 AND ua."workspaceId" = 2 AND ua."projectId" = NULLIF(2, 0)) AS project_exists,
    --COUNT(*) FILTER(WHERE ua."userId" = NULLIF(0, 0) OR ua."userGroupId" = NULLIF(1, 0)) AS user_group_1,
	(SELECT "name" FROM ugg) AS ooooo,
	ua."userId" AS user_id,
    ug."name" AS user_group,
    ur."name" AS role,
    --rp."name" AS permission,
	COALESCE(json_agg(rp."name")::jsonb, '[]'::jsonb) AS permissions
FROM 
    "system"."UserAssignment" AS ua
LEFT JOIN "system"."UserGroup" AS ug ON ug.id = ua."userGroupId" AND ug."users" @> '[]'
LEFT JOIN "system"."Role" AS ur ON ur.id = ua."userRoleId"
LEFT JOIN "system"."Permission" AS rp ON rp."roleId" = ur.id
WHERE 
    ua."orgId" = 2
    AND (ua."workspaceId" = 2 OR 0 = 0)
    AND (ua."projectId" = 2 OR 0 = 0)
    AND (ua."userId" = 2 OR 0 = 0)
GROUP BY 
    ua.id,
    ug.id,
    ur.id
    --rp.id;
`

const q5 = `
--EXPLAIN ANALYZE
SELECT
	ua."orgId",
	ua."workspaceId",
	ua."projectId",
	ua."userId" AS user_id,
    ug."name" AS user_group,
    ur."name" AS role,
	COALESCE(json_agg(rp."name")::jsonb, '[]'::jsonb) AS permissions
FROM 
    "system"."UserAssignment" AS ua
LEFT JOIN "system"."UserGroup" AS ug ON ug.id = ua."userGroupId"
LEFT JOIN "system"."Role" AS ur ON ur.id = ua."userRoleId"
LEFT JOIN "system"."Permission" AS rp ON rp."roleId" = ur.id
WHERE (
	ua."orgId" = 2 
	AND (ua."workspaceId" = 2 OR 0 = 0) 
	AND (ua."projectId" = 2 OR 0 = 0)
	)
	AND(
		(ua."userId" = 1 OR 10 = 0) 
		OR ug."users" @> '[100]'
	)
GROUP BY 
    ua.id,
    ug.id,
    ur.id;
`
const q6 = `
SELECT jsonb_array_length(permissions)::bool FROM (
	SELECT
		COALESCE(json_agg(rp."name")::jsonb, '[]'::jsonb) AS permissions
	FROM 
		"system"."UserAssignment" AS ua
	LEFT JOIN "system"."UserGroup" AS ug ON ug.id = ua."userGroupId"
	LEFT JOIN "system"."Role" AS ur ON ur.id = ua."userRoleId"
	LEFT JOIN "system"."Permission" AS rp ON rp."roleId" = ur.id
	WHERE 
	(
		ua."orgId" = 2 AND (ua."workspaceId" = 2 OR 0 = 0) AND (ua."projectId" = 2 OR 0 = 0)
	)
	AND
	(
		(ua."userId" = 100 OR 10 = 0) OR ug."users" @> '[10]'
	)
	AND
	(
		rp."name" = 'org:create'
	)
) AS permission_exists
`
