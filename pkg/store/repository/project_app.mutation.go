package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// CreateProjectApp

const createProjectAppMutation = `
INSERT INTO
  "system"."ProjectApp" ("userId", "projectId", "title", "metadata", "data")
VALUES
  (@userId, @projectId, @title, @metadata, @data)
RETURNING 
  "id", "title", "status", "metadata", "data", "createdAt";
`

func (r *Repository) CreateProjectApp(ctx context.Context, userId string, projectId int, pSubmission *model.App) (*model.App, error) {
	row := r.ProjectDbConn.QueryRow(
		ctx,
		createProjectAppMutation,
		pgx.NamedArgs{
			"userId":    userId,
			"projectId": projectId,
			"title":     pSubmission.Title,
			"metadata":  pSubmission.Metadata,
			"data":      pSubmission.Data,
		},
	)

	var app model.App
	err := row.Scan(
		&app.Id,
		&app.Title,
		&app.Status,
		&app.Metadata,
		&app.Data,
		&app.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	return &app, nil
}

/// UpdateProjectAppById

const updateProjectAppByIdMutation = `
UPDATE
  "system"."ProjectApp"
SET
  "title"     = COALESCE(@title, "title"),
  "status"    = COALESCE(@status, "status"),
  "labels"    = COALESCE(@labels, "labels"),
  "dueDate"   = COALESCE(@dueDate, "dueDate"),
  "metadata"  = COALESCE(@metadata, "metadata"),
  "data"      = COALESCE(@data, "data"),
  "updatedAt" = COALESCE(@updatedAt, "updatedAt")
WHERE
  "id" = @id;
`

func (r *Repository) UpdateProjectAppById(ctx context.Context, projectAppId string, pSubmission *model.App) error {
	commandTag, err := r.ProjectDbConn.Exec(
		ctx,
		updateProjectAppByIdMutation,
		pgx.NamedArgs{
			"id":        projectAppId,
			"title":     pSubmission.Title,
			"status":    pSubmission.Status,
			"metadata":  pSubmission.Metadata,
			"data":      pSubmission.Data,
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

/// DeleteProjectAppById

const deleteProjectAppByIdMutation = `
DELETE FROM
  "system"."ProjectApp"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteProjectAppById(ctx context.Context, projectAppId string) error {
	commandTag, err := r.ProjectDbConn.Exec(ctx, deleteProjectAppByIdMutation, pgx.NamedArgs{
		"id": projectAppId,
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
