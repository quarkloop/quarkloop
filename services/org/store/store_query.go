package store

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
	orgErrors "github.com/quarkloop/quarkloop/services/org/errors"
)

/// GetOrgId

const getOrgIdQuery = `
SELECT 
    "id"
FROM 
    "system"."Org"
WHERE 
    "sid" = @sid;
`

func (store *orgStore) GetOrgId(ctx context.Context, orgSid string) (int64, error) {
	row := store.Conn.QueryRow(ctx, getOrgIdQuery, pgx.NamedArgs{"sid": orgSid})

	var orgId int64
	err := row.Scan(&orgId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return orgId, orgErrors.ErrOrgNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return orgId, err
	}

	return orgId, nil
}

/// GetOrgById

const getOrgByIdQuery = `
SELECT 
    "id",
    "sid",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."Org"
WHERE 
    "id" = @id;
`

type GetOrgByIdQuery struct {
	OrgId int64
}

func (store *orgStore) GetOrgById(ctx context.Context, query *GetOrgByIdQuery) (*model.Org, error) {
	row := store.Conn.QueryRow(ctx, getOrgByIdQuery, pgx.NamedArgs{"id": query.OrgId})

	var o model.Org
	err := row.Scan(
		&o.Id,
		&o.ScopeId,
		&o.Name,
		&o.Description,
		&o.Visibility,
		&o.CreatedAt,
		&o.CreatedBy,
		&o.UpdatedAt,
		&o.UpdatedBy,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, orgErrors.ErrOrgNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &o, nil
}

/// GetOrgVisibilityById

const getOrgVisibilityByIdQuery = `
SELECT 
    "visibility"
FROM 
    "system"."Org"
WHERE 
    "id" = @id;
`

type GetOrgVisibilityByIdQuery struct {
	OrgId int64
}

func (store *orgStore) GetOrgVisibilityById(ctx context.Context, query *GetOrgVisibilityByIdQuery) (model.ScopeVisibility, error) {
	row := store.Conn.QueryRow(ctx, getOrgVisibilityByIdQuery, pgx.NamedArgs{"id": query.OrgId})

	var visibility model.ScopeVisibility
	if err := row.Scan(&visibility); err != nil {
		if err == pgx.ErrNoRows {
			return "", orgErrors.ErrOrgNotFound
		}
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return "", err
	}

	return visibility, nil
}

// /// GetOrg

// const getOrgQuery = `
// SELECT
//     "id",
//     "sid",
//     "name",
//     "description",
//     "visibility",
//     "createdAt",
//     "createdBy",
//     "updatedAt",
//     "updatedBy"
// FROM
//     "system"."Org"
// WHERE
// `

// func (store *orgStore) GetOrg(ctx context.Context, organization *model.Org) (*model.Org, error) {
// 	availableFields := []string{}
// 	organizationFields := map[string]interface{}{
// 		"sid":        organization.ScopeId,
// 		"name":       organization.Name,
// 		"visibility": organization.Visibility,
// 		"createdAt":  organization.CreatedAt,
// 		"updatedAt":  organization.UpdatedAt,
// 	}
// 	for key, value := range organizationFields {
// 		switch v := value.(type) {
// 		case int:
// 			if v != 0 {
// 				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%d'", key, v))
// 			}
// 		case float64:
// 			if v != 0.0 {
// 				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%f'", key, v))
// 			}
// 		case string:
// 			if v != "" {
// 				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%s'", key, v))
// 			}
// 		case *time.Time:
// 			if v != nil {
// 				availableFields = append(availableFields, fmt.Sprintf("\"%s\" = '%s'", key, v))
// 			}
// 		}
// 	}
// 	finalQuery := getOrgQuery + strings.Join(availableFields, " AND ")

// 	row := store.Conn.QueryRow(ctx, finalQuery)

// 	var org model.Org
// 	err := row.Scan(
// 		&org.Id,
// 		&org.ScopeId,
// 		&org.Name,
// 		&org.Description,
// 		&org.Visibility,
// 		&org.CreatedAt,
// 		&org.CreatedBy,
// 		&org.UpdatedAt,
// 		&org.UpdatedBy,
// 	)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
// 		return nil, err
// 	}

// 	return &org, nil
// }

/// GetOrgList

const listOrgsQuery = `
SELECT 
    "id",
    "sid",
    "name",
    "description",
    "visibility",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
    "system"."Org"
WHERE 
    "id" = ANY (@orgIdList)
%s	
`

type GetOrgListQuery struct {
	OrgIdList  []int64
	Visibility model.ScopeVisibility
}

func (store *orgStore) GetOrgList(ctx context.Context, query *GetOrgListQuery) ([]*model.Org, error) {
	var finalQuery strings.Builder
	if query.Visibility == model.PublicVisibility || query.Visibility == model.PrivateVisibility {
		finalQuery.WriteString(fmt.Sprintf(listOrgsQuery, `AND "visibility" = @visibility;`))
	} else {
		finalQuery.WriteString(fmt.Sprintf(listOrgsQuery, ";"))
	}

	rows, err := store.Conn.Query(ctx, finalQuery.String(), pgx.NamedArgs{
		"orgIdList":  query.OrgIdList,
		"visibility": query.Visibility,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var orgList []*model.Org = []*model.Org{}

	for rows.Next() {
		var org model.Org
		err := rows.Scan(
			&org.Id,
			&org.ScopeId,
			&org.Name,
			&org.Description,
			&org.Visibility,
			&org.CreatedAt,
			&org.CreatedBy,
			&org.UpdatedAt,
			&org.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		orgList = append(orgList, &org)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return orgList, nil
}

/// GetWorkspaceList

const listWorkspacesQuery = `
SELECT 
    ws."id",
    ws."sid",
    ws."orgId",
    org."sid",
    ws."name",
    ws."description",
    ws."visibility",
    ws."createdAt",
    ws."createdBy",
    ws."updatedAt",
    ws."updatedBy"
FROM 
    "system"."Workspace" AS ws
LEFT JOIN 
    "system"."Org" AS org ON org."id" = ws."orgId"
WHERE 
    ws."orgId" = @orgId
%s	
`

type GetWorkspaceListQuery struct {
	OrgId      int64
	Visibility model.ScopeVisibility
}

func (store *orgStore) GetWorkspaceList(ctx context.Context, query *GetWorkspaceListQuery) ([]*model.Workspace, error) {
	var finalQuery strings.Builder
	if query.Visibility == model.PublicVisibility || query.Visibility == model.PrivateVisibility {
		finalQuery.WriteString(fmt.Sprintf(listWorkspacesQuery, `AND ws."visibility" = @visibility;`))
	} else {
		finalQuery.WriteString(fmt.Sprintf(listWorkspacesQuery, ";"))
	}

	rows, err := store.Conn.Query(ctx, finalQuery.String(), pgx.NamedArgs{
		"orgId":      query.OrgId,
		"visibility": query.Visibility,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var wsList []*model.Workspace = []*model.Workspace{}
	for rows.Next() {
		var workspace model.Workspace
		err := rows.Scan(
			&workspace.Id,
			&workspace.ScopeId,
			&workspace.OrgId,
			&workspace.OrgScopeId,
			&workspace.Name,
			&workspace.Description,
			&workspace.Visibility,
			&workspace.CreatedAt,
			&workspace.CreatedBy,
			&workspace.UpdatedAt,
			&workspace.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		wsList = append(wsList, &workspace)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return wsList, nil
}

// /// GetProjectList

// const listProjectsQuery = `
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
//     "system"."Workspace" AS ws ON ws.id = p."workspaceId"
// LEFT JOIN
//     "system"."Org" AS org ON org.id = p."orgId"
// WHERE
//     p."orgId" = @orgId
// %s
// `

// type GetProjectListQuery struct {
// 	OrgId      int64
// 	Visibility model.ScopeVisibility
// }

// func (store *orgStore) GetProjectList(ctx context.Context, query *GetProjectListQuery) ([]*model.Project, error) {
// 	var finalQuery strings.Builder
// 	if query.Visibility == model.PublicVisibility || query.Visibility == model.PrivateVisibility {
// 		finalQuery.WriteString(fmt.Sprintf(listProjectsQuery, `AND p."visibility" = @visibility;`))
// 	} else {
// 		finalQuery.WriteString(fmt.Sprintf(listProjectsQuery, ";"))
// 	}

// 	rows, err := store.Conn.Query(ctx, finalQuery.String(), pgx.NamedArgs{
// 		"orgId":      query.OrgId,
// 		"visibility": query.Visibility,
// 	})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var projectList []*model.Project = []*model.Project{}

// 	for rows.Next() {
// 		var project model.Project
// 		err := rows.Scan(
// 			&project.Id,
// 			&project.ScopeId,
// 			&project.OrgId,
// 			&project.WorkspaceId,
// 			&project.OrgScopeId,
// 			&project.WorkspaceScopeId,
// 			&project.Name,
// 			&project.Description,
// 			&project.Visibility,
// 			&project.CreatedAt,
// 			&project.CreatedBy,
// 			&project.UpdatedAt,
// 			&project.UpdatedBy,
// 		)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
// 			return nil, err
// 		}

// 		projectList = append(projectList, &project)
// 	}

// 	if err := rows.Err(); err != nil {
// 		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
// 		return nil, err
// 	}

// 	return projectList, nil
// }

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
// WHERE
//     ua."orgId" = @orgId
// GROUP BY
//     ua.id,
//     ur."name"
// ORDER BY id ASC;
// `

// func (store *orgStore) GetUserAssignmentList(ctx context.Context, query *org.GetUserAssignmentListQuery) ([]*user.UserAssignment, error) {
// 	rows, err := store.Conn.Query(ctx, getUserAssignmentListQuery, pgx.NamedArgs{"orgId": query.OrgId})
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
