package model

import "github.com/quarkloop/quarkloop/pkg/db"

type File struct {
	Status   db.DatabaseResponseStatus `json:"status"`
	Database struct {
		AppId      string `json:"appId"`
		InstanceId string `json:"instanceId"`
		Id         string `json:"id,omitempty"`
		Enable     bool   `json:"enable"`
	} `json:"database,omitempty"`
}
