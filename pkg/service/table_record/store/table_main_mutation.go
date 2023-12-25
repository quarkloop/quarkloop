package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

/// CreateMainRecord

const createMainRecordMutation = `
INSERT INTO "project"."TableMain" (
    "projectId",
    "branchId",
    "name",
    "type",
    "description",
    "metadata",
    "data"
)
VALUES (
    @projectId,
    @branchId,
    @name,
    @type,
    @description,
    @metadata,
    @data
)
RETURNING 
    "id",
    "name",
    "type",
    "description",
    "metadata",
    "data",
    "createdAt";
`

func (store *tableRecordStore) CreateMainRecord(ctx context.Context, projectId int, branchId int, table *table_record.MainRecord) (*table_record.MainRecord, error) {
	row := store.Conn.QueryRow(
		ctx,
		createMainRecordMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"branchId":    branchId,
			"id":          table.Id,
			"name":        table.Name,
			"type":        table.Type,
			"description": table.Description,
			"metadata":    table.Metadata,
			"data":        table.Data,
		},
	)

	var service table_record.MainRecord
	err := row.Scan(
		&service.Id,
		&service.Name,
		&service.Type,
		&service.Description,
		&service.Metadata,
		&service.Data,
		&service.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	return &service, nil
}

/// CreateBulkMainRecord

func (store *tableRecordStore) CreateBulkMainRecords(ctx context.Context, projectId int, branchId int, tableList []table_record.MainRecord) (int64, error) {
	rowsAffected, err := store.Conn.CopyFrom(
		ctx,
		pgx.Identifier{"system", "TableMain"},
		[]string{"projectId", "branchId", "name", "type", "description", "data"},
		pgx.CopyFromSlice(len(tableList), func(i int) ([]interface{}, error) {
			return []interface{}{
				projectId,
				branchId,
				tableList[i].Name,
				tableList[i].Type,
				tableList[i].Description,
				//tableList[i].Metadata,
				tableList[i].Data,
			}, nil
		}),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return 0, err
	}

	if rowsAffected == 0 {
		notFoundErr := errors.New("cannot bulk create")
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", notFoundErr)
		return 0, notFoundErr
	}

	return rowsAffected, nil
}

/// UpdateMainRecordById

const updateMainRecordByIdMutation = `
UPDATE
    "project"."TableMain"
SET
    "name"        = @name,
    "type"        = @type,
    "description" = @description,
    "metadata"    = @metadata,
    "data"        = @data,
    "updatedAt"   = @updatedAt
WHERE
    "projectId" = @projectId
AND
    "branchId" = @branchId
AND
    "id" = @id;
`

func (store *tableRecordStore) UpdateMainRecordById(ctx context.Context, projectId int, branchId int, mainId string, table *table_record.MainRecord) error {
	commandTag, err := store.Conn.Exec(
		ctx,
		updateMainRecordByIdMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"branchId":    branchId,
			"id":          mainId,
			"name":        table.Name,
			"type":        table.Type,
			"description": table.Description,
			"metadata":    table.Metadata,
			"data":        table.Data,
			"updatedAt":   time.Now(),
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

/// DeleteMainRecordById

const deleteMainRecordByIdMutation = `
DELETE FROM
    "project"."TableMain"
WHERE
    "projectId" = @projectId
AND
    "branchId" = @branchId
AND
    "id" = @id;
`

func (store *tableRecordStore) DeleteMainRecordById(ctx context.Context, projectId int, branchId int, mainId string) error {
	commandTag, err := store.Conn.Exec(ctx, deleteMainRecordByIdMutation, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        mainId,
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
