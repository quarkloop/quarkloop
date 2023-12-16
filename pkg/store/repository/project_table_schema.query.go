package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListTableSchemas

const listTableSchemasQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TableSchema"
WHERE
  "projectId" = @projectId;
`

func (r *Repository) ListTableSchemas(ctx context.Context, projectId int) ([]model.TableSchema, error) {
	rows, err := r.ProjectDbConn.Query(ctx, listTableSchemasQuery, pgx.NamedArgs{"projectId": projectId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var instanceList []model.TableSchema

	for rows.Next() {
		var schema model.TableSchema
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

		instanceList = append(instanceList, schema)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return instanceList, nil
}

/// GetTableSchemaById

const getTableSchemaByIdQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TableSchema"
WHERE
  "projectId" = @projectId
AND
  "id" = @id;
`

func (r *Repository) GetTableSchemaById(ctx context.Context, projectId int, schemaId string) (*model.TableSchema, error) {
	row := r.ProjectDbConn.QueryRow(ctx, getTableSchemaByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"id":        schemaId,
	})

	var schema model.TableSchema
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
