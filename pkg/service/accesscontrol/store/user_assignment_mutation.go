package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
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
	"userId",
    "userGroupId",
    "userRoleId",
    "createdBy"
)
VALUES (
	NULLIF(@orgId, 0),
    NULLIF(@workspaceId, 0),
    NULLIF(@projectId, 0),
	NULLIF(@userId, 0),
	NULLIF(@userGroupId, 0),
	NULLIF(@userRoleId, 0),
    @createdBy
)
RETURNING 
    "id",
    "orgId",
    "workspaceId",
    "projectId",
	"userId",
    "userGroupId",
    "userRoleId",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy";
`

func (store *accessControlStore) CreateUserAssignment(ctx context.Context, cmd *accesscontrol.CreateUserAssignmentCommand) (*accesscontrol.UserAssignment, error) {
	row := store.Conn.QueryRow(ctx, createUserAssignmentQuery, pgx.NamedArgs{
		"orgId":       cmd.OrgId,
		"workspaceId": cmd.WorkspaceId,
		"projectId":   cmd.ProjectId,
		"userId":      cmd.UserId,
		"userGroupId": cmd.UserGroupId,
		"userRoleId":  cmd.UserRoleId,
		"createdBy":   cmd.CreatedBy,
	})

	var ua accesscontrol.UserAssignment
	rowErr := row.Scan(
		&ua.Id,
		&ua.OrgId,
		&ua.WorkspaceId,
		&ua.ProjectId,
		&ua.UserId,
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
    "userRoleId" = @userRoleId,
    "updatedAt"  = @updatedAt,
    "updatedBy"  = @updatedBy,
WHERE
    "id" = @id;
`

func (store *accessControlStore) UpdateUserAssignmentById(ctx context.Context, cmd *accesscontrol.UpdateUserAssignmentByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updateUserAssignmentByIdQuery, pgx.NamedArgs{
		"id":         cmd.UserAssignmentId,
		"userRoleId": cmd.UserRole.UserRoleId,
		"updatedBy":  cmd.UserRole.UpdatedBy,
		"updatedAt":  time.Now(),
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
    "id" = @id
%s;
`

func (store *accessControlStore) DeleteUserAssignmentById(ctx context.Context, cmd *accesscontrol.DeleteUserAssignmentByIdCommand) error {
	var whereClause []string
	if cmd.OrgId != 0 {
		whereClause = append(whereClause, `AND "orgId" = @orgId`)
	}
	if cmd.OrgId != 0 && cmd.WorkspaceId != 0 {
		whereClause = append(whereClause, `"workspaceId" = @workspaceId`)
	}
	if cmd.OrgId != 0 && cmd.WorkspaceId != 0 && cmd.ProjectId != 0 {
		whereClause = append(whereClause, `"projectId" = @projectId`)
	}

	finalQuery := fmt.Sprintf(deleteUserAssignmentByIdQuery, strings.Join(whereClause, " AND "))

	commandTag, err := store.Conn.Exec(ctx, finalQuery, pgx.NamedArgs{
		"id":          cmd.UserAssignmentId,
		"orgId":       cmd.OrgId,
		"workspaceId": cmd.WorkspaceId,
		"projectId":   cmd.ProjectId,
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
