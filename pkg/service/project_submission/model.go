package project_submission

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetProjectSubmissionListParams struct {
	Context   context.Context
	ProjectId string
}

type GetProjectSubmissionByIdParams struct {
	Context             context.Context
	ProjectId           string
	ProjectSubmissionId string
}

type CreateProjectSubmissionParams struct {
	Context           context.Context
	UserId            string
	ProjectId         string
	ProjectSubmission *model.ProjectSubmission
}

type UpdateProjectSubmissionByIdParams struct {
	Context             context.Context
	ProjectSubmissionId string
	ProjectSubmission   *model.ProjectSubmission
}

type DeleteProjectSubmissionByIdParams struct {
	Context             context.Context
	ProjectSubmissionId string
}
