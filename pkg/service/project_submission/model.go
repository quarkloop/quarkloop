package project_submission

import (
	"context"

	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetAppSubmissionListParams struct {
	Context   context.Context
	ProjectId int
}

type GetAppSubmissionByIdParams struct {
	Context         context.Context
	ProjectId       int
	AppSubmissionId string
}

type CreateAppSubmissionParams struct {
	Context       context.Context
	UserId        string
	ProjectId     int
	AppSubmission *model.AppSubmission
}

type UpdateAppSubmissionByIdParams struct {
	Context         context.Context
	AppSubmissionId string
	AppSubmission   *model.AppSubmission
}

type DeleteAppSubmissionByIdParams struct {
	Context         context.Context
	AppSubmissionId string
}
