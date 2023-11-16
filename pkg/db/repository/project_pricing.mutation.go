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
	ProjectId      string
	ProjectPricing model.ProjectPricing
}

const createProjectPricingMutation = `
INSERT INTO
  "app"."ProjectPricing" ("projectId", "name", "description", "metadata", "data")
VALUES
  (@projectId, @name, @description, @metadata, @data)
RETURNING
  "id", "projectId", "name", "description", "metadata", "data", "createdAt";
`

func (r *Repository) CreateProjectPricing(p *CreateProjectPricingParams) (*model.ProjectPricing, error) {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		createProjectPricingMutation,
		pgx.NamedArgs{
			"projectId":   p.ProjectId,
			"name":        p.ProjectPricing.Name,
			"description": p.ProjectPricing.Description,
			"metadata":    p.ProjectPricing.Metadata,
			"data":        p.ProjectPricing.Data,
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
	ProjectId      string
	Id             int
	ProjectPricing model.ProjectPricing
}

const updateProjectPricingByIdMutation = `
UPDATE
  "app"."ProjectPricing"
set
  "name"        = @name,
  "description" = @description,
  "metadata"    = @metadata,
  "data"        = @data,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) UpdateProjectPricingById(p *UpdateProjectPricingByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		updateProjectPricingByIdMutation,
		pgx.NamedArgs{
			"projectId":   p.ProjectId,
			"id":          p.Id,
			"name":        p.ProjectPricing.Name,
			"description": p.ProjectPricing.Description,
			"metadata":    p.ProjectPricing.Metadata,
			"data":        p.ProjectPricing.Data,
			"updatedAt":   time.Now(),
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
	Context   context.Context
	ProjectId string
	Id        int
}

const deleteProjectPricingByIdMutation = `
DELETE FROM
  "app"."ProjectPricing"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) DeleteProjectPricingById(p *DeleteProjectPricingByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(p.Context, deleteProjectPricingByIdMutation, pgx.NamedArgs{
		"projectId": p.ProjectId,
		"id":        p.Id,
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
