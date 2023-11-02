package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// ListOperatingSystems

type ListOperatingSystemsParams struct {
	Context context.Context
}

const listOperatingSystemsQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "app"."OperatingSystem";
`

func (r *Repository) ListOperatingSystems(p *ListOperatingSystemsParams) ([]model.OperatingSystem, error) {
	rows, err := r.Conn.Query(p.Context, listOperatingSystemsQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var osList []model.OperatingSystem

	for rows.Next() {
		var operatingSystem model.OperatingSystem
		err := rows.Scan(
			&operatingSystem.Id,
			&operatingSystem.Name,
			&operatingSystem.Description,
			&operatingSystem.Path,
			&operatingSystem.CreatedAt,
			&operatingSystem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		osList = append(osList, operatingSystem)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}

	return osList, nil
}

/// GetOperatingSystemById

type GetOperatingSystemByIdParams struct {
	Context context.Context
	Id      string
}

const getOperatingSystemByIdQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "app"."OperatingSystem" 
WHERE 
  "id" = @id;
`

func (r *Repository) GetOperatingSystemById(p *GetOperatingSystemByIdParams) (*model.OperatingSystem, error) {
	row := r.Conn.QueryRow(p.Context, getOperatingSystemByIdQuery, pgx.NamedArgs{"id": p.Id})

	var operatingSystem model.OperatingSystem
	err := row.Scan(
		&operatingSystem.Id,
		&operatingSystem.Name,
		&operatingSystem.Description,
		&operatingSystem.Path,
		&operatingSystem.CreatedAt,
		&operatingSystem.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &operatingSystem, nil
}
