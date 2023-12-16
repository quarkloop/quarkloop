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
  "system"."Project" ("orgId", "workspaceId", "sid", "name", "description", "accessType", "createdBy")
VALUES
  (@orgId, @workspaceId, @sid, @name, @description, @accessType, @createdBy)
RETURNING 
  "id", "sid", "orgId", "workspaceId",
  "name", "description", "accessType",
  "createdAt", "createdBy", "updatedAt", "updatedBy";
`

func (r *Repository) CreateProject(ctx context.Context, orgId int, workspaceId int, project *model.Project) (*model.Project, error) {
	if project.ScopedId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		project.ScopedId = sid
	}

	row := r.SystemDbConn.QueryRow(
		ctx,
		createProjectMutation,
		pgx.NamedArgs{
			"orgId":       orgId,
			"workspaceId": workspaceId,
			"sid":         project.ScopedId,
			"name":        project.Name,
			"description": project.Description,
			"accessType":  project.AccessType,
			"createdBy":   project.CreatedBy,
		},
	)

	var prj model.Project
	rowErr := row.Scan(
		&prj.Id,
		&prj.ScopedId,
		&prj.OrgId,
		&prj.WorkspaceId,
		&prj.Name,
		&prj.Description,
		&prj.AccessType,
		&prj.CreatedAt,
		&prj.CreatedBy,
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
  "sid"         = @sid,
  "name"        = @name,
  "description" = @description,
  "updatedAt"   = @updatedAt,
  "updatedBy"   = @updatedBy,
WHERE
  "id" = @id;
`

func (r *Repository) UpdateProjectById(ctx context.Context, projectId int, project *model.Project) error {
	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		updateProjectByIdMutation,
		pgx.NamedArgs{
			"id":          projectId,
			"sid":         project.ScopedId,
			"name":        project.Name,
			"description": project.Description,
			"updatedAt":   time.Now(),
			"updatedBy":   project.UpdatedBy,
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

func (r *Repository) DeleteProjectById(ctx context.Context, projectId int) error {
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
