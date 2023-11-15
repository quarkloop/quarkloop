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
	AppId              string
	ProjectJsonDataset model.ProjectJsonDataset
}

const createProjectJsonDatasetMutation = `
INSERT INTO
  "app"."ProjectJsonDataset" ("appId", "title", "rowCount", "rows", "updatedAt")
VALUES
  (@appId, @title, @rowCount, @rows, @updatedAt)
RETURNING
  "id", "appId", "title", "rowCount", "rows", "createdAt", "updatedAt";
`

func (r *Repository) CreateProjectJsonDataset(p *CreateProjectJsonDatasetParams) (*model.ProjectJsonDataset, error) {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		createProjectJsonDatasetMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"title":     p.ProjectJsonDataset.Title,
			"rowCount":  p.ProjectJsonDataset.RowCount,
			"rows":      p.ProjectJsonDataset.Rows,
			"updatedAt": p.ProjectJsonDataset.UpdatedAt,
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
	AppId              string
	Id                 int
	ProjectJsonDataset model.ProjectJsonDataset
}

const updateProjectJsonDatasetByIdMutation = `
UPDATE
  "app"."ProjectJsonDataset"
set
  "title"     = @title,
  "rowCount"  = @rowCount,
  "rows"      = @rows,
  "updatedAt" = @updatedAt
WHERE
  "id" = @id
AND
  "appId" = @appId;
`

func (r *Repository) UpdateProjectJsonDatasetById(p *UpdateProjectJsonDatasetByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		updateProjectJsonDatasetByIdMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"id":        p.Id,
			"title":     p.ProjectJsonDataset.Title,
			"rowCount":  p.ProjectJsonDataset.RowCount,
			"rows":      p.ProjectJsonDataset.Rows,
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

/// DeleteProjectJsonDatasetById

type DeleteProjectJsonDatasetByIdParams struct {
	Context context.Context
	AppId   string
	Id      int
}

const deleteProjectJsonDatasetByIdMutation = `
DELETE FROM
  "app"."ProjectJsonDataset"
WHERE
  "id" = @id
AND
  "appId" = @appId;
`

func (r *Repository) DeleteProjectJsonDatasetById(p *DeleteProjectJsonDatasetByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(p.Context, deleteProjectJsonDatasetByIdMutation, pgx.NamedArgs{
		"appId": p.AppId,
		"id":    p.Id,
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
