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

type CreateWorkspaceParams struct {
	Context   context.Context
	OrgId     string
	Workspace model.Workspace
}

const createWorkspaceMutation = `
INSERT INTO
  "system"."Workspace" ("orgId", "id", "name", "accessType", "description", "path")
VALUES
  (@orgId, @id, @name, @accessType, @description, @path)
RETURNING 
  "orgId", "id", "name", "accessType", "description", "path", "createdAt";
`

func (r *Repository) CreateWorkspace(p *CreateWorkspaceParams) (*model.Workspace, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	p.Workspace.Id = id
	p.Workspace.Path = fmt.Sprintf("/org/%s/%s", p.OrgId, p.Workspace.Id)

	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		createWorkspaceMutation,
		pgx.NamedArgs{
			"orgId":       p.OrgId,
			"id":          p.Workspace.Id,
			"name":        p.Workspace.Name,
			"accessType":  p.Workspace.AccessType,
			"description": p.Workspace.Description,
			"path":        p.Workspace.Path,
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to create")
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", notFoundErr)
		return nil, notFoundErr
	}

	return &p.Workspace, nil
}

/// UpdateWorkspaceById

type UpdateWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId string
	Workspace   model.Workspace
}

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

func (r *Repository) UpdateWorkspaceById(p *UpdateWorkspaceByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		updateWorkspaceByIdMutation,
		pgx.NamedArgs{
			"id":          p.WorkspaceId,
			"name":        p.Workspace.Name,
			"accessType":  p.Workspace.AccessType,
			"description": p.Workspace.Description,
			"path":        p.Workspace.Path,
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

type DeleteWorkspaceByIdParams struct {
	Context     context.Context
	WorkspaceId string
}

const deleteWorkspaceByIdMutation = `
DELETE FROM
  "system"."Workspace"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteWorkspaceById(p *DeleteWorkspaceByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(p.Context, deleteWorkspaceByIdMutation, pgx.NamedArgs{"id": p.WorkspaceId})
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
