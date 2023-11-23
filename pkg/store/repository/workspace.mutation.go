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
  "system"."Workspace" ("orgId", "id", "name", "accessType", "description", "path")
VALUES
  (@orgId, @id, @name, @accessType, @description, @path)
RETURNING 
  "id", "name", "accessType", "path", "description", "createdAt";
`

func (r *Repository) CreateWorkspace(ctx context.Context, orgId string, workspace *model.Workspace) (*model.Workspace, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	workspace.Id = id
	workspace.Path = fmt.Sprintf("/org/%s/%s", orgId, workspace.Id)

	row := r.SystemDbConn.QueryRow(
		ctx,
		createWorkspaceMutation,
		pgx.NamedArgs{
			"orgId":       orgId,
			"id":          workspace.Id,
			"name":        workspace.Name,
			"accessType":  workspace.AccessType,
			"description": workspace.Description,
			"path":        workspace.Path,
		},
	)

	var ws model.Workspace
	rowErr := row.Scan(
		&ws.Id,
		&ws.Name,
		&ws.AccessType,
		&ws.Path,
		&ws.Description,
		&ws.CreatedAt,
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
  "name"        = @name,
  "accessType"  = @accessType,
  "description" = @description,
  "path"        = @path,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateWorkspaceById(ctx context.Context, workspaceId string, workspace *model.Workspace) error {
	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		updateWorkspaceByIdMutation,
		pgx.NamedArgs{
			"id":          workspaceId,
			"name":        workspace.Name,
			"accessType":  *workspace.AccessType,
			"description": workspace.Description,
			"path":        workspace.Path,
			"updatedAt":   time.Now(),
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

func (r *Repository) DeleteWorkspaceById(ctx context.Context, workspaceId string) error {
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
