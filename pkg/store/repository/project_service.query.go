package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListProjectServices

const listProjectServicesQuery = `
SELECT 
  "id", "name", "type", "description", "metadata", "data", "createdAt", "updatedAt"
FROM 
  "system"."ProjectService"
WHERE
  "projectId" = @projectId;
`

func (r *Repository) ListProjectServices(ctx context.Context, projectId string) ([]model.ProjectService, error) {
	rows, err := r.SystemDbConn.Query(ctx, listProjectServicesQuery, pgx.NamedArgs{
		"projectId": projectId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var serviceList []model.ProjectService = []model.ProjectService{}

	for rows.Next() {
		var service model.ProjectService
		err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.Type,
			&service.Description,
			&service.Metadata,
			&service.Data,
			&service.CreatedAt,
			&service.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		serviceList = append(serviceList, service)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]:Rows failed: %v\n", err)
		return nil, err
	}

	return serviceList, nil
}

/// FindUniqueProjectService

const findUniqueProjectServiceQuery = `
SELECT
  "id", "name", "type", "description", "metadata", "data", "createdAt", "updatedAt"
FROM
  "system"."ProjectService"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) FindUniqueProjectService(ctx context.Context, projectId string, projectServiceId string) (*model.ProjectService, error) {
	row := r.SystemDbConn.QueryRow(ctx, findUniqueProjectServiceQuery, pgx.NamedArgs{
		"projectId": projectId,
		"id":        projectServiceId,
	})

	var service model.ProjectService
	err := row.Scan(
		&service.Id,
		&service.Name,
		&service.Type,
		&service.Description,
		&service.Metadata,
		&service.Data,
		&service.CreatedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &service, nil
}
