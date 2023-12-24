package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

/// ListDocumentRecords

const listDocumentRecordsQuery = `
SELECT 
	"id",
    "name",
    "description",
    "metadata",
    "data",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
	"project"."TableDocument"
WHERE 
	"projectId" = @projectId
AND 
	"branchId" = @branchId;
`

func (store *tableRecordStore) ListDocumentRecords(ctx context.Context, projectId int, branchId int) ([]table_record.DocumentRecord, error) {
	rows, err := store.Conn.Query(ctx, listDocumentRecordsQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var documentList []table_record.DocumentRecord

	for rows.Next() {
		var document table_record.DocumentRecord
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

		documentList = append(documentList, document)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return documentList, nil
}

/// GetDocumentRecordById

const getDocumentRecordByIdQuery = `
SELECT 
	"id",
    "name",
    "description",
    "metadata",
    "data",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
	"project"."TableDocument"
WHERE 
	"projectId" = @projectId
AND 
	"branchId" = @branchId
AND 
	"id" = @id;
`

func (store *tableRecordStore) GetDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string) (*table_record.DocumentRecord, error) {
	row := store.Conn.QueryRow(ctx, getDocumentRecordByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        documentId,
	})

	var document table_record.DocumentRecord
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
