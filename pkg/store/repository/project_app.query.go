package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListProjectApps

const listProjectAppsQuery = `
SELECT 
  "id", "title", "status", "labels", "dueDate", "metadata", "data", "createdAt", "updatedAt"
FROM 
  "system"."ProjectApp"
WHERE
  "projectId" = @projectId;
`

func (r *Repository) ListProjectApps(ctx context.Context, projectId int) ([]model.App, error) {
	rows, err := r.ProjectDbConn.Query(ctx, listProjectAppsQuery, pgx.NamedArgs{
		"projectId": projectId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var appList []model.App = []model.App{}

	for rows.Next() {
		var app model.App

		err := rows.Scan(
			&app.Id,
			&app.Title,
			&app.Status,
			&app.Metadata,
			&app.Data,
			&app.CreatedAt,
			&app.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		appList = append(appList, app)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return appList, nil
}

/// GetProjectAppById

const getProjectAppByIdQuery = `
SELECT
  "id", "title", "status", "labels", "dueDate", "metadata", "data", "createdAt", "updatedAt"
FROM
  "system"."ProjectApp"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) GetProjectAppById(ctx context.Context, projectId int, projectAppId string) (*model.App, error) {
	row := r.ProjectDbConn.QueryRow(ctx, getProjectAppByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"id":        projectAppId,
	})

	var app model.App
	err := row.Scan(
		&app.Id,
		&app.Title,
		&app.Status,
		&app.Metadata,
		&app.Data,
		&app.CreatedAt,
		&app.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &app, nil
}
