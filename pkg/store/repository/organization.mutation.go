package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// CreateOrganization

type CreateOrganizationParams struct {
	Context      context.Context
	Organization model.Organization
}

const createOrganizationMutation = `
INSERT INTO
  "system"."Organization" ("id", "name", "accessType", "description", "path")
VALUES
  (@id, @name, @accessType, @description, @path)
RETURNING 
  "id", "name", "accessType", "description", "path", "createdAt";
`

func (r *Repository) CreateOrganization(p *CreateOrganizationParams) (*model.Organization, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	p.Organization.Id = id
	p.Organization.Path = fmt.Sprintf("/org/%s", id)

	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		createOrganizationMutation,
		pgx.NamedArgs{
			"id":          p.Organization.Id,
			"name":        p.Organization.Name,
			"accessType":  p.Organization.AccessType,
			"description": p.Organization.Description,
			"path":        p.Organization.Path,
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

	return &p.Organization, nil
}

/// UpdateOrganizationById

type UpdateOrganizationByIdParams struct {
	Context      context.Context
	OrgId        string
	Organization model.Organization
}

const updateOrganizationByIdMutation = `
UPDATE
  "system"."Organization"
SET
  "name"        = @name,
  "accessType"  = @accessType,
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
			"id":          p.OrgId,
			"name":        p.Organization.Name,
			"accessType":  p.Organization.AccessType,
			"description": p.Organization.Description,
			"path":        p.Organization.Path,
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
	OrgId   string
}

const deleteOrganizationByIdMutation = `
DELETE FROM
  "system"."Organization"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteOrganizationById(p *DeleteOrganizationByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(p.Context, deleteOrganizationByIdMutation, pgx.NamedArgs{"id": p.OrgId})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", err)
		return err
	}

	fmt.Printf("\nDE: %v\n", *p)

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to delete")
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}
