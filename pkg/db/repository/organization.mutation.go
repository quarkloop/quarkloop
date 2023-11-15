package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/quarkloop/quarkloop/pkg/db/model"
)

/// CreateOrganization

type CreateOrganizationParams struct {
	Context context.Context
	Os      model.Organization
}

const createOrganizationMutation = `
INSERT INTO
  "system"."Organization" ("id", "name", "description", "path")
VALUES
  (@id, @name, @description, @path)
RETURNING 
  "id", "name", "description", "path", "createdAt";
`

func (r *Repository) CreateOrganization(p *CreateOrganizationParams) (*model.Organization, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	p.Os.Id = id
	p.Os.Path = fmt.Sprintf("/os/%s", id)

	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		createOrganizationMutation,
		pgx.NamedArgs{
			"id":          p.Os.Id,
			"name":        p.Os.Name,
			"description": p.Os.Description,
			"path":        p.Os.Path,
			// p.Os.Overview,
			// p.Os.ImageUrl,
			// p.Os.WebsiteUrl,
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

	return &p.Os, nil
}

/// UpdateOrganizationById

type UpdateOrganizationByIdParams struct {
	Context context.Context
	OsId    string
	Os      model.Organization
}

const updateOrganizationByIdMutation = `
UPDATE
  "system"."Organization"
SET
  "name"        = @name,
  "description" = @description,
  "path"        = @path,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateOrganizationById(p *UpdateOrganizationByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		updateOrganizationByIdMutation,
		pgx.NamedArgs{
			"id":          p.OsId,
			"name":        p.Os.Name,
			"description": p.Os.Description,
			"path":        p.Os.Path,
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

/// DeleteOrganizationById

type DeleteOrganizationByIdParams struct {
	Context context.Context
	OsId    string
}

const deleteOrganizationByIdMutation = `
DELETE FROM
  "system"."Organization"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteOrganizationById(p *DeleteOrganizationByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(p.Context, deleteOrganizationByIdMutation, pgx.NamedArgs{"id": p.OsId})
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
