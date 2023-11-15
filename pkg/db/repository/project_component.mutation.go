package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// CreateProjectComponent

type CreateProjectComponentParams struct {
	Context          context.Context
	ProjectId        string
	ProjectComponent model.ProjectComponent
}

const createProjectComponentMutation = `
INSERT INTO
  "system"."ProjectComponent" ("appId", "id", "name", "settings", "updatedAt")
VALUES
  (@appId, @id, @name, @settings, @updatedAt)
RETURNING 
  "id", "name", "settings", "createdAt", "updatedAt";
`

func (r *Repository) CreateProjectComponent(p *CreateProjectComponentParams) (*model.ProjectComponent, error) {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		createProjectComponentMutation,
		pgx.NamedArgs{
			"appId":     p.ProjectId,
			"id":        p.ProjectComponent.Id,
			"name":      p.ProjectComponent.Name,
			"settings":  p.ProjectComponent.Settings,
			"createdAt": p.ProjectComponent.CreatedAt,
			"updatedAt": time.Now(),
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

	return &p.ProjectComponent, nil
}

/// UpdateProjectComponentById

type UpdateProjectComponentByIdParams struct {
	Context          context.Context
	ProjectId        string
	ProjectComponent model.ProjectComponent
}

const updateProjectComponentByIdMutation = `
UPDATE
  "system"."ProjectComponent"
SET
  "name"        = @name,
  "path"        = @path,
  "description" = @description,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateProjectComponentById(p *UpdateProjectComponentByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		updateProjectComponentByIdMutation,
		pgx.NamedArgs{
			"id":        p.ProjectComponent.Id,
			"name":      p.ProjectComponent.Name,
			"settings":  p.ProjectComponent.Settings,
			"updatedAt": time.Now(),
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

/// DeleteProjectComponentById

type DeleteProjectComponentByIdParams struct {
	Context            context.Context
	ProjectComponentId string
}

const deleteProjectComponentByIdMutation = `
DELETE FROM
  "system"."ProjectComponent"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteProjectComponentById(p *DeleteProjectComponentByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(p.Context, deleteProjectComponentByIdMutation, pgx.NamedArgs{"id": p.ProjectComponentId})
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
