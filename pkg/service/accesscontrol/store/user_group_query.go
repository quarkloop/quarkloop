package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

/// GetUserGroupList

const listUserGroupsQuery = `
SELECT 
    "id",
    "orgId",
    "name",
	"users",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."UserGroup"
WHERE 
    "orgId" = @orgId;
`

func (store *accessControlStore) GetUserGroupList(ctx context.Context, query *accesscontrol.GetUserGroupListQuery) ([]*accesscontrol.UserGroup, error) {
	rows, err := store.Conn.Query(ctx, listUserGroupsQuery, pgx.NamedArgs{"orgId": query.OrgId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var aclList []*accesscontrol.UserGroup = []*accesscontrol.UserGroup{}

	for rows.Next() {
		var ug accesscontrol.UserGroup
		err := rows.Scan(
			&ug.Id,
			&ug.OrgId,
			&ug.Name,
			&ug.Users,
			&ug.CreatedAt,
			&ug.CreatedBy,
			&ug.UpdatedAt,
			&ug.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		aclList = append(aclList, &ug)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return aclList, nil
}

/// GetUserGroupById

const getUserGroupByIdQuery = `
SELECT 
    "id",
    "orgId",
    "name",
	"users",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."UserGroup"
WHERE 
    "orgId" = @orgId
AND
    "id" = @id;
`

func (store *accessControlStore) GetUserGroupById(ctx context.Context, query *accesscontrol.GetUserGroupByIdQuery) (*accesscontrol.UserGroup, error) {
	row := store.Conn.QueryRow(ctx, getUserGroupByIdQuery, pgx.NamedArgs{
		"orgId": query.OrgId,
		"id":    query.UserGroupId,
	})

	var ug accesscontrol.UserGroup
	err := row.Scan(
		&ug.Id,
		&ug.OrgId,
		&ug.Name,
		&ug.Users,
		&ug.CreatedAt,
		&ug.CreatedBy,
		&ug.UpdatedAt,
		&ug.UpdatedBy,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, accesscontrol.ErrUserGroupNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &ug, nil
}
