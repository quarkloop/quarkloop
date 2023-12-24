package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

/// ListMainRecords

const listMainRecordsQuery = `
SELECT 
    tbl."id",
    tbl."name",
    tbl."type",
    tbl."description",
    tbl."metadata",
    tbl."data",
    tbl."createdAt",
    tbl."createdBy",
    tbl."updatedAt",
    tbl."updatedBy" COUNT(doc.id) AS doc,
    COUNT(form.id) AS form,
    COUNT(payment.id) AS payment
FROM 
    project."TableMain" AS tbl
LEFT JOIN 
    project."TableDocument" AS doc ON tbl.id = doc."mainId"
LEFT JOIN 
    project."TableForm" AS form ON tbl.id = form."mainId"
LEFT JOIN 
    project."TablePayment" AS payment ON tbl.id = payment."mainId"
WHERE 
    tbl."projectId" = @projectId
AND 
    tbl."branchId" = @branchId
GROUP BY 
    tbl.id;
`

func (store *tableRecordStore) ListMainRecords(ctx context.Context, projectId int, branchId int) ([]table_record.MainRecordWithRelationCount, error) {
	rows, err := store.Conn.Query(ctx, listMainRecordsQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var tableList []table_record.MainRecordWithRelationCount = []table_record.MainRecordWithRelationCount{}

	for rows.Next() {
		var table table_record.MainRecordWithRelationCount
		err := rows.Scan(
			&table.Table.Id,
			&table.Table.Name,
			&table.Table.Type,
			&table.Table.Description,
			&table.Table.Metadata,
			&table.Table.Data,
			&table.Table.CreatedAt,
			&table.Table.CreatedBy,
			&table.Table.UpdatedAt,
			&table.Table.UpdatedBy,
			&table.Relations.Document,
			&table.Relations.Form,
			&table.Relations.Payment,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		tableList = append(tableList, table)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return tableList, nil
}

/// GetMainRecordById

const getMainRecordByIdQuery = `
SELECT 
    tbl."id",
    tbl."name",
    tbl."type",
    tbl."description",
    tbl."metadata",
    tbl."data",
    tbl."createdAt",
    tbl."createdBy",
    tbl."updatedAt",
    tbl."updatedBy" COUNT(doc.id) AS doc,
    COUNT(form.id) AS form,
    COUNT(payment.id) AS payment
FROM 
    "project"."TableMain" AS tbl
LEFT JOIN 
    "project"."TableDocument" AS doc ON tbl.id = doc."mainId"
LEFT JOIN 
    "project"."TableForm" AS form ON tbl.id = form."mainId"
LEFT JOIN 
    "project"."TablePayment" AS payment ON tbl.id = payment."mainId"
WHERE 
    tbl."projectId" = @projectId
AND 
    tbl."branchId" = @branchId
AND 
    tbl.id = @id
GROUP BY 
    tbl.id;
`

func (store *tableRecordStore) GetMainRecordById(ctx context.Context, projectId int, branchId int, mainId string) (*table_record.MainRecordWithRelationCount, error) {
	row := store.Conn.QueryRow(ctx, getMainRecordByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        mainId,
	})

	var table table_record.MainRecordWithRelationCount
	err := row.Scan(
		&table.Table.Id,
		&table.Table.Name,
		&table.Table.Type,
		&table.Table.Description,
		&table.Table.Metadata,
		&table.Table.Data,
		&table.Table.CreatedAt,
		&table.Table.CreatedBy,
		&table.Table.UpdatedAt,
		&table.Table.UpdatedBy,
		&table.Relations.Document,
		&table.Relations.Form,
		&table.Relations.Payment,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &table, nil
}
