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
	ProjectId   string
	ProjectForm model.ProjectForm
}

const createProjectFormMutation = `
INSERT INTO
  "app"."ProjectForm" ("projectId", "name", "description", "metadata", "data")
VALUES
  (@projectId, @name, @description, @metadata, @data)
RETURNING
  "id", "projectId", "name", "description", "metadata", "data", "createdAt";
`

func (r *Repository) CreateProjectForm(p *CreateProjectFormParams) (*model.ProjectForm, error) {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		createProjectFormMutation,
		pgx.NamedArgs{
			"projectId":   p.ProjectId,
			"name":        p.ProjectForm.Name,
			"description": p.ProjectForm.Description,
			"metadata":    p.ProjectForm.Metadata,
			"data":        p.ProjectForm.Data,
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
	ProjectId   string
	Id          int
	ProjectForm model.ProjectForm
}

const updateProjectFormByIdMutation = `
UPDATE
  "app"."ProjectForm"
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

func (r *Repository) UpdateProjectFormById(p *UpdateProjectFormByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(
		p.Context,
		updateProjectFormByIdMutation,
		pgx.NamedArgs{
			"projectId":   p.ProjectId,
			"id":          p.Id,
			"name":        p.ProjectForm.Name,
			"description": p.ProjectForm.Description,
			"metadata":    p.ProjectForm.Metadata,
			"data":        p.ProjectForm.Data,
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

/// DeleteProjectFormById

type DeleteProjectFormByIdParams struct {
	Context   context.Context
	ProjectId string
	Id        int
}

const deleteProjectFormByIdMutation = `
DELETE FROM
  "app"."ProjectForm"
WHERE
  "id" = @id
AND
  "projectId" = @projectId;
`

func (r *Repository) DeleteProjectFormById(p *DeleteProjectFormByIdParams) error {
	commandTag, err := r.AppDbConn.Exec(p.Context, deleteProjectFormByIdMutation, pgx.NamedArgs{
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
