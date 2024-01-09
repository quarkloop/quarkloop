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
    r."id",
    r."orgId",
    r."name" AS role,
	perm."name" AS permission,
    r."createdAt",
    r."createdBy",
    r."updatedAt",
    r."updatedBy"
FROM 
    "system"."UserRole" AS r
LEFT JOIN "system"."Permission" AS perm ON perm."roleId" = r.id
WHERE 
    r."orgId" = @orgId;
`

func (store *accessControlStore) GetUserRoleList(ctx context.Context, query *accesscontrol.GetUserRoleListQuery) ([]*accesscontrol.UserRole, error) {
	rows, err := store.Conn.Query(ctx, listUserRolesQuery, pgx.NamedArgs{"orgId": query.OrgId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	permissions := map[string]*accesscontrol.UserRole{}
	for rows.Next() {
		var ur accesscontrol.UserRole
		var permission string
		err := rows.Scan(
			&ur.Id,
			&ur.OrgId,
			&ur.Name,
			&permission,
			&ur.CreatedAt,
			&ur.CreatedBy,
			&ur.UpdatedAt,
			&ur.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		_, ok := permissions[ur.Name]
		if !ok {
			permissions[ur.Name] = &accesscontrol.UserRole{
				Id:          ur.Id,
				OrgId:       ur.OrgId,
				Name:        ur.Name,
				Permissions: []string{},
				CreatedAt:   ur.CreatedAt,
				CreatedBy:   ur.CreatedBy,
				UpdatedAt:   ur.UpdatedAt,
				UpdatedBy:   ur.UpdatedBy,
			}
		}
		permissions[ur.Name].Permissions = append(permissions[ur.Name].Permissions, permission)
	}

	var roleList []*accesscontrol.UserRole = []*accesscontrol.UserRole{}
	for _, role := range permissions {
		roleList = append(roleList, role)
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
    r."id",
    r."orgId",
    r."name" AS role,
	perm."name" AS permission,
    r."createdAt",
    r."createdBy",
    r."updatedAt",
    r."updatedBy"
FROM 
    "system"."UserRole" AS r
LEFT JOIN "system"."Permission" AS perm ON perm."roleId" = r.id
WHERE 
    r."orgId" = @orgId
AND	
    r."id" = @id;
`

func (store *accessControlStore) GetUserRoleById(ctx context.Context, query *accesscontrol.GetUserRoleByIdQuery) (*accesscontrol.UserRole, error) {
	rows, err := store.Conn.Query(ctx, getUserRoleByIdQuery, pgx.NamedArgs{
		"orgId": query.OrgId,
		"id":    query.UserRoleId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	permissions := map[string]*accesscontrol.UserRole{}
	for rows.Next() {
		var ur accesscontrol.UserRole
		var permission string
		err := rows.Scan(
			&ur.Id,
			&ur.OrgId,
			&ur.Name,
			&permission,
			&ur.CreatedAt,
			&ur.CreatedBy,
			&ur.UpdatedAt,
			&ur.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}
		_, ok := permissions[ur.Name]
		if !ok {
			permissions[ur.Name] = &accesscontrol.UserRole{
				Id:          ur.Id,
				OrgId:       ur.OrgId,
				Name:        ur.Name,
				Permissions: []string{},
				CreatedAt:   ur.CreatedAt,
				CreatedBy:   ur.CreatedBy,
				UpdatedAt:   ur.UpdatedAt,
				UpdatedBy:   ur.UpdatedBy,
			}
		}
		permissions[ur.Name].Permissions = append(permissions[ur.Name].Permissions, permission)
	}
	if len(permissions) > 1 {
		panic("only a single role should be returned")
	}

	var role *accesscontrol.UserRole
	for _, r := range permissions {
		role = r
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}
	if role == nil {
		return nil, accesscontrol.ErrRoleNotFound
	}

	return role, nil
}
