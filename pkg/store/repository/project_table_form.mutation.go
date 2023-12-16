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

/// CreateFormRecord

const createFormRecordMutation = `
INSERT INTO
  "project"."TableForm" ("projectId", "branchId", "name", "description", "metadata", "data")
VALUES
  (@projectId, @branchId, @name, @description, @metadata, @data)
RETURNING
  "id", "name", "description", "metadata", "data", "createdAt";
`

func (r *Repository) CreateFormRecord(ctx context.Context, projectId int, branchId int, form *model.FormRecord) (*model.FormRecord, error) {
	commandTag, err := r.ProjectDbConn.Exec(
		ctx,
		createFormRecordMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"branchId":    branchId,
			"name":        form.Name,
			"description": form.Description,
			"metadata":    form.Metadata,
			"data":        form.Data,
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

	return form, nil
}

/// UpdateFormRecordById

const updateFormRecordByIdMutation = `
UPDATE
  "project"."TableForm"
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

func (r *Repository) UpdateFormRecordById(ctx context.Context, projectId int, branchId int, formId string, form *model.FormRecord) error {
	commandTag, err := r.ProjectDbConn.Exec(
		ctx,
		updateFormRecordByIdMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"branchId":    branchId,
			"id":          formId,
			"name":        form.Name,
			"description": form.Description,
			"metadata":    form.Metadata,
			"data":        form.Data,
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

/// DeleteFormRecordById

const deleteFormRecordByIdMutation = `
DELETE FROM
  "project"."TableForm"
WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId
AND
  "id" = @id;
`

func (r *Repository) DeleteFormRecordById(ctx context.Context, projectId int, branchId int, formId string) error {
	commandTag, err := r.ProjectDbConn.Exec(ctx, deleteFormRecordByIdMutation, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        formId,
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
