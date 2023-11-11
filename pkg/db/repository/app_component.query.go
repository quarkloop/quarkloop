package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// ListAppComponents

type ListAppComponentsParams struct {
	Context context.Context
	AppId   string
}

const listAppComponentsQuery = `
SELECT 
  "id", "name", "settings", "createdAt", "updatedAt"
FROM 
  "system"."AppComponent"
WHERE
  "appId" = @appId;
`

func (r *Repository) ListAppComponents(p *ListAppComponentsParams) ([]model.AppComponent, error) {
	rows, err := r.SystemDbConn.Query(p.Context, listAppComponentsQuery, pgx.NamedArgs{
		"appId": p.AppId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var compList []model.AppComponent

	for rows.Next() {
		var component model.AppComponent
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

/// FindUniqueAppComponent

type FindUniqueAppComponentParams struct {
	Context     context.Context
	AppId       string
	ComponentId string
}

const findUniqueAppComponentQuery = `
SELECT
  "id", "name", "settings", "createdAt", "updatedAt"
FROM
  "system"."AppComponent"
WHERE
  "id" = @id
AND
  "appId" = @appId;
`

func (r *Repository) FindUniqueAppComponent(p *FindUniqueAppComponentParams) (*model.AppComponent, error) {
	row := r.SystemDbConn.QueryRow(p.Context, findUniqueAppComponentQuery, pgx.NamedArgs{
		"appId": p.AppId,
		"id":    p.ComponentId,
	})

	var app model.AppComponent
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
