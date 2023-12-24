package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

/// ListUserAssignments

const listUserAssignmentsQuery = `
SELECT 
  "id", "orgId", "workspaceId", "projectId", "userGroupId", "userRoleId",
  "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM 
  "system"."UserAssignment"
WHERE
  "orgId" = @orgId;
`

func (store *accessControlStore) ListUserAssignments(ctx context.Context, orgId int) ([]accesscontrol.UserAssignment, error) {
	rows, err := store.Conn.Query(ctx, listUserAssignmentsQuery, pgx.NamedArgs{"orgId": orgId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var aclList []accesscontrol.UserAssignment = []accesscontrol.UserAssignment{}

	for rows.Next() {
		var ua accesscontrol.UserAssignment
		err := rows.Scan(
			&ua.Id,
			&ua.OrgId,
			&ua.WorkspaceId,
			&ua.ProjectId,
			&ua.UserGroupId,
			&ua.UserRoleId,
			&ua.CreatedAt,
			&ua.CreatedBy,
			&ua.UpdatedAt,
			&ua.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		aclList = append(aclList, ua)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return aclList, nil
}

/// GetUserAssignmentById

const getUserAssignmentByIdQuery = `
SELECT 
  "id", "orgId", "workspaceId", "projectId", "userGroupId", "userRoleId",
  "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM 
  "system"."UserAssignment" 
WHERE 
  "id" = @id;
`

func (store *accessControlStore) GetUserAssignmentById(ctx context.Context, userAssignmentId int) (*accesscontrol.UserAssignment, error) {
	row := store.Conn.QueryRow(ctx, getUserAssignmentByIdQuery, pgx.NamedArgs{"id": userAssignmentId})

	var ua accesscontrol.UserAssignment
	err := row.Scan(
		&ua.Id,
		&ua.OrgId,
		&ua.WorkspaceId,
		&ua.ProjectId,
		&ua.UserGroupId,
		&ua.UserRoleId,
		&ua.CreatedAt,
		&ua.CreatedBy,
		&ua.UpdatedAt,
		&ua.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &ua, nil
}
