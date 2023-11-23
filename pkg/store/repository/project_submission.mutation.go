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

/// CreateProjectSubmission

const createProjectSubmissionMutation = `
INSERT INTO
  "system"."ProjectSubmission" ("userId", "projectId", "title", "metadata", "data")
VALUES
  (@userId, @projectId, @title, @metadata, @data)
RETURNING 
  "id", "title", "metadata", "data", "createdAt";
`

func (r *Repository) CreateProjectSubmission(ctx context.Context, userId string, projectId string, pSubmission *model.ProjectSubmission) (*model.ProjectSubmission, error) {
	row := r.SystemDbConn.QueryRow(
		ctx,
		createProjectSubmissionMutation,
		pgx.NamedArgs{
			"userId":    userId,
			"projectId": projectId,
			"title":     pSubmission.Title,
			"metadata":  pSubmission.Metadata,
			"data":      pSubmission.Data,
		},
	)

	var submission model.ProjectSubmission
	err := row.Scan(
		&submission.Id,
		&submission.Title,
		&submission.Metadata,
		&submission.Data,
		&submission.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	return &submission, nil
}

/// UpdateProjectSubmissionById

const updateProjectSubmissionByIdMutation = `
UPDATE
  "system"."ProjectSubmission"
SET
  "title"     = COALESCE(@title, "title"),
  "metadata"  = COALESCE(@metadata, "metadata"),
  "data"      = COALESCE(@data, "data"),
  "updatedAt" = COALESCE(@updatedAt, "updatedAt")
WHERE
  "id" = @id;
`

func (r *Repository) UpdateProjectSubmissionById(ctx context.Context, projectSubmissionId string, pSubmission *model.ProjectSubmission) error {
	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		updateProjectSubmissionByIdMutation,
		pgx.NamedArgs{
			"id":        projectSubmissionId,
			"title":     pSubmission.Title,
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

/// DeleteProjectSubmissionById

const deleteProjectSubmissionByIdMutation = `
DELETE FROM
  "system"."ProjectSubmission"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteProjectSubmissionById(ctx context.Context, projectSubmissionId string) error {
	commandTag, err := r.SystemDbConn.Exec(ctx, deleteProjectSubmissionByIdMutation, pgx.NamedArgs{
		"id": projectSubmissionId,
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
