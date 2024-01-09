package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

type AccessControlStore interface {
	// access control query
	Evaluate(context.Context, *accesscontrol.EvaluateQuery) (bool, error)
	GetUserAssignmentList(context.Context, *accesscontrol.GetUserAssignmentListQuery) ([]accesscontrol.UserAssignment, error)
	GetUserAssignmentById(context.Context, *accesscontrol.GetUserAssignmentByIdQuery) (*accesscontrol.UserAssignment, error)

	// access control mutation
	CreateUserAssignment(context.Context, *accesscontrol.CreateUserAssignmentCommand) (*accesscontrol.UserAssignment, error)
	UpdateUserAssignmentById(context.Context, *accesscontrol.UpdateUserAssignmentByIdCommand) error
	DeleteUserAssignmentById(context.Context, *accesscontrol.DeleteUserAssignmentByIdCommand) error

	// user group query
	GetUserGroupList(context.Context, *accesscontrol.GetUserGroupListQuery) ([]*accesscontrol.UserGroup, error)
	GetUserGroupById(context.Context, *accesscontrol.GetUserGroupByIdQuery) (*accesscontrol.UserGroup, error)

	// user group mutation
	CreateUserGroup(context.Context, *accesscontrol.CreateUserGroupCommand) (*accesscontrol.UserGroup, error)
	UpdateUserGroupById(context.Context, *accesscontrol.UpdateUserGroupByIdCommand) error
	DeleteUserGroupById(context.Context, *accesscontrol.DeleteUserGroupByIdCommand) error

	// role query
	GetUserRoleList(context.Context, *accesscontrol.GetUserRoleListQuery) ([]*accesscontrol.UserRole, error)
	GetUserRoleById(context.Context, *accesscontrol.GetUserRoleByIdQuery) (*accesscontrol.UserRole, error)

	// role mutation
	CreateUserRole(context.Context, *accesscontrol.CreateUserRoleCommand) (*accesscontrol.UserRole, error)
	//UpdateUserRoleById(context.Context, *accesscontrol.UpdateUserRoleByIdCommand) error
	DeleteUserRoleById(context.Context, *accesscontrol.DeleteUserRoleByIdCommand) error
}

type accessControlStore struct {
	Conn *pgx.Conn
}

func NewAccessControlStore(conn *pgx.Conn) *accessControlStore {
	return &accessControlStore{
		Conn: conn,
	}
}

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
LEFT JOIN "system"."UserRole" AS ur ON ur.id = ua."userRoleId"	
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
LEFT JOIN "system"."UserRole" AS ur ON ur.id = ua."userRoleId"
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
LEFT JOIN "system"."UserRole" AS ur ON ur.id = ua."userRoleId"
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
LEFT JOIN "system"."UserRole" AS ur ON ur.id = ua."userRoleId"
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
	LEFT JOIN "system"."UserRole" AS ur ON ur.id = ua."userRoleId"
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
