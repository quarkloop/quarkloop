package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListFormRecords

const listFormRecordsQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TableForm"
WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId;
`

func (r *Repository) ListFormRecords(ctx context.Context, projectId int, branchId int) ([]model.FormRecord, error) {
	rows, err := r.ProjectDbConn.Query(ctx, listFormRecordsQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var instanceList []model.FormRecord

	for rows.Next() {
		var form model.FormRecord
		err := rows.Scan(
			&form.Id,
			&form.Name,
			&form.Description,
			&form.Metadata,
			&form.Data,
			&form.CreatedAt,
			&form.CreatedBy,
			&form.UpdatedAt,
			&form.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		instanceList = append(instanceList, form)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return instanceList, nil
}

/// GetFormRecordById

const getFormRecordByIdQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TableForm"
WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId
AND
  "id" = @id;
`

func (r *Repository) GetFormRecordById(ctx context.Context, projectId int, branchId int, formId string) (*model.FormRecord, error) {
	row := r.ProjectDbConn.QueryRow(ctx, getFormRecordByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        formId,
	})

	var form model.FormRecord
	err := row.Scan(
		&form.Id,
		&form.Name,
		&form.Description,
		&form.Metadata,
		&form.Data,
		&form.CreatedAt,
		&form.CreatedBy,
		&form.UpdatedAt,
		&form.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &form, nil
}
