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

/// CreateOperatingSystem

type CreateOperatingSystemParams struct {
	Context context.Context
	Os      model.OperatingSystem
}

const createOperatingSystemMutation = `
INSERT INTO
  "system"."OperatingSystem" ("id", "name", "description", "path")
VALUES
  (@id, @name, @description, @path)
RETURNING 
  "id", "name", "description", "path", "createdAt";
`

func (r *Repository) CreateOperatingSystem(p *CreateOperatingSystemParams) (*model.OperatingSystem, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	p.Os.Id = id
	p.Os.Path = fmt.Sprintf("/os/%s", id)

	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		createOperatingSystemMutation,
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

/// UpdateOperatingSystemById

type UpdateOperatingSystemByIdParams struct {
	Context context.Context
	OsId    string
	Os      model.OperatingSystem
}

const updateOperatingSystemByIdMutation = `
UPDATE
  "system"."OperatingSystem"
SET
  "name"        = @name,
  "description" = @description,
  "path"        = @path,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateOperatingSystemById(p *UpdateOperatingSystemByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		updateOperatingSystemByIdMutation,
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

/// DeleteOperatingSystemById

type DeleteOperatingSystemByIdParams struct {
	Context context.Context
	OsId    string
}

const deleteOperatingSystemByIdMutation = `
DELETE FROM
  "system"."OperatingSystem"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteOperatingSystemById(p *DeleteOperatingSystemByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(p.Context, deleteOperatingSystemByIdMutation, pgx.NamedArgs{"id": p.OsId})
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
