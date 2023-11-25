package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListProjectSubmissions

const listProjectSubmissionsQuery = `
SELECT 
  "id", "title", "status", "labels", "dueDate", "metadata", "data", "createdAt", "updatedAt"
FROM 
  "system"."ProjectSubmission"
WHERE
  "projectId" = @projectId;
`

func (r *Repository) ListProjectSubmissions(ctx context.Context, projectId string) ([]model.ProjectSubmission, error) {
	rows, err := r.SystemDbConn.Query(ctx, listProjectSubmissionsQuery, pgx.NamedArgs{
		"projectId": projectId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var submissionList []model.ProjectSubmission = []model.ProjectSubmission{}

	for rows.Next() {
		var submission model.ProjectSubmission
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

		submissionList = append(submissionList, submission)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]:Rows failed: %v\n", err)
		return nil, err
	}

	return submissionList, nil
}

/// FindUniqueProjectSubmission

const findUniqueProjectSubmissionQuery = `
SELECT
  "id", "title", "status", "labels", "dueDate", "metadata", "data", "createdAt", "updatedAt"
FROM
  "system"."ProjectSubmission"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) FindUniqueProjectSubmission(ctx context.Context, projectId string, projectSubmissionId string) (*model.ProjectSubmission, error) {
	row := r.SystemDbConn.QueryRow(ctx, findUniqueProjectSubmissionQuery, pgx.NamedArgs{
		"projectId": projectId,
		"id":        projectSubmissionId,
	})

	var submission model.ProjectSubmission
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

	return &submission, nil
}
