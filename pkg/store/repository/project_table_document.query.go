package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListDocumentRecords

const listDocumentRecordsQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TableDocument"
WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId;
`

func (r *Repository) ListDocumentRecords(ctx context.Context, projectId int, branchId int) ([]model.DocumentRecord, error) {
	rows, err := r.ProjectDbConn.Query(ctx, listDocumentRecordsQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var instanceList []model.DocumentRecord

	for rows.Next() {
		var document model.DocumentRecord
		err := rows.Scan(
			&document.Id,
			&document.Name,
			&document.Description,
			&document.Metadata,
			&document.Data,
			&document.CreatedAt,
			&document.CreatedBy,
			&document.UpdatedAt,
			&document.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		instanceList = append(instanceList, document)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return instanceList, nil
}

/// GetDocumentRecordById

const getDocumentRecordByIdQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TableDocument"
WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId
AND
  "id" = @id;
`

func (r *Repository) GetDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string) (*model.DocumentRecord, error) {
	row := r.ProjectDbConn.QueryRow(ctx, getDocumentRecordByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        documentId,
	})

	var document model.DocumentRecord
	err := row.Scan(
		&document.Id,
		&document.Name,
		&document.Description,
		&document.Metadata,
		&document.Data,
		&document.CreatedAt,
		&document.CreatedBy,
		&document.UpdatedAt,
		&document.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &document, nil
}
