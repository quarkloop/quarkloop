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

/// CreateProjectForm

type CreateProjectFormParams struct {
	Context     context.Context
	AppId       string
	ProjectForm model.ProjectForm
}

const createProjectFormMutation = `
INSERT INTO
  "app"."ProjectForm" ("appId", "title", "rowCount", "rows", "updatedAt")
VALUES
  (@appId, @title, @rowCount, @rows, @updatedAt)
RETURNING
  "id", "appId", "title", "rowCount", "rows", "createdAt", "updatedAt";
`

func (r *Repository) CreateProjectForm(p *CreateProjectFormParams) (*model.ProjectForm, error) {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		createProjectFormMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"title":     p.ProjectForm.Title,
			"rowCount":  p.ProjectForm.RowCount,
			"rows":      p.ProjectForm.Rows,
			"updatedAt": p.ProjectForm.UpdatedAt,
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

	return &p.ProjectForm, nil
}

/// UpdateProjectFormById

type UpdateProjectFormByIdParams struct {
	Context     context.Context
	AppId       string
	Id          int
	ProjectForm model.ProjectForm
}

const updateProjectFormByIdMutation = `
UPDATE
  "app"."ProjectForm"
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

func (r *Repository) UpdateProjectFormById(p *UpdateProjectFormByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		updateProjectFormByIdMutation,
		pgx.NamedArgs{
			"appId":     p.AppId,
			"id":        p.Id,
			"title":     p.ProjectForm.Title,
			"rowCount":  p.ProjectForm.RowCount,
			"rows":      p.ProjectForm.Rows,
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

/// DeleteProjectFormById

type DeleteProjectFormByIdParams struct {
	Context context.Context
	AppId   string
	Id      int
}

const deleteProjectFormByIdMutation = `
DELETE FROM
  "app"."ProjectForm"
WHERE
  "id" = @id
AND
  "appId" = @appId;
`

func (r *Repository) DeleteProjectFormById(p *DeleteProjectFormByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(p.Context, deleteProjectFormByIdMutation, pgx.NamedArgs{
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
