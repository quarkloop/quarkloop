package store

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
	projectErrors "github.com/quarkloop/quarkloop/services/project/errors"
)

/// GetProjectId

const getProjectIdQuery = `
SELECT 
    org."id",
    ws."id",
	p."id"
FROM 
    "system"."Project" AS p
LEFT JOIN 
    "system"."Org" AS org ON org."id" = p."orgId"
LEFT JOIN 
    system."Workspace" AS ws ON ws."id" = p."workspaceId"
WHERE (
	org."sid" = @orgSid
	AND
	ws."sid" = @workspaceSid
	AND
	p."sid" = @projectSid
);
`

type GetProjectIdQuery struct {
	OrgSid       string
	WorkspaceSid string
	ProjectSid   string
}

func (store *projectStore) GetProjectId(ctx context.Context, query *GetProjectIdQuery) (int64, int64, int64, error) {
	row := store.Conn.QueryRow(ctx, getProjectIdQuery, pgx.NamedArgs{
		"orgSid":       query.OrgSid,
		"workspaceSid": query.WorkspaceSid,
		"projectSid":   query.ProjectSid,
	})

	var orgId, workspaceId, projectId int64
	err := row.Scan(&orgId, &workspaceId, &projectId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return orgId, workspaceId, projectId, projectErrors.ErrProjectNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return orgId, workspaceId, projectId, err
	}

	return orgId, workspaceId, projectId, nil
}

/// GetProjectById

const getProjectByIdQuery = `
SELECT 
    p."id",
    p."sid",
    p."orgId",
    p."workspaceId",
    org."sid",
    ws."sid",
    p."name",
    p."description",
    p."visibility",
    p."createdAt",
    p."createdBy",
    p."updatedAt",
    p."updatedBy"
FROM 
    "system"."Project" AS p
LEFT JOIN 
    "system"."Org" AS org ON org."id" = p."orgId"
LEFT JOIN 
    system."Workspace" AS ws ON ws."id" = p."workspaceId"
WHERE (
	p."orgId" = @orgId
	AND
	p."workspaceId" = @workspaceId
	AND
	p."id" = @id
);
`

type GetProjectByIdQuery struct {
	OrgId       int64
	WorkspaceId int64
	ProjectId   int64
}

func (store *projectStore) GetProjectById(ctx context.Context, query *GetProjectByIdQuery) (*model.Project, error) {
	row := store.Conn.QueryRow(ctx, getProjectByIdQuery, pgx.NamedArgs{
		"orgId":       query.OrgId,
		"workspaceId": query.WorkspaceId,
		"id":          query.ProjectId,
	})

	var p model.Project
	err := row.Scan(
		&p.Id,
		&p.ScopeId,
		&p.OrgId,
		&p.WorkspaceId,
		&p.OrgScopeId,
		&p.WorkspaceScopeId,
		&p.Name,
		&p.Description,
		&p.Visibility,
		&p.CreatedAt,
		&p.CreatedBy,
		&p.UpdatedAt,
		&p.UpdatedBy,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, projectErrors.ErrProjectNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &p, nil
}

/// GetProjectVisibilityById

const getProjectVisibilityByIdQuery = `
SELECT 
    "visibility"
FROM 
    "system"."Project"
WHERE (
	"orgId" = @orgId
	AND
	"workspaceId" = @workspaceId
	AND
	"id" = @id
);
`

type GetProjectVisibilityByIdQuery struct {
	OrgId       int64
	WorkspaceId int64
	ProjectId   int64
}

func (store *projectStore) GetProjectVisibilityById(ctx context.Context, query *GetProjectVisibilityByIdQuery) (model.ScopeVisibility, error) {
	row := store.Conn.QueryRow(ctx, getProjectVisibilityByIdQuery, pgx.NamedArgs{
		"orgId":       query.OrgId,
		"workspaceId": query.WorkspaceId,
		"id":          query.ProjectId,
	})

	var visibility model.ScopeVisibility
	if err := row.Scan(&visibility); err != nil {
		if err == pgx.ErrNoRows {
			return 0, projectErrors.ErrProjectNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return 0, err
	}

	return visibility, nil
}

/// GetProjectList

const listProjectsQuery = `
SELECT 
    p."id",
    p."sid",
    p."orgId",
    p."workspaceId",
    org."sid",
    ws."sid",
    p."name",
    p."description",
    p."visibility",
    p."createdAt",
    p."createdBy",
    p."updatedAt",
    p."updatedBy"
FROM 
    "system"."Project" AS p
LEFT JOIN 
    system."Workspace" AS ws ON ws."id" = p."workspaceId"
LEFT JOIN 
    "system"."Org" AS org ON org."id" = p."orgId"	
WHERE
    p."id" = ANY (@projectIdList)
%s
`

type GetProjectListQuery struct {
	ProjectIdList []int64
	Visibility    model.ScopeVisibility
}

func (store *projectStore) GetProjectList(ctx context.Context, query *GetProjectListQuery) ([]*model.Project, error) {
	var finalQuery strings.Builder
	if query.Visibility == model.PublicVisibility || query.Visibility == model.PrivateVisibility {
		finalQuery.WriteString(fmt.Sprintf(listProjectsQuery, `AND p."visibility" = @visibility;`))
	} else {
		finalQuery.WriteString(fmt.Sprintf(listProjectsQuery, ";"))
	}

	rows, err := store.Conn.Query(ctx, finalQuery.String(), pgx.NamedArgs{
		"projectIdList": query.ProjectIdList,
		"visibility":    query.Visibility,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var projectList []*model.Project = []*model.Project{}
	for rows.Next() {
		var project model.Project
		err := rows.Scan(
			&project.Id,
			&project.ScopeId,
			&project.OrgId,
			&project.WorkspaceId,
			&project.OrgScopeId,
			&project.WorkspaceScopeId,
			&project.Name,
			&project.Description,
			&project.Visibility,
			&project.CreatedAt,
			&project.CreatedBy,
			&project.UpdatedAt,
			&project.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		projectList = append(projectList, &project)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return projectList, nil
}

// /// GetUserAssignmentList

// const getUserAssignmentListQuery = `
// SELECT
//     ua.id,
//     ua."userId",
//     ur."name",
//     ua."createdAt",
//     ua."createdBy",
//     ua."updatedAt",
//     ua."updatedBy"
// FROM
//     "system"."UserAssignment" AS ua
// LEFT JOIN
//     "system"."UserRole" AS ur ON ur.id = ua."userRoleId"
// WHERE (
// 	ua."orgId" = @orgId
// 	AND
// 	ua."workspaceId" = @workspaceId
// 	AND
// 	ua."projectId" = @projectId
// )
// GROUP BY
//     ua.id,
//     ur."name"
// ORDER BY id ASC;
// `

// type GetUserAssignmentListQuery struct {
// 	OrgId       int64
// 	WorkspaceId int64
// 	ProjectId   int64
// }

// func (store *projectStore) GetUserAssignmentList(ctx context.Context, query *GetUserAssignmentListQuery) ([]*user.UserAssignment, error) {
// 	rows, err := store.Conn.Query(ctx, getUserAssignmentListQuery, pgx.NamedArgs{
// 		"orgId":       query.OrgId,
// 		"workspaceId": query.WorkspaceId,
// 		"projectId":   query.ProjectId,
// 	})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var uaList []*user.UserAssignment = []*user.UserAssignment{}

// 	for rows.Next() {
// 		var ua user.UserAssignment
// 		err := rows.Scan(
// 			&ua.Id,
// 			&ua.UserId,
// 			&ua.Role,
// 			&ua.CreatedAt,
// 			&ua.CreatedBy,
// 			&ua.UpdatedAt,
// 			&ua.UpdatedBy,
// 		)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
// 			return nil, err
// 		}

// 		uaList = append(uaList, &ua)
// 	}

// 	if err := rows.Err(); err != nil {
// 		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
// 		return nil, err
// 	}

// 	return uaList, nil
// }

// /// GetProject

// const getProjectQuery = `
// SELECT
//     p."id",
//     p."sid",
//     p."orgId",
//     p."workspaceId",
//     org."sid",
//     ws."sid",
//     p."name",
//     p."description",
//     p."visibility",
//     p."createdAt",
//     p."createdBy",
//     p."updatedAt",
//     p."updatedBy"
// FROM
//     "system"."Project" AS p
// LEFT JOIN
//     "system"."Org" AS org ON org."id" = p."orgId"
// LEFT JOIN
//     system."Workspace" AS ws ON ws."id" = p."workspaceId"
// WHERE
//     %s
// ORDER BY "updatedAt" ASC
// LIMIT 1;
// `

// func (store *projectStore) GetProject(ctx context.Context, p *model.Project) (*model.Project, error) {
// 	availableFields := []string{}
// 	projectFields := map[string]interface{}{
// 		"sid":        p.ScopeId,
// 		"name":       p.Name,
// 		"visibility": p.Visibility,
// 		"createdAt":  p.CreatedAt,
// 		"updatedAt":  p.UpdatedAt,
// 	}
// 	for key, value := range projectFields {
// 		switch v := value.(type) {
// 		case int:
// 			if v != 0 {
// 				availableFields = append(availableFields, fmt.Sprintf("p.\"%s\" = '%d'", key, v))
// 			}
// 		case float64:
// 			if v != 0.0 {
// 				availableFields = append(availableFields, fmt.Sprintf("p.\"%s\" = '%f'", key, v))
// 			}
// 		case string:
// 			if v != "" {
// 				availableFields = append(availableFields, fmt.Sprintf("p.\"%s\" = '%s'", key, v))
// 			}
// 		case *time.Time:
// 			if v != nil {
// 				availableFields = append(availableFields, fmt.Sprintf("p.\"%s\" = '%s'", key, v))
// 			}
// 		}
// 	}
// 	finalQuery := fmt.Sprintf(getProjectQuery, strings.Join(availableFields, " AND "))

// 	row := store.Conn.QueryRow(ctx, finalQuery)

// 	var project model.Project
// 	err := row.Scan(
// 		&project.Id,
// 		&project.ScopeId,
// 		&project.OrgId,
// 		&project.WorkspaceId,
// 		&project.OrgScopeId,
// 		&project.WorkspaceScopeId,
// 		&project.Name,
// 		&project.Description,
// 		&project.Visibility,
// 		&project.CreatedAt,
// 		&project.CreatedBy,
// 		&project.UpdatedAt,
// 		&project.UpdatedBy,
// 	)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
// 		return nil, err
// 	}

// 	return &project, nil
// }
