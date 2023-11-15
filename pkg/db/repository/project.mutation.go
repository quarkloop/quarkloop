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

/// CreateProject

type CreateProjectParams struct {
	Context     context.Context
	OsId        string
	WorkspaceId string
	Project     model.Project
}

const createProjectMutation = `
INSERT INTO
  "system"."Project" ("osId", "workspaceId", "id", "name", "accessType", "path", "description", "updatedAt")
VALUES
  (@osId, @workspaceId, @id, @name, @accessType, @path, @description, @updatedAt)
RETURNING 
  "id", "name", "accessType", "path", "description", "createdAt", "updatedAt";
`

func (r *Repository) CreateProject(p *CreateProjectParams) (*model.Project, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	p.Project.Id = id
	p.Project.Path = fmt.Sprintf("/os/%s/%s/%s", p.OsId, p.WorkspaceId, id)

	fmt.Printf("\n%v\n", p)

	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		createProjectMutation,
		pgx.NamedArgs{
			"osId":        p.OsId,
			"workspaceId": p.WorkspaceId,
			"id":          p.Project.Id,
			"name":        p.Project.Name,
			"accessType":  p.Project.AccessType,
			"path":        p.Project.Path,
			"description": p.Project.Description,
			"updatedAt":   time.Now(),
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

	return &p.Project, nil
}

/// UpdateProjectById

type UpdateProjectByIdParams struct {
	Context   context.Context
	ProjectId string
	Project   model.Project
}

const updateProjectByIdMutation = `
UPDATE
  "system"."Project"
SET
  "name"        = @name,
  "path"        = @path,
  "description" = @description,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateProjectById(p *UpdateProjectByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		updateProjectByIdMutation,
		pgx.NamedArgs{
			"id":          p.ProjectId,
			"name":        p.Project.Name,
			"path":        p.Project.Path,
			"description": p.Project.Description,
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

/// DeleteProjectById

type DeleteProjectByIdParams struct {
	Context   context.Context
	ProjectId string
}

const deleteProjectByIdMutation = `
DELETE FROM
  "system"."Project"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteProjectById(p *DeleteProjectByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(p.Context, deleteProjectByIdMutation, pgx.NamedArgs{"id": p.ProjectId})
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
