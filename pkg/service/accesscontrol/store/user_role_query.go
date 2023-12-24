package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

/// ListUserRoles

const listUserRolesQuery = `
SELECT 
    "id",
    "orgId",
    "name",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."UserRole"
WHERE 
    "orgId" = @orgId;
`

func (store *accessControlStore) ListUserRoles(ctx context.Context, orgId int) ([]accesscontrol.UserRole, error) {
	rows, err := store.Conn.Query(ctx, listUserRolesQuery, pgx.NamedArgs{"orgId": orgId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var roleList []accesscontrol.UserRole = []accesscontrol.UserRole{}

	for rows.Next() {
		var ur accesscontrol.UserRole
		err := rows.Scan(
			&ur.Id,
			&ur.OrgId,
			&ur.Name,
			&ur.CreatedAt,
			&ur.CreatedBy,
			&ur.UpdatedAt,
			&ur.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		roleList = append(roleList, ur)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return roleList, nil
}

/// GetUserRoleById

const getUserRoleByIdQuery = `
SELECT 
    "id",
    "orgId",
    "name",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."UserRole"
WHERE 
    "id" = @id;
`

func (store *accessControlStore) GetUserRoleById(ctx context.Context, userRoleId int) (*accesscontrol.UserRole, error) {
	row := store.Conn.QueryRow(ctx, getUserRoleByIdQuery, pgx.NamedArgs{"id": userRoleId})

	var ur accesscontrol.UserRole
	err := row.Scan(
		&ur.Id,
		&ur.OrgId,
		&ur.Name,
		&ur.CreatedAt,
		&ur.CreatedBy,
		&ur.UpdatedAt,
		&ur.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &ur, nil
}
