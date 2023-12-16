package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/quarkloop/quarkloop/pkg/service/workspace"
)

/// CreateWorkspace

const createWorkspaceMutation = `
INSERT INTO
  "system"."Workspace" ("orgId", "sid", "name", "description", "accessType", "createdBy")
VALUES
  (@orgId, @sid, @name, @description, @accessType, @createdBy)
RETURNING 
  "id", "sid", "orgId",
  "name", "description", "accessType",
  "createdAt", "createdBy", "updatedAt", "updatedBy";
`

func (store *workspaceStore) CreateWorkspace(ctx context.Context, orgId int, ws *workspace.Workspace) (*workspace.Workspace, error) {
	if ws.ScopedId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		ws.ScopedId = sid
	}

	row := store.Conn.QueryRow(
		ctx,
		createWorkspaceMutation,
		pgx.NamedArgs{
			"orgId":       orgId,
			"sid":         ws.ScopedId,
			"name":        ws.Name,
			"description": ws.Description,
			"accessType":  ws.AccessType,
			"createdBy":   ws.CreatedBy,
		},
	)

	var workspace workspace.Workspace
	rowErr := row.Scan(
		&workspace.Id,
		&workspace.ScopedId,
		&workspace.OrgId,
		&workspace.Name,
		&workspace.Description,
		&workspace.AccessType,
		&workspace.CreatedAt,
		&workspace.CreatedBy,
		&workspace.UpdatedAt,
		&workspace.UpdatedBy,
	)
	if rowErr != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", rowErr)
		return nil, rowErr
	}

	return &workspace, nil
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

func (store *workspaceStore) UpdateWorkspaceById(ctx context.Context, workspaceId int, workspace *workspace.Workspace) error {
	commandTag, err := store.Conn.Exec(
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

func (store *workspaceStore) DeleteWorkspaceById(ctx context.Context, workspaceId int) error {
	commandTag, err := store.Conn.Exec(ctx, deleteWorkspaceByIdMutation, pgx.NamedArgs{"id": workspaceId})
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
