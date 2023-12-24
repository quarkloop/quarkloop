package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

/// ListFormRecords

const listFormRecordsQuery = `
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
    "project"."TableForm"
WHERE 
    "projectId" = @projectId
AND 
    "branchId" = @branchId;
`

func (store *tableRecordStore) ListFormRecords(ctx context.Context, projectId int, branchId int) ([]table_record.FormRecord, error) {
	rows, err := store.Conn.Query(ctx, listFormRecordsQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var formList []table_record.FormRecord

	for rows.Next() {
		var form table_record.FormRecord
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

		formList = append(formList, form)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return formList, nil
}

/// GetFormRecordById

const getFormRecordByIdQuery = `
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
    "project"."TableForm"
WHERE 
    "projectId" = @projectId
AND 
    "branchId" = @branchId
AND 
    "id" = @id;
`

func (store *tableRecordStore) GetFormRecordById(ctx context.Context, projectId int, branchId int, formId string) (*table_record.FormRecord, error) {
	row := store.Conn.QueryRow(ctx, getFormRecordByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        formId,
	})

	var form table_record.FormRecord
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
