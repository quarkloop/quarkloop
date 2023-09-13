package api

import (
	"encoding/json"
	"time"
)

type DatabaseResponseStatus struct {
	StatusCode       int          `json:"statusCode"`
	StatusCodeString string       `json:"statusCodeString"`
	Timestamp        time.Time    `json:"timestamp"`
	Message          *string      `json:"message,omitempty"`
	Details          *interface{} `json:"details,omitempty"`
}

type DatabaseResponsePayload struct {
	Status   DatabaseResponseStatus `json:"status"`
	Database *struct {
		App *struct {
			Records      json.RawMessage `json:"records"`
			TotalRecords int             `json:"totalRecords"`
		} `json:"app,omitempty"`
		AppInstance *struct {
			Records      json.RawMessage `json:"records"`
			TotalRecords int             `json:"totalRecords"`
		} `json:"appInstance,omitempty"`
		File *struct {
			Records      json.RawMessage `json:"records"`
			TotalRecords int             `json:"totalRecords"`
		} `json:"appFile,omitempty"`
	} `json:"database,omitempty"`
}
