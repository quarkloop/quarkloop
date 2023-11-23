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

/// CreateProject

const createProjectMutation = `
INSERT INTO
  "system"."Project" ("orgId", "workspaceId", "id", "name", "accessType", "path", "description", "updatedAt")
VALUES
  (@orgId, @workspaceId, @id, @name, @accessType, @path, @description, @updatedAt)
RETURNING 
  "id", "name", "accessType", "path", "description", "createdAt";
`

func (r *Repository) CreateProject(ctx context.Context, orgId string, workspaceId string, project *model.Project) (*model.Project, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	project.Id = id
	project.Path = fmt.Sprintf("/org/%s/%s/%s", orgId, workspaceId, id)

	row := r.SystemDbConn.QueryRow(
		ctx,
		createProjectMutation,
		pgx.NamedArgs{
			"orgId":       orgId,
			"workspaceId": workspaceId,
			"id":          project.Id,
			"name":        project.Name,
			"accessType":  project.AccessType,
			"path":        project.Path,
			"description": project.Description,
			"updatedAt":   time.Now(),
		},
	)

	var prj model.Project
	rowErr := row.Scan(
		&prj.Id,
		&prj.Name,
		&prj.AccessType,
		&prj.Path,
		&prj.Description,
		&prj.CreatedAt,
	)
	if rowErr != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", rowErr)
		return nil, rowErr
	}

	return &prj, nil
}

/// UpdateProjectById

const updateProjectByIdMutation = `
UPDATE
  "system"."Project"
SET
  "name"        = @name,
  "path"        = @path,
  "description" = @description,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateProjectById(ctx context.Context, projectId string, project *model.Project) error {
	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		updateProjectByIdMutation,
		pgx.NamedArgs{
			"id":          projectId,
			"name":        project.Name,
			"path":        project.Path,
			"description": project.Description,
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

/// DeleteProjectById

const deleteProjectByIdMutation = `
DELETE FROM
  "system"."Project"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteProjectById(ctx context.Context, projectId string) error {
	commandTag, err := r.SystemDbConn.Exec(ctx, deleteProjectByIdMutation, pgx.NamedArgs{"id": projectId})
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
