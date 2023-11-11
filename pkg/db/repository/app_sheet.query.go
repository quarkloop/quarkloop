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

/// ListSheetInstances

type ListSheetInstancesParams struct {
	Context context.Context
	AppId   string
}

const listSheetInstancesQuery = `
SELECT 
  "id", "title", "rowCount", "rows", "createdAt", "updatedAt"
FROM 
  "app"."SheetInstance"
WHERE
  "appId" = @appId;
`

func (r *Repository) ListSheetInstances(p *ListSheetInstancesParams) ([]model.SheetInstance, error) {
	rows, err := r.AppDbConn.Query(p.Context, listSheetInstancesQuery, pgx.NamedArgs{"appId": p.AppId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var instanceList []model.SheetInstance

	for rows.Next() {
		var sheet model.SheetInstance
		err := rows.Scan(
			&sheet.Id,
			&sheet.Title,
			&sheet.RowCount,
			&sheet.Rows,
			&sheet.CreatedAt,
			&sheet.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		instanceList = append(instanceList, sheet)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]:Rows failed: %v\n", err)
		return nil, err
	}

	return instanceList, nil
}

/// FindUniqueSheetInstance

type FindUniqueSheetInstanceParams struct {
	Context context.Context
	AppId   string
	Id      int
}

const findUniqueSheetInstanceQuery = `
SELECT
  "id", "title", "rowCount", "rows", "createdAt", "updatedAt"
FROM
  "app"."SheetInstance"
WHERE
  "id" = @id;
`

func (r *Repository) FindUniqueSheetInstance(p *FindUniqueSheetInstanceParams) (*model.SheetInstance, error) {
	row := r.AppDbConn.QueryRow(p.Context, findUniqueSheetInstanceQuery, pgx.NamedArgs{"id": p.Id})

	var sheet model.SheetInstance
	err := row.Scan(
		&sheet.Id,
		&sheet.Title,
		&sheet.RowCount,
		&sheet.Rows,
		&sheet.CreatedAt,
		&sheet.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &sheet, nil
}

/// FindFirstSheetInstance

type FindFirstSheetInstanceParams struct {
	Context       context.Context
	AppId         string
	SheetInstance model.SheetInstance
}

const findFirstSheetInstanceQuery = `
SELECT
  "id", "title", "rowCount", "rows", "createdAt", "updatedAt"
FROM
  "app"."SheetInstance"
WHERE 
  %s
ORDER BY "updatedAt" ASC 
LIMIT 1;
`

func (r *Repository) FindFirstSheetInstance(p *FindFirstSheetInstanceParams) (*model.SheetInstance, error) {
	availableFields := []string{}
	sheetFields := map[string]interface{}{
		"appId":     p.AppId,
		"id":        p.SheetInstance.Id,
		"createdAt": p.SheetInstance.CreatedAt,
		"updatedAt": p.SheetInstance.UpdatedAt,
	}
	for key, value := range sheetFields {
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
	finalQuery := fmt.Sprintf(findFirstSheetInstanceQuery, strings.Join(availableFields, " AND "))

	row := r.AppDbConn.QueryRow(p.Context, finalQuery)

	var sheet model.SheetInstance
	err := row.Scan(
		&sheet.Id,
		&sheet.Title,
		&sheet.RowCount,
		&sheet.Rows,
		&sheet.CreatedAt,
		&sheet.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &sheet, nil
}
