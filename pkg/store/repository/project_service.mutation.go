package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// CreateProjectService

const createProjectServiceMutation = `
INSERT INTO
  "system"."ProjectService" ("projectId", "name", "type", "description", "metadata", "data")
VALUES
  (@projectId, @name, @type, @description, @metadata, @data)
RETURNING 
  "id", "name", "type", "description", "metadata", "data", "createdAt";
`

func (r *Repository) CreateProjectService(ctx context.Context, projectId string, pService *model.ProjectService) (*model.ProjectService, error) {
	row := r.SystemDbConn.QueryRow(
		ctx,
		createProjectServiceMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"id":          pService.Id,
			"name":        pService.Name,
			"type":        pService.Type,
			"description": pService.Description,
			"metadata":    pService.Metadata,
			"data":        pService.Data,
		},
	)

	var service model.ProjectService
	err := row.Scan(
		&service.Id,
		&service.Name,
		&service.Type,
		&service.Description,
		&service.Metadata,
		&service.Data,
		&service.CreatedAt,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	return &service, nil
}

/// CreateBulkProjectService

func (r *Repository) CreateBulkProjectService(ctx context.Context, projectId string, pServiceList []model.ProjectService) (int64, error) {
	rowsAffected, err := r.SystemDbConn.CopyFrom(
		ctx,
		pgx.Identifier{"system", "ProjectService"},
		[]string{"projectId", "name", "type", "description", "data"},
		pgx.CopyFromSlice(len(pServiceList), func(i int) ([]interface{}, error) {
			return []interface{}{
				projectId,
				pServiceList[i].Name,
				pServiceList[i].Type,
				pServiceList[i].Description,
				//pServiceList[i].Metadata,
				pServiceList[i].Data,
			}, nil
		}),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return 0, err
	}

	if rowsAffected == 0 {
		notFoundErr := errors.New("cannot bulk create")
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", notFoundErr)
		return 0, notFoundErr
	}

	return rowsAffected, nil
}

/// UpdateProjectServiceById

const updateProjectServiceByIdMutation = `
UPDATE
  "system"."ProjectService"
SET
  "name"        = @name,
  "type"        = @type,
  "description" = @description,
  "metadata"    = @metadata,
  "data"        = @data,
  "updatedAt"   = @updatedAt
WHERE
  "id" = @id;
`

func (r *Repository) UpdateProjectServiceById(ctx context.Context, projectServiceId string, pService *model.ProjectService) error {
	commandTag, err := r.SystemDbConn.Exec(
		ctx,
		updateProjectServiceByIdMutation,
		pgx.NamedArgs{
			"id":          projectServiceId,
			"name":        pService.Name,
			"type":        pService.Type,
			"description": pService.Description,
			"metadata":    pService.Metadata,
			"data":        pService.Data,
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

/// DeleteProjectServiceById

const deleteProjectServiceByIdMutation = `
DELETE FROM
  "system"."ProjectService"
WHERE
  "id" = @id;
`

func (r *Repository) DeleteProjectServiceById(ctx context.Context, projectServiceId string) error {
	commandTag, err := r.SystemDbConn.Exec(ctx, deleteProjectServiceByIdMutation, pgx.NamedArgs{
		"id": projectServiceId,
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
