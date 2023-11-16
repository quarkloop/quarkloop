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

/// CreateProjectJsonDataset

type CreateProjectJsonDatasetParams struct {
	Context            context.Context
	ProjectId          string
	ProjectJsonDataset model.ProjectJsonDataset
}

const createProjectJsonDatasetMutation = `
INSERT INTO
  "app"."ProjectJsonDataset" ("projectId", "name", "description", "metadata", "data")
VALUES
  (@projectId, @name, @description, @metadata, @data)
RETURNING
  "id", "projectId", name", "description", "metadata", "data", "createdAt";
`

func (r *Repository) CreateProjectJsonDataset(p *CreateProjectJsonDatasetParams) (*model.ProjectJsonDataset, error) {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		createProjectJsonDatasetMutation,
		pgx.NamedArgs{
			"projectId":   p.ProjectId,
			"name":        p.ProjectJsonDataset.Name,
			"description": p.ProjectJsonDataset.Description,
			"metadata":    p.ProjectJsonDataset.Metadata,
			"data":        p.ProjectJsonDataset.Data,
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

	return &p.ProjectJsonDataset, nil
}

/// UpdateProjectJsonDatasetById

type UpdateProjectJsonDatasetByIdParams struct {
	Context            context.Context
	ProjectId          string
	Id                 int
	ProjectJsonDataset model.ProjectJsonDataset
}

const updateProjectJsonDatasetByIdMutation = `
UPDATE
  "app"."ProjectJsonDataset"
set
  "name"        = @name,
  "description" = @description,
  "metadata"    = @metadata,
  "data"        = @data,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) UpdateProjectJsonDatasetById(p *UpdateProjectJsonDatasetByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		updateProjectJsonDatasetByIdMutation,
		pgx.NamedArgs{
			"projectId":   p.ProjectId,
			"id":          p.Id,
			"name":        p.ProjectJsonDataset.Name,
			"description": p.ProjectJsonDataset.Description,
			"metadata":    p.ProjectJsonDataset.Metadata,
			"data":        p.ProjectJsonDataset.Data,
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

/// DeleteProjectJsonDatasetById

type DeleteProjectJsonDatasetByIdParams struct {
	Context   context.Context
	ProjectId string
	Id        int
}

const deleteProjectJsonDatasetByIdMutation = `
DELETE FROM
  "app"."ProjectJsonDataset"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) DeleteProjectJsonDatasetById(p *DeleteProjectJsonDatasetByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(p.Context, deleteProjectJsonDatasetByIdMutation, pgx.NamedArgs{
		"projectId": p.ProjectId,
		"id":        p.Id,
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
