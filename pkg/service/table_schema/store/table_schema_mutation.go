package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_schema"
)

/// CreateTableSchema

const createTableSchemaMutation = `
INSERT INTO "project"."TableSchema" (
    "projectId",
    "name",
    "description",
    "metadata",
    "data"
)
VALUES (
    @projectId,
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

func (store *tableSchemaStore) CreateTableSchema(ctx context.Context, projectId int, schema *table_schema.TableSchema) (*table_schema.TableSchema, error) {
	commandTag, err := store.Conn.Exec(
		ctx,
		createTableSchemaMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"name":        schema.Name,
			"description": schema.Description,
			"metadata":    schema.Metadata,
			"data":        schema.Data,
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

	return schema, nil
}

/// UpdateTableSchemaById

const updateTableSchemaByIdMutation = `
UPDATE
    "project"."TableSchema"
SET
    "name"        = @name,
    "description" = @description,
    "metadata"    = @metadata,
    "data"        = @data,
    "updatedAt"   = @updatedAt
WHERE
    "id" = @id
AND
    "projectId" = @projectId;
`

func (store *tableSchemaStore) UpdateTableSchemaById(ctx context.Context, projectId int, schemaId string, schema *table_schema.TableSchema) error {
	commandTag, err := store.Conn.Exec(
		ctx,
		updateTableSchemaByIdMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"id":          schemaId,
			"name":        schema.Name,
			"description": schema.Description,
			"metadata":    schema.Metadata,
			"data":        schema.Data,
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

/// DeleteTableSchemaById

const deleteTableSchemaByIdMutation = `
DELETE FROM
    "project"."TableSchema"
WHERE
    "id" = @id
AND
    "projectId" = @projectId;
`

func (store *tableSchemaStore) DeleteTableSchemaById(ctx context.Context, projectId int, schemaId string) error {
	commandTag, err := store.Conn.Exec(ctx, deleteTableSchemaByIdMutation, pgx.NamedArgs{
		"projectId": projectId,
		"id":        schemaId,
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
