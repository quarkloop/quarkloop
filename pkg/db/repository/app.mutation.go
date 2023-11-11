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

/// CreateApp

type CreateAppParams struct {
	Context     context.Context
	OsId        string
	WorkspaceId string
	App         model.App
}

const createAppMutation = `
INSERT INTO
  "system"."App" ("osId", "workspaceId", "id", "name", "accessType", "path", "description", "updatedAt")
VALUES
  (@osId, @workspaceId, @id, @name, @accessType, @path, @description, @updatedAt)
RETURNING 
  "id", "name", "accessType", "path", "description", "createdAt", "updatedAt";
`

func (r *Repository) CreateApp(p *CreateAppParams) (*model.App, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	p.App.Id = id
	p.App.Path = fmt.Sprintf("/os/%s/%s/%s", p.OsId, p.WorkspaceId, id)

	fmt.Printf("\n%v\n", p)

	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		createAppMutation,
		pgx.NamedArgs{
			"osId":        p.OsId,
			"workspaceId": p.WorkspaceId,
			"id":          p.App.Id,
			"name":        p.App.Name,
			"accessType":  p.App.AccessType,
			"path":        p.App.Path,
			"description": p.App.Description,
			"updatedAt":   time.Now(),
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

	return &p.App, nil
}

/// UpdateAppById

type UpdateAppByIdParams struct {
	Context context.Context
	AppId   string
	App     model.App
}

const updateAppByIdMutation = `
UPDATE
  "system"."App"
SET
  "name"        = @name,
  "path"        = @path,
  "description" = @description,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateAppById(p *UpdateAppByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		updateAppByIdMutation,
		pgx.NamedArgs{
			"id":          p.AppId,
			"name":        p.App.Name,
			"path":        p.App.Path,
			"description": p.App.Description,
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

/// DeleteAppById

type DeleteAppByIdParams struct {
	Context context.Context
	AppId   string
}

const deleteAppByIdMutation = `
DELETE FROM
  "system"."App"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteAppById(p *DeleteAppByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(p.Context, deleteAppByIdMutation, pgx.NamedArgs{"id": p.AppId})
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
