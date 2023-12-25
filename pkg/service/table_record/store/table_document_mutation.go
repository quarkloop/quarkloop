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

/// CreateDocumentRecord

const createDocumentRecordMutation = `
INSERT INTO "project"."TableDocument" (
    "projectId",
    "branchId",
    "name",
    "description",
    "metadata",
    "data"
)
VALUES (
    @projectId,
    @branchId,
    @name,
    @description,
    @metadata,
    @data
)
RETURNING 
    "id",
    "name",
    "description",
    "metadata",
    "data",
    "createdAt";
`

func (store *tableRecordStore) CreateDocumentRecord(ctx context.Context, projectId int, branchId int, doc *table_record.DocumentRecord) (*table_record.DocumentRecord, error) {
	commandTag, err := store.Conn.Exec(
		ctx,
		createDocumentRecordMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"branchId":    branchId,
			"name":        doc.Name,
			"description": doc.Description,
			"metadata":    doc.Metadata,
			"data":        doc.Data,
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to create")
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", notFoundErr)
		return nil, notFoundErr
	}

	return doc, nil
}

/// UpdateDocumentRecordById

const updateDocumentRecordByIdMutation = `
UPDATE
    "project"."TableDocument"
SET
    "name"        = @name,
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

func (store *tableRecordStore) UpdateDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string, doc *table_record.DocumentRecord) error {
	commandTag, err := store.Conn.Exec(
		ctx,
		updateDocumentRecordByIdMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"branchId":    branchId,
			"id":          documentId,
			"name":        doc.Name,
			"description": doc.Description,
			"metadata":    doc.Metadata,
			"data":        doc.Data,
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

/// DeleteDocumentRecordById

const deleteDocumentRecordByIdMutation = `
DELETE FROM
    "project"."TableDocument"
WHERE
    "projectId" = @projectId
AND
    "branchId" = @branchId
AND
    "id" = @id;
`

func (store *tableRecordStore) DeleteDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string) error {
	commandTag, err := store.Conn.Exec(ctx, deleteDocumentRecordByIdMutation, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        documentId,
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
