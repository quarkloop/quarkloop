package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

/// CreateUserRole

const createUserRoleQuery = `
WITH 
role AS (
    INSERT INTO "system"."UserRole" (
		"orgId", 
		"name", 
		"createdBy"
	)
    VALUES (
		@orgId, 
		@name, 
		@createdBy
	)
    RETURNING 
	    "id",
        "orgId",
        "name",
        "createdAt",
        "createdBy",
        "updatedAt",
        "updatedBy"
), 
permission AS (
    INSERT INTO "system"."Permission" ("roleId", "name") VALUES %s
)
SELECT
    "id",
    "orgId",
    "name",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    role;
`

func (store *accessControlStore) CreateUserRole(ctx context.Context, cmd *accesscontrol.CreateUserRoleCommand) (*accesscontrol.UserRole, error) {
	permissions := []string{}
	for _, perm := range cmd.Permissions {
		permissions = append(permissions, fmt.Sprintf("((SELECT id FROM role),'%s')", perm.Name))
	}
	finalQuery := fmt.Sprintf(createUserRoleQuery, strings.Join(permissions, ","))

	row := store.Conn.QueryRow(ctx, finalQuery, pgx.NamedArgs{
		"orgId":     cmd.OrgId,
		"name":      cmd.Name,
		"createdBy": cmd.CreatedBy,
	})

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
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	return &ur, nil
}

// /// UpdateUserRoleById

// const updateUserRoleByIdQuery = `
// UPDATE
//     "system"."UserRole"
// SET
//     "name"        = @name,
//     "updatedAt"   = @updatedAt,
//     "updatedBy"   = @updatedBy,
// WHERE
//     "id" = @id;
// `

// func (store *accessControlStore) UpdateUserRoleById(ctx context.Context, cmd *accesscontrol.UpdateUserRoleByIdCommand) error {
// 	commandTag, err := store.Conn.Exec(ctx, updateUserRoleByIdQuery, pgx.NamedArgs{
// 		"id":        cmd.UserRoleId,
// 		"name":      cmd.UserRole.Name,
// 		"updatedBy": cmd.UserRole.UpdatedBy,
// 		"updatedAt": time.Now(),
// 	})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
// 		return err
// 	}

// 	if commandTag.RowsAffected() != 1 {
// 		notFoundErr := errors.New("cannot find to update")
// 		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", notFoundErr)
// 		return notFoundErr
// 	}

// 	return nil
// }

/// DeleteUserRoleById

const deleteUserRoleByIdQuery = `
WITH 
permission AS (
	DELETE FROM
       "system"."UserRole"
	WHERE
		"orgId" = @orgId
	AND
		"id" = @id
	RETURNING "id"
)
DELETE FROM
    "system"."Permission"
WHERE
    "roleId" = (SELECT "id" FROM permission);
`

func (store *accessControlStore) DeleteUserRoleById(ctx context.Context, cmd *accesscontrol.DeleteUserRoleByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteUserRoleByIdQuery, pgx.NamedArgs{
		"orgId": cmd.OrgId,
		"id":    cmd.UserRoleId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() == 0 {
		err = errors.New("no rows has been deleted")
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", err)
		return err
	}

	return nil
}
