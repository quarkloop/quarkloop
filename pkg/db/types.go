package db

import (
	"time"
)

type DatabaseResponsePayload struct {
	Status struct {
		StatusCode       int       `json:"statusCode"`
		StatusCodeString string    `json:"statusCodeString"`
		DefaultMessage   string    `json:"defaultMessage"`
		Timestamp        time.Time `json:"time"`
	} `json:"status"`
	Database struct {
		App struct {
			Records      interface{} `json:"records"`
			TotalRecords int         `json:"totalRecords"`
		} `json:"app,omitempty"`
		AppInstance struct {
			Records      interface{} `json:"records"`
			TotalRecords int         `json:"totalRecords"`
		} `json:"appInstance,omitempty"`
	} `json:"database,omitempty"`
}
