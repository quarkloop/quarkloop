package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_schema"
)

/// ListTableSchemas

const listTableSchemasQuery = `
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
	"project"."TableSchema"
WHERE 
	"projectId" = @projectId;
`

func (store *tableSchemaStore) ListTableSchemas(ctx context.Context, projectId int) ([]table_schema.TableSchema, error) {
	rows, err := store.Conn.Query(ctx, listTableSchemasQuery, pgx.NamedArgs{"projectId": projectId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var schemaList []table_schema.TableSchema

	for rows.Next() {
		var schema table_schema.TableSchema
		err := rows.Scan(
			&schema.Id,
			&schema.Name,
			&schema.Description,
			&schema.Metadata,
			&schema.Data,
			&schema.CreatedAt,
			&schema.CreatedBy,
			&schema.UpdatedAt,
			&schema.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		schemaList = append(schemaList, schema)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return schemaList, nil
}

/// GetTableSchemaById

const getTableSchemaByIdQuery = `
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
	"project"."TableSchema"
WHERE 
	"projectId" = @projectId
AND 
	"id" = @id;
`

func (store *tableSchemaStore) GetTableSchemaById(ctx context.Context, projectId int, schemaId string) (*table_schema.TableSchema, error) {
	row := store.Conn.QueryRow(ctx, getTableSchemaByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"id":        schemaId,
	})

	var schema table_schema.TableSchema
	err := row.Scan(
		&schema.Id,
		&schema.Name,
		&schema.Description,
		&schema.Metadata,
		&schema.Data,
		&schema.CreatedAt,
		&schema.CreatedBy,
		&schema.UpdatedAt,
		&schema.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &schema, nil
}
