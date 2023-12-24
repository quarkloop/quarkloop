package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/accesscontrol"
)

/// CreateUserRole

const createUserRoleQuery = `
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
    "updatedBy";
`

func (store *accessControlStore) CreateUserRole(ctx context.Context, orgId int, userRole *accesscontrol.UserRole) (*accesscontrol.UserRole, error) {
	row := store.Conn.QueryRow(ctx, createUserRoleQuery, pgx.NamedArgs{
		"orgId":     userRole.OrgId,
		"name":      userRole.Name,
		"createdBy": userRole.CreatedBy,
	})

	var ur accesscontrol.UserRole
	rowErr := row.Scan(
		&ur.Id,
		&ur.OrgId,
		&ur.Name,
		&ur.CreatedAt,
		&ur.CreatedBy,
		&ur.UpdatedAt,
		&ur.UpdatedBy,
	)
	if rowErr != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", rowErr)
		return nil, rowErr
	}

	return &ur, nil
}

/// UpdateUserRoleById

const updateUserRoleByIdQuery = `
UPDATE
  "system"."UserRole"
SET
  "name"        = @name,
  "updatedAt"   = @updatedAt,
  "updatedBy"   = @updatedBy,
WHERE
  "id" = @id;
`

func (store *accessControlStore) UpdateUserRoleById(ctx context.Context, userRoleId int, userRole *accesscontrol.UserRole) error {
	commandTag, err := store.Conn.Exec(ctx, updateUserRoleByIdQuery, pgx.NamedArgs{
		"id":        userRoleId,
		"name":      userRole.Name,
		"updatedAt": time.Now(),
		"updatedBy": userRole.UpdatedBy,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to update")
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}

/// DeleteUserRoleById

const deleteUserRoleByIdQuery = `
DELETE FROM
  "system"."UserRole"
WHERE
  "orgId" = @orgId
AND
  "id" = @id;
`

func (store *accessControlStore) DeleteUserRoleById(ctx context.Context, orgId int, userRoleId int) error {
	commandTag, err := store.Conn.Exec(ctx, deleteUserRoleByIdQuery, pgx.NamedArgs{
		"orgId": orgId,
		"id":    userRoleId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to delete")
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}
