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

/// CreateUserGroup

const createUserGroupQuery = `
INSERT INTO "system"."UserGroup" (
    "orgId", 
    "name", 
	"users",
    "createdBy"
)
VALUES (
    @orgId, 
    @name, 
	@users,
    @createdBy
)
RETURNING 
    "id",
    "orgId",
    "name",
	"users",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy";
`

func (store *accessControlStore) CreateUserGroup(ctx context.Context, cmd *accesscontrol.CreateUserGroupCommand) (*accesscontrol.UserGroup, error) {
	row := store.Conn.QueryRow(ctx, createUserGroupQuery, pgx.NamedArgs{
		"orgId":     cmd.OrgId,
		"name":      cmd.Name,
		"users":     cmd.Users,
		"createdBy": cmd.CreatedBy,
	})

	var ug accesscontrol.UserGroup
	rowErr := row.Scan(
		&ug.Id,
		&ug.OrgId,
		&ug.Name,
		&ug.Users,
		&ug.CreatedAt,
		&ug.CreatedBy,
		&ug.UpdatedAt,
		&ug.UpdatedBy,
	)
	if rowErr != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", rowErr)
		return nil, rowErr
	}

	return &ug, nil
}

/// UpdateUserGroupById

const updateUserGroupByIdQuery = `
UPDATE
    "system"."UserGroup"
SET
    "name"        = @name,
    "updatedAt"   = @updatedAt,
    "updatedBy"   = @updatedBy,
WHERE
    "id" = @id;
`

func (store *accessControlStore) UpdateUserGroupById(ctx context.Context, cmd *accesscontrol.UpdateUserGroupByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updateUserGroupByIdQuery, pgx.NamedArgs{
		"id":        cmd.UserGroupId,
		"name":      cmd.UserGroup.Name,
		"updatedBy": cmd.UserGroup.UpdatedBy,
		"updatedAt": time.Now(),
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

/// DeleteUserGroupById

const deleteUserGroupByIdQuery = `
DELETE FROM
    "system"."UserGroup"
WHERE
    "orgId" = @orgId
AND
    "id" = @id;
`

func (store *accessControlStore) DeleteUserGroupById(ctx context.Context, cmd *accesscontrol.DeleteUserGroupByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteUserGroupByIdQuery, pgx.NamedArgs{
		"orgId": cmd.OrgId,
		"id":    cmd.UserGroupId,
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
