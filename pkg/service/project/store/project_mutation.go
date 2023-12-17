package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/quarkloop/quarkloop/pkg/service/project"
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

func (store *projectStore) CreateProject(ctx context.Context, orgId int, workspaceId int, p *project.Project) (*project.Project, error) {
	if p.ScopedId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		p.ScopedId = sid
	}

	row := store.Conn.QueryRow(
		ctx,
		createProjectMutation,
		pgx.NamedArgs{
			"orgId":       orgId,
			"workspaceId": workspaceId,
			"sid":         p.ScopedId,
			"name":        p.Name,
			"description": p.Description,
			"accessType":  p.AccessType,
			"createdBy":   p.CreatedBy,
		},
	)

	var project project.Project
	rowErr := row.Scan(
		&project.Id,
		&project.ScopedId,
		&project.OrgId,
		&project.WorkspaceId,
		&project.Name,
		&project.Description,
		&project.AccessType,
		&project.CreatedAt,
		&project.CreatedBy,
	)
	if rowErr != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", rowErr)
		return nil, rowErr
	}

	return &project, nil
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

func (store *projectStore) UpdateProjectById(ctx context.Context, projectId int, project *project.Project) error {
	commandTag, err := store.Conn.Exec(
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

func (store *projectStore) DeleteProjectById(ctx context.Context, projectId int) error {
	commandTag, err := store.Conn.Exec(ctx, deleteProjectByIdMutation, pgx.NamedArgs{"id": projectId})
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