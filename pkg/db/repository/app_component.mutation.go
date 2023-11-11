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

/// CreateAppComponent

type CreateAppComponentParams struct {
	Context      context.Context
	AppId        string
	AppComponent model.AppComponent
}

const createAppComponentMutation = `
INSERT INTO
  "system"."AppComponent" ("appId", "id", "name", "settings", "updatedAt")
VALUES
  (@appId, @id, @name, @settings, @updatedAt)
RETURNING 
  "id", "name", "settings", "createdAt", "updatedAt";
`

func (r *Repository) CreateAppComponent(p *CreateAppComponentParams) (*model.AppComponent, error) {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		createAppComponentMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"id":        p.AppComponent.Id,
			"name":      p.AppComponent.Name,
			"settings":  p.AppComponent.Settings,
			"createdAt": p.AppComponent.CreatedAt,
			"updatedAt": time.Now(),
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

	return &p.AppComponent, nil
}

/// UpdateAppComponentById

type UpdateAppComponentByIdParams struct {
	Context      context.Context
	AppId        string
	AppComponent model.AppComponent
}

const updateAppComponentByIdMutation = `
UPDATE
  "system"."AppComponent"
SET
  "name"        = @name,
  "path"        = @path,
  "description" = @description,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateAppComponentById(p *UpdateAppComponentByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(
		p.Context,
		updateAppComponentByIdMutation,
		pgx.NamedArgs{
			"id":        p.AppComponent.Id,
			"name":      p.AppComponent.Name,
			"settings":  p.AppComponent.Settings,
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

/// DeleteAppComponentById

type DeleteAppComponentByIdParams struct {
	Context        context.Context
	AppComponentId string
}

const deleteAppComponentByIdMutation = `
DELETE FROM
  "system"."AppComponent"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteAppComponentById(p *DeleteAppComponentByIdParams) error {
	commandTag, err := r.SystemDbConn.Exec(p.Context, deleteAppComponentByIdMutation, pgx.NamedArgs{"id": p.AppComponentId})
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
