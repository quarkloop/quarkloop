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

/// CreateAppSubmission

const createAppSubmissionMutation = `
INSERT INTO
  "project"."AppSubmission" ("userId", "projectId", "title", "metadata", "data")
VALUES
  (@userId, @projectId, @title, @metadata, @data)
RETURNING 
  "id", "title", "status", "metadata", "data", "createdAt";
`

func (r *Repository) CreateAppSubmission(ctx context.Context, userId string, projectId int, pSubmission *model.AppSubmission) (*model.AppSubmission, error) {
	row := r.ProjectDbConn.QueryRow(
		ctx,
		createAppSubmissionMutation,
		pgx.NamedArgs{
			"userId":    userId,
			"projectId": projectId,
			"title":     pSubmission.Title,
			"metadata":  pSubmission.Metadata,
			"data":      pSubmission.Data,
		},
	)

	var submission model.AppSubmission
	err := row.Scan(
		&submission.Id,
		&submission.Title,
		&submission.Status,
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

/// UpdateAppSubmissionById

const updateAppSubmissionByIdMutation = `
UPDATE
  "project"."AppSubmission"
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

func (r *Repository) UpdateAppSubmissionById(ctx context.Context, appSubmissionId string, pSubmission *model.AppSubmission) error {
	commandTag, err := r.ProjectDbConn.Exec(
		ctx,
		updateAppSubmissionByIdMutation,
		pgx.NamedArgs{
			"id":        appSubmissionId,
			"title":     pSubmission.Title,
			"status":    pSubmission.Status,
			"labels":    pSubmission.LabelList,
			"dueDate":   pSubmission.DueDate,
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

/// DeleteAppSubmissionById

const deleteAppSubmissionByIdMutation = `
DELETE FROM
  "project"."AppSubmission"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteAppSubmissionById(ctx context.Context, appSubmissionId string) error {
	commandTag, err := r.ProjectDbConn.Exec(ctx, deleteAppSubmissionByIdMutation, pgx.NamedArgs{
		"id": appSubmissionId,
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
