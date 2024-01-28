package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/quarkloop/quarkloop/pkg/model"
	"github.com/quarkloop/quarkloop/pkg/service/project"
)

/// CreateProject

const createProjectMutation = `
INSERT INTO "system"."Project" (
    "orgId",
    "workspaceId",
    "sid",
    "name",
    "description",
    "visibility",
    "createdBy"
)
VALUES (
    @orgId,
    @workspaceId,
    @sid,
    @name,
    @description,
    @visibility,
    @createdBy
)
RETURNING 
    "id",
    "sid",
    "orgId",
    "workspaceId",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy";
`

func (store *projectStore) CreateProject(ctx context.Context, cmd *project.CreateProjectCommand) (*model.Project, error) {
	if cmd.ScopeId == "" {
		sid, err := gonanoid.New()
		if err != nil {
			return nil, err
		}
		cmd.ScopeId = sid
	}

	row := store.Conn.QueryRow(ctx, createProjectMutation, pgx.NamedArgs{
		"orgId":       cmd.OrgId,
		"workspaceId": cmd.WorkspaceId,
		"sid":         cmd.ScopeId,
		"name":        cmd.Name,
		"description": cmd.Description,
		"visibility":  cmd.Visibility,
		"createdBy":   cmd.CreatedBy,
	})

	var pr model.Project
	err := row.Scan(
		&pr.Id,
		&pr.ScopeId,
		&pr.OrgId,
		&pr.WorkspaceId,
		&pr.Name,
		&pr.Description,
		&pr.Visibility,
		&pr.CreatedAt,
		&pr.CreatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, project.HandleError(err)
	}

	return &pr, nil
}

/// UpdateProjectById

const updateProjectByIdMutation = `
UPDATE
    "system"."Project"
SET
    "sid"         = COALESCE (NULLIF(@sid, ''), "sid"),
    "name"        = COALESCE (NULLIF(@name, ''), "name"),
    "description" = COALESCE (NULLIF(@description, ''), "description"),
    "visibility"  = COALESCE (NULLIF(@visibility, 0), "visibility"),
    "updatedAt"   = @updatedAt,
    "updatedBy"   = @updatedBy
WHERE
    "id" = @id
AND
    "orgId" = @orgId
AND
    "workspaceId" = @workspaceId;	
`

func (store *projectStore) UpdateProjectById(ctx context.Context, cmd *project.UpdateProjectByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, updateProjectByIdMutation, pgx.NamedArgs{
		"orgId":       cmd.OrgId,
		"workspaceId": cmd.WorkspaceId,
		"id":          cmd.ProjectId,
		"sid":         cmd.ScopeId,
		"name":        cmd.Name,
		"description": cmd.Description,
		"visibility":  cmd.Visibility,
		"updatedBy":   cmd.UpdatedBy,
		"updatedAt":   time.Now(),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
		return project.HandleError(err)
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
    "id" = @id
AND
    "orgId" = @orgId
AND
    "workspaceId" = @workspaceId;		
`

func (store *projectStore) DeleteProjectById(ctx context.Context, cmd *project.DeleteProjectByIdCommand) error {
	commandTag, err := store.Conn.Exec(ctx, deleteProjectByIdMutation, pgx.NamedArgs{
		"orgId":       cmd.OrgId,
		"workspaceId": cmd.WorkspaceId,
		"id":          cmd.ProjectId,
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
