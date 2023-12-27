package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

/// Evaluate

const evaluateQuery = `
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
		((@orgId = 0 AND ua."orgId" IS NULL) OR ua."orgId" = @orgId) 
		AND
		((@workspaceId = 0 AND ua."workspaceId" IS NULL) OR ua."workspaceId" = @workspaceId)
		AND
		((@projectId = 0 AND ua."projectId" IS NULL) OR ua."projectId" = @projectId)		
	)
	AND
	(
		((@userId = 0 AND ua."userId" IS NULL) OR ua."userId" = @userId) 
		OR
		ug."users" @> '[@userId]'
	)
	AND
	(
		rp."name" = @permission
	)
) AS permission_exists
`

func (store *accessControlStore) Evaluate(ctx context.Context, permission string, orgId, workspaceId, projectId, userId int) (bool, error) {
	row := store.Conn.QueryRow(ctx, evaluateQuery, pgx.NamedArgs{
		"orgId":       orgId,
		"workspaceId": workspaceId,
		"projectId":   projectId,
		"userId":      userId,
		"permission":  permission,
	})

	var permission_exists bool
	err := row.Scan(&permission_exists)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return false, err
	}

	return permission_exists, nil
}
