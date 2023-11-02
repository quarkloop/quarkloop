package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// CreateWorkspace

type CreateWorkspaceParams struct {
	Context   context.Context
	OsId      string
	Workspace model.Workspace
}

const createWorkspaceMutation = `
INSERT INTO
  "app"."Workspace" ("id", "osId", "name", "description", "path")
VALUES
  (@id, @osId, @name, @description, @path)
RETURNING 
  "id", "osId", "name", "description", "path", "createdAt";
`

func (r *Repository) CreateWorkspace(p *CreateWorkspaceParams) (*model.Workspace, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	p.Workspace.Id = id
	p.Workspace.Path = fmt.Sprintf("/os/%s/%s", p.OsId, p.Workspace.Id)

	commandTag, err := r.Conn.Exec(
		p.Context,
		createWorkspaceMutation,
		pgx.NamedArgs{
			"id":          p.Workspace.Id,
			"osId":        p.OsId,
			"name":        p.Workspace.Name,
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
	OsId        string
	WorkspaceId string
	Workspace   model.Workspace
}

const updateWorkspaceByIdMutation = `
UPDATE
  "app"."Workspace"
set
  name = @name,
  description = @description,
  path = @path
  updatedAt = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateWorkspaceById(p *UpdateWorkspaceByIdParams) error {
	commandTag, err := r.Conn.Exec(
		p.Context,
		updateWorkspaceByIdMutation,
		pgx.NamedArgs{
			"id":          p.WorkspaceId,
			"name":        p.Workspace.Name,
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
	Context context.Context
	Id      string
}

const deleteWorkspaceByIdMutation = `
DELETE FROM
  "app"."Workspace"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteWorkspaceById(p *DeleteWorkspaceByIdParams) error {
	commandTag, err := r.Conn.Exec(p.Context, deleteWorkspaceByIdMutation, pgx.NamedArgs{"id": p.Id})
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
