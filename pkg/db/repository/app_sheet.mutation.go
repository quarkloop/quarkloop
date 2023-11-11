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

/// CreateSheetInstance

type CreateSheetInstanceParams struct {
	Context       context.Context
	AppId         string
	SheetInstance model.SheetInstance
}

const createSheetInstanceMutation = `
INSERT INTO
  "app"."SheetInstance" ("appId", "title", "rowCount", "rows", "updatedAt")
VALUES
  (@appId, @title, @rowCount, @rows, @updatedAt)
RETURNING 
  "id", "appId", "title", "rowCount", "rows", "createdAt", "updatedAt";
`

func (r *Repository) CreateSheetInstance(p *CreateSheetInstanceParams) (*model.SheetInstance, error) {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		createSheetInstanceMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"title":     p.SheetInstance.Title,
			"rowCount":  p.SheetInstance.RowCount,
			"rows":      p.SheetInstance.Rows,
			"updatedAt": p.SheetInstance.UpdatedAt,
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

	return &p.SheetInstance, nil
}

/// UpdateSheetInstanceById

type UpdateSheetInstanceByIdParams struct {
	Context       context.Context
	AppId         string
	Id            int
	SheetInstance model.SheetInstance
}

const updateSheetInstanceByIdMutation = `
UPDATE
  "app"."SheetInstance"
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

func (r *Repository) UpdateSheetInstanceById(p *UpdateSheetInstanceByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		updateSheetInstanceByIdMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"id":        p.Id,
			"title":     p.SheetInstance.Title,
			"rowCount":  p.SheetInstance.RowCount,
			"rows":      p.SheetInstance.Rows,
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

/// DeleteSheetInstanceById

type DeleteSheetInstanceByIdParams struct {
	Context context.Context
	AppId   string
	Id      int
}

const deleteSheetInstanceByIdMutation = `
DELETE FROM
  "app"."SheetInstance"
WHERE
  "id" = @id
AND
  "appId" = @appId;
`

func (r *Repository) DeleteSheetInstanceById(p *DeleteSheetInstanceByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(p.Context, deleteSheetInstanceByIdMutation, pgx.NamedArgs{
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
