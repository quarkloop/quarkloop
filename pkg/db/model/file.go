package model

type File struct {
	InstanceId string `json:"appInstanceId"`
	ID         string `json:"id,omitempty"`
	Enable     bool   `json:"enable"`
}
