package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// ListProjectComponents

type ListProjectComponentsParams struct {
	Context   context.Context
	ProjectId string
}

const listProjectComponentsQuery = `
SELECT 
  "id", "name", "settings", "createdAt", "updatedAt"
FROM 
  "system"."ProjectComponent"
WHERE
  "appId" = @appId;
`

func (r *Repository) ListProjectComponents(p *ListProjectComponentsParams) ([]model.ProjectComponent, error) {
	rows, err := r.SystemDbConn.Query(p.Context, listProjectComponentsQuery, pgx.NamedArgs{
		"appId": p.ProjectId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var compList []model.ProjectComponent

	for rows.Next() {
		var component model.ProjectComponent
		err := rows.Scan(
			&component.Id,
			&component.Name,
			&component.Settings,
			&component.CreatedAt,
			&component.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		compList = append(compList, component)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]:Rows failed: %v\n", err)
		return nil, err
	}

	return compList, nil
}

/// FindUniqueProjectComponent

type FindUniqueProjectComponentParams struct {
	Context     context.Context
	ProjectId   string
	ComponentId string
}

const findUniqueProjectComponentQuery = `
SELECT
  "id", "name", "settings", "createdAt", "updatedAt"
FROM
  "system"."ProjectComponent"
WHERE
  "id" = @id
AND
  "appId" = @appId;
`

func (r *Repository) FindUniqueProjectComponent(p *FindUniqueProjectComponentParams) (*model.ProjectComponent, error) {
	row := r.SystemDbConn.QueryRow(p.Context, findUniqueProjectComponentQuery, pgx.NamedArgs{
		"appId": p.ProjectId,
		"id":    p.ComponentId,
	})

	var app model.ProjectComponent
	err := row.Scan(
		&app.Id,
		&app.Name,
		&app.Settings,
		&app.CreatedAt,
		&app.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &app, nil
}
