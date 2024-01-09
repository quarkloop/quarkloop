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
    "id",
    "orgId",
    "workspaceId",
    "projectId",
    "userGroupId",
    "userRoleId",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."UserAssignment"
WHERE 
%s;
`

func (store *accessControlStore) GetUserAssignmentList(ctx context.Context, query *accesscontrol.GetUserAssignmentListQuery) ([]*accesscontrol.UserAssignment, error) {
	var whereClause string
	if query.OrgId != 0 && query.WorkspaceId != 0 && query.ProjectId != 0 {
		whereClause = `"orgId" = @orgId AND "workspaceId" = @workspaceId AND "projectId" = @projectId;`
	} else if query.OrgId != 0 && query.WorkspaceId != 0 {
		whereClause = `"orgId" = @orgId AND "workspaceId" = @workspaceId AND "projectId" is NULL;`
	} else if query.OrgId != 0 {
		whereClause = `"orgId" = @orgId AND "workspaceId" is NULL AND "projectId" is NULL;`
	}
	finalQuery := fmt.Sprintf(listUserAssignmentsQuery, whereClause)

	rows, err := store.Conn.Query(ctx, finalQuery, pgx.NamedArgs{
		"orgId":       query.OrgId,
		"workspaceId": query.WorkspaceId,
		"projectId":   query.ProjectId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var uaList []*accesscontrol.UserAssignment = []*accesscontrol.UserAssignment{}
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
		uaList = append(uaList, &ua)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return uaList, nil
}

/// GetUserAssignmentById

const getUserAssignmentByIdQuery = `
SELECT 
    "id",
    "orgId",
    "workspaceId",
    "projectId",
    "userGroupId",
    "userRoleId",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."UserAssignment"
WHERE 
    "id" = @id;
`

func (store *accessControlStore) GetUserAssignmentById(ctx context.Context, query *accesscontrol.GetUserAssignmentByIdQuery) (*accesscontrol.UserAssignment, error) {
	row := store.Conn.QueryRow(ctx, getUserAssignmentByIdQuery, pgx.NamedArgs{"id": query.UserAssignmentId})

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
