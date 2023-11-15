package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// ListProjectForms

type ListProjectFormsParams struct {
	Context context.Context
	AppId   string
}

const listProjectFormsQuery = `
SELECT
  "id", "title", "rowCount", "rows", "createdAt", "updatedAt"
FROM
  "app"."ProjectForm"
WHERE
  "appId" = @appId;
`

func (r *Repository) ListProjectForms(p *ListProjectFormsParams) ([]model.ProjectForm, error) {
	rows, err := r.AppDbConn.Query(p.Context, listProjectFormsQuery, pgx.NamedArgs{"appId": p.AppId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var instanceList []model.ProjectForm

	for rows.Next() {
		var sheet model.ProjectForm
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

/// FindUniqueProjectForm

type FindUniqueProjectFormParams struct {
	Context context.Context
	AppId   string
	Id      int
}

const findUniqueProjectFormQuery = `
SELECT
  "id", "title", "rowCount", "rows", "createdAt", "updatedAt"
FROM
  "app"."ProjectForm"
WHERE
  "id" = @id;
`

func (r *Repository) FindUniqueProjectForm(p *FindUniqueProjectFormParams) (*model.ProjectForm, error) {
	row := r.AppDbConn.QueryRow(p.Context, findUniqueProjectFormQuery, pgx.NamedArgs{"id": p.Id})

	var sheet model.ProjectForm
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
