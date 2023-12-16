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

/// CreateDocumentRecord

const createDocumentRecordMutation = `
INSERT INTO
  "project"."TableDocument" ("projectId", "branchId", "name", "description", "metadata", "data")
VALUES
  (@projectId, @branchId, @name, @description, @metadata, @data)
RETURNING
  "id", "name", "description", "metadata", "data", "createdAt";
`

func (r *Repository) CreateDocumentRecord(ctx context.Context, projectId int, branchId int, doc *model.DocumentRecord) (*model.DocumentRecord, error) {
	commandTag, err := r.ProjectDbConn.Exec(
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

func (r *Repository) UpdateDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string, doc *model.DocumentRecord) error {
	commandTag, err := r.ProjectDbConn.Exec(
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

func (r *Repository) DeleteDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string) error {
	commandTag, err := r.ProjectDbConn.Exec(ctx, deleteDocumentRecordByIdMutation, pgx.NamedArgs{
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
