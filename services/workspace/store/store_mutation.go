package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/quarkloop/quarkloop/pkg/model"
	wsErrors "github.com/quarkloop/quarkloop/services/workspace/errors"
)

/// CreateWorkspace

const createWorkspaceMutation = `
INSERT INTO "system"."Workspace" (
    "orgId",
    "sid",
    "name",
    "description",
    "visibility",
    "createdBy"
)
VALUES (
    @orgId,
    @sid,
    @name,
    @description,
    @visibility,
    @createdBy
)
RETURNING 
    "id",
    "sid",
    "orgId",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy";
`

type CreateWorkspaceCommand struct {
	OrgId     int64
	CreatedBy string

	ScopeId     string
	Name        string
	Description string
	Visibility  model.ScopeVisibility
}

func (store *workspaceStore) CreateWorkspace(ctx context.Context, cmd *CreateWorkspaceCommand) (*model.Workspace, error) {
	if cmd.ScopeId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		cmd.ScopeId = sid
	}

	row := store.Conn.QueryRow(ctx, createWorkspaceMutation, pgx.NamedArgs{
		"orgId":       cmd.OrgId,
		"sid":         cmd.ScopeId,
		"name":        cmd.Name,
		"description": cmd.Description,
		"visibility":  cmd.Visibility,
		"createdBy":   cmd.CreatedBy,
	})

	var ws model.Workspace
	err := row.Scan(
		&ws.Id,
		&ws.ScopeId,
		&ws.OrgId,
		&ws.Name,
		&ws.Description,
		&ws.Visibility,
		&ws.CreatedAt,
		&ws.CreatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, wsErrors.HandleError(err)
	}

	return &ws, nil
}

/// UpdateWorkspaceById

const updateWorkspaceByIdMutation = `
UPDATE
    "system"."Workspace"
SET
    "sid"         = COALESCE (NULLIF(@sid, ''), "sid"),
    "name"        = COALESCE (NULLIF(@name, ''), "name"),
    "description" = COALESCE (NULLIF(@description, ''), "description"),
    "visibility"  = COALESCE (NULLIF(@visibility, 0), "visibility"),
    "updatedAt"   = @updatedAt,
    "updatedBy"   = @updatedBy
WHERE
    "id" = @id
AND
    "orgId" = @orgId;
`

type UpdateWorkspaceByIdCommand struct {
	OrgId       int64
	WorkspaceId int64
	UpdatedBy   string

	ScopeId     string
	Name        string
	Description string
	Visibility  model.ScopeVisibility
}

func (store *workspaceStore) UpdateWorkspaceById(ctx context.Context, cmd *UpdateWorkspaceByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updateWorkspaceByIdMutation, pgx.NamedArgs{
		"orgId":       cmd.OrgId,
		"id":          cmd.WorkspaceId,
		"sid":         cmd.ScopeId,
		"name":        cmd.Name,
		"description": cmd.Description,
		"visibility":  cmd.Visibility,
		"updatedBy":   cmd.UpdatedBy,
		"updatedAt":   time.Now(),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
		return wsErrors.HandleError(err)
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
    "id" = @id
AND
    "orgId" = @orgId;	
`

type DeleteWorkspaceByIdCommand struct {
	OrgId       int64
	WorkspaceId int64
}

func (store *workspaceStore) DeleteWorkspaceById(ctx context.Context, cmd *DeleteWorkspaceByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteWorkspaceByIdMutation, pgx.NamedArgs{
		"id":    cmd.WorkspaceId,
		"orgId": cmd.OrgId,
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
