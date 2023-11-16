package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// ListProjectPricings

type ListProjectPricingsParams struct {
	Context   context.Context
	ProjectId string
}

const listProjectPricingsQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "updatedAt"
FROM
  "app"."ProjectPricing"
WHERE
  "projectId" = @projectId;
`

func (r *Repository) ListProjectPricings(p *ListProjectPricingsParams) ([]model.ProjectPricing, error) {
	rows, err := r.AppDbConn.Query(p.Context, listProjectPricingsQuery, pgx.NamedArgs{"projectId": p.ProjectId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var instanceList []model.ProjectPricing

	for rows.Next() {
		var sheet model.ProjectPricing
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

/// FindUniqueProjectPricing

type FindUniqueProjectPricingParams struct {
	Context context.Context
	Id      int
}

const findUniqueProjectPricingQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "updatedAt"
FROM
  "app"."ProjectPricing"
WHERE
  "id" = @id;
`

func (r *Repository) FindUniqueProjectPricing(p *FindUniqueProjectPricingParams) (*model.ProjectPricing, error) {
	row := r.AppDbConn.QueryRow(p.Context, findUniqueProjectPricingQuery, pgx.NamedArgs{"id": p.Id})

	var sheet model.ProjectPricing
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
