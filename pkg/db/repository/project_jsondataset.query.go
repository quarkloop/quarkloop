package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// ListProjectJsonDatasets

type ListProjectJsonDatasetsParams struct {
	Context   context.Context
	ProjectId string
}

const listProjectJsonDatasetsQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "updatedAt"
FROM
  "app"."ProjectJsonDataset"
WHERE
  "projectId" = @projectId;
`

func (r *Repository) ListProjectJsonDatasets(p *ListProjectJsonDatasetsParams) ([]model.ProjectJsonDataset, error) {
	rows, err := r.AppDbConn.Query(p.Context, listProjectJsonDatasetsQuery, pgx.NamedArgs{"projectId": p.ProjectId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var instanceList []model.ProjectJsonDataset

	for rows.Next() {
		var sheet model.ProjectJsonDataset
		err := rows.Scan(
			&sheet.Id,
			&sheet.Name,
			&sheet.Description,
			&sheet.Metadata,
			&sheet.Data,
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

/// FindUniqueProjectJsonDataset

type FindUniqueProjectJsonDatasetParams struct {
	Context context.Context
	Id      int
}

const findUniqueProjectJsonDatasetQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "updatedAt"
FROM
  "app"."ProjectJsonDataset"
WHERE
  "id" = @id;
`

func (r *Repository) FindUniqueProjectJsonDataset(p *FindUniqueProjectJsonDatasetParams) (*model.ProjectJsonDataset, error) {
	row := r.AppDbConn.QueryRow(p.Context, findUniqueProjectJsonDatasetQuery, pgx.NamedArgs{"id": p.Id})

	var sheet model.ProjectJsonDataset
	err := row.Scan(
		&sheet.Id,
		&sheet.Name,
		&sheet.Description,
		&sheet.Metadata,
		&sheet.Data,
		&sheet.CreatedAt,
		&sheet.UpdatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &sheet, nil
}
