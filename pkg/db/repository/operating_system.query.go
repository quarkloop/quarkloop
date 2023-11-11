package repository

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

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
  "system"."OperatingSystem";
`

func (r *Repository) ListOperatingSystems(p *ListOperatingSystemsParams) ([]model.OperatingSystem, error) {
	rows, err := r.SystemDbConn.Query(p.Context, listOperatingSystemsQuery)
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

/// FindUniqueOperatingSystem

type FindUniqueOperatingSystemParams struct {
	Context context.Context
	Id      string
}

const findUniqueOperatingSystemQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."OperatingSystem" 
WHERE 
  "id" = @id;
`

func (r *Repository) FindUniqueOperatingSystem(p *FindUniqueOperatingSystemParams) (*model.OperatingSystem, error) {
	row := r.SystemDbConn.QueryRow(p.Context, findUniqueOperatingSystemQuery, pgx.NamedArgs{"id": p.Id})

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

/// FindFirstOperatingSystem

type FindFirstOperatingSystemParams struct {
	Context         context.Context
	OperatingSystem model.OperatingSystem
}

const findFirstOperatingSystemQuery = `
SELECT 
  "id", "name", "description", "path", "createdAt", "updatedAt"
FROM 
  "system"."OperatingSystem" 
WHERE
`

func (r *Repository) FindFirstOperatingSystem(p *FindFirstOperatingSystemParams) (*model.OperatingSystem, error) {
	availableFields := []string{}
	operatingSystemFields := map[string]interface{}{
		"id":        p.OperatingSystem.Id,
		"name":      p.OperatingSystem.Name,
		"path":      p.OperatingSystem.Path,
		"createdAt": p.OperatingSystem.CreatedAt,
		"updatedAt": p.OperatingSystem.UpdatedAt,
	}
	for key, value := range operatingSystemFields {
		switch v := value.(type) {
		case int:
			if v != 0 {
				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%d'", key, v))
			}
		case float64:
			if v != 0.0 {
				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%f'", key, v))
			}
		case string:
			if v != "" {
				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%s'", key, v))
			}
		case *time.Time:
			if v != nil {
				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%s'", key, v))
			}
		}
	}
	finalQuery := findFirstOperatingSystemQuery + strings.Join(availableFields, " AND ")

	row := r.SystemDbConn.QueryRow(p.Context, finalQuery)

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
