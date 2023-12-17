package project_submission

import (
	"github.com/quarkloop/quarkloop/pkg/model"
)

type GetAppSubmissionListParams struct {
	ProjectId int
}

type GetAppSubmissionByIdParams struct {
	ProjectId       int
	AppSubmissionId string
}

type CreateAppSubmissionParams struct {
	UserId        string
	ProjectId     int
	AppSubmission *model.AppSubmission
}

type UpdateAppSubmissionByIdParams struct {
	AppSubmissionId string
	AppSubmission   *model.AppSubmission
}

type DeleteAppSubmissionByIdParams struct {
	AppSubmissionId string
}
