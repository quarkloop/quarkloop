package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// CreateWorkspace

const createWorkspaceMutation = `
INSERT INTO
  "system"."Workspace" ("orgId", "sid", "name", "description", "accessType", "createdBy")
VALUES
  (@orgId, @sid, @name, @description, @accessType, @createdBy)
RETURNING 
  "id", "sid", "name", "description", "accessType", "createdAt", "createdBy";
`

func (r *Repository) CreateWorkspace(ctx context.Context, orgId int, workspace *model.Workspace) (*model.Workspace, error) {
	if workspace.ScopedId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		workspace.ScopedId = sid
	}

	row := r.SystemDbConn.QueryRow(
		ctx,
		createWorkspaceMutation,
		pgx.NamedArgs{
			"orgId":       orgId,
			"sid":         workspace.ScopedId,
			"name":        workspace.Name,
			"description": workspace.Description,
			"accessType":  workspace.AccessType,
			"createdBy":   workspace.CreatedBy,
		},
	)

	var ws model.Workspace
	rowErr := row.Scan(
		&ws.Id,
		&ws.ScopedId,
		&ws.Name,
		&ws.Description,
		&ws.AccessType,
		&ws.CreatedAt,
		&ws.CreatedBy,
	)
	if rowErr != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", rowErr)
		return nil, rowErr
	}

	return &ws, nil
}

/// UpdateWorkspaceById

const updateWorkspaceByIdMutation = `
UPDATE
  "system"."Workspace"
SET
  "sid"         = @sid,
  "name"        = @name,
  "description" = @description,
  "accessType"  = @accessType,
  "updatedAt"   = @updatedAt,
  "updatedBy"   = @updatedBy,
WHERE
  "id" = @id;
`

func (r *Repository) UpdateWorkspaceById(ctx context.Context, workspaceId int, workspace *model.Workspace) error {
	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		updateWorkspaceByIdMutation,
		pgx.NamedArgs{
			"id":          workspaceId,
			"sid":         workspace.ScopedId,
			"name":        workspace.Name,
			"description": workspace.Description,
			"accessType":  *workspace.AccessType,
			"updatedAt":   time.Now(),
			"updatedBy":   workspace.UpdatedBy,
		},
	)
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

/// DeleteWorkspaceById

const deleteWorkspaceByIdMutation = `
DELETE FROM
  "system"."Workspace"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteWorkspaceById(ctx context.Context, workspaceId int) error {
	commandTag, err := r.SystemDbConn.Exec(ctx, deleteWorkspaceByIdMutation, pgx.NamedArgs{"id": workspaceId})
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
