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

/// CreateUserAssignment

const createUserAssignmentQuery = `
INSERT INTO "system"."UserAssignment" (
    "orgId",
    "workspaceId",
    "projectId",
    "userGroupId",
    "userRoleId",
    "createdBy"
)
VALUES (
    @orgId,
    NULLIF(@workspaceId, 0),
    NULLIF(@projectId, 0),
    @userGroupId,
    @userRoleId,
    @createdBy
)
RETURNING 
    "id",
    "orgId",
    "workspaceId",
    "projectId",
    "userGroupId",
    "userRoleId",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy";
`

func (store *accessControlStore) CreateUserAssignment(ctx context.Context, orgId int, userRole *accesscontrol.UserAssignment) (*accesscontrol.UserAssignment, error) {
	row := store.Conn.QueryRow(ctx, createUserAssignmentQuery, pgx.NamedArgs{
		"orgId":       userRole.OrgId,
		"workspaceId": userRole.WorkspaceId,
		"projectId":   userRole.ProjectId,
		"userGroupId": userRole.UserGroupId,
		"userRoleId":  userRole.UserRoleId,
		"createdBy":   userRole.CreatedBy,
	})

	var ua accesscontrol.UserAssignment
	rowErr := row.Scan(
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
	if rowErr != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", rowErr)
		return nil, rowErr
	}

	return &ua, nil
}

/// UpdateUserAssignmentById

const updateUserAssignmentByIdQuery = `
UPDATE
    "system"."UserAssignment"
SET
    "userRoleId"  = @userRoleId,
    "updatedAt"   = @updatedAt,
    "updatedBy"   = @updatedBy,
WHERE
    "id" = @id;
`

func (store *accessControlStore) UpdateUserAssignmentById(ctx context.Context, userAssignmentId int, userRole *accesscontrol.UserAssignment) error {
	commandTag, err := store.Conn.Exec(ctx, updateUserAssignmentByIdQuery, pgx.NamedArgs{
		"id":         userAssignmentId,
		"userRoleId": userRole.UserRoleId,
		"updatedAt":  time.Now(),
		"updatedBy":  userRole.UpdatedBy,
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

/// DeleteUserAssignmentById

const deleteUserAssignmentByIdQuery = `
DELETE FROM
    "system"."UserAssignment"
WHERE
    "orgId" = @orgId
AND
    "id" = @id;
`

func (store *accessControlStore) DeleteUserAssignmentById(ctx context.Context, orgId int, userAssignmentId int) error {
	commandTag, err := store.Conn.Exec(ctx, deleteUserAssignmentByIdQuery, pgx.NamedArgs{
		"orgId": orgId,
		"id":    userAssignmentId,
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
