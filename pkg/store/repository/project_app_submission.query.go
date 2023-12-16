package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListAppSubmissions

const listAppSubmissionsQuery = `
SELECT 
  "id", "title", "status", "labels", "dueDate", "metadata", "data", "createdAt", "updatedAt"
FROM 
  "project"."AppSubmission"
WHERE
  "projectId" = @projectId;
`

func (r *Repository) ListAppSubmissions(ctx context.Context, projectId int) ([]model.AppSubmission, error) {
	rows, err := r.ProjectDbConn.Query(ctx, listAppSubmissionsQuery, pgx.NamedArgs{
		"projectId": projectId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var submissionList []model.AppSubmission = []model.AppSubmission{}

	for rows.Next() {
		var submission model.AppSubmission

		err := rows.Scan(
			&submission.Id,
			&submission.Title,
			&submission.Status,
			&submission.LabelList,
			&submission.DueDate,
			&submission.Metadata,
			&submission.Data,
			&submission.CreatedAt,
			&submission.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if submission.LabelList == nil {
			submission.LabelList = []string{}
		}
		submissionList = append(submissionList, submission)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return submissionList, nil
}

/// GetAppSubmissionById

const getAppSubmissionByIdQuery = `
SELECT
  "id", "title", "status", "labels", "dueDate", "metadata", "data", "createdAt", "updatedAt"
FROM
  "project"."AppSubmission"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) GetAppSubmissionById(ctx context.Context, projectId int, appSubmissionId string) (*model.AppSubmission, error) {
	row := r.ProjectDbConn.QueryRow(ctx, getAppSubmissionByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"id":        appSubmissionId,
	})

	var submission model.AppSubmission
	err := row.Scan(
		&submission.Id,
		&submission.Title,
		&submission.Status,
		&submission.LabelList,
		&submission.DueDate,
		&submission.Metadata,
		&submission.Data,
		&submission.CreatedAt,
		&submission.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	if submission.LabelList == nil {
		submission.LabelList = []string{}
	}

	return &submission, nil
}
