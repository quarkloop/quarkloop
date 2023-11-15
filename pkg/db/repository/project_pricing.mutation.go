package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// CreateProjectPricing

type CreateProjectPricingParams struct {
	Context        context.Context
	AppId          string
	ProjectPricing model.ProjectPricing
}

const createProjectPricingMutation = `
INSERT INTO
  "app"."ProjectPricing" ("appId", "title", "rowCount", "rows", "updatedAt")
VALUES
  (@appId, @title, @rowCount, @rows, @updatedAt)
RETURNING
  "id", "appId", "title", "rowCount", "rows", "createdAt", "updatedAt";
`

func (r *Repository) CreateProjectPricing(p *CreateProjectPricingParams) (*model.ProjectPricing, error) {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		createProjectPricingMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"title":     p.ProjectPricing.Title,
			"rowCount":  p.ProjectPricing.RowCount,
			"rows":      p.ProjectPricing.Rows,
			"updatedAt": p.ProjectPricing.UpdatedAt,
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to create")
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", notFoundErr)
		return nil, notFoundErr
	}

	return &p.ProjectPricing, nil
}

/// UpdateProjectPricingById

type UpdateProjectPricingByIdParams struct {
	Context        context.Context
	AppId          string
	Id             int
	ProjectPricing model.ProjectPricing
}

const updateProjectPricingByIdMutation = `
UPDATE
  "app"."ProjectPricing"
set
  "title"     = @title,
  "rowCount"  = @rowCount,
  "rows"      = @rows,
  "updatedAt" = @updatedAt
WHERE
  "id" = @id
AND
  "appId" = @appId;
`

func (r *Repository) UpdateProjectPricingById(p *UpdateProjectPricingByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		updateProjectPricingByIdMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"id":        p.Id,
			"title":     p.ProjectPricing.Title,
			"rowCount":  p.ProjectPricing.RowCount,
			"rows":      p.ProjectPricing.Rows,
			"updatedAt": time.Now(),
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to update")
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}

/// DeleteProjectPricingById

type DeleteProjectPricingByIdParams struct {
	Context context.Context
	AppId   string
	Id      int
}

const deleteProjectPricingByIdMutation = `
DELETE FROM
  "app"."ProjectPricing"
WHERE
  "id" = @id
AND
  "appId" = @appId;
`

func (r *Repository) DeleteProjectPricingById(p *DeleteProjectPricingByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(p.Context, deleteProjectPricingByIdMutation, pgx.NamedArgs{
		"appId": p.AppId,
		"id":    p.Id,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to delete")
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}
