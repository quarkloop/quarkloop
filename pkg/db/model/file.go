package model

type File struct {
	AppId      string `json:"appId"`
	InstanceId string `json:"instanceId"`
	Id         string `json:"id,omitempty"`
	Enable     bool   `json:"enable"`
}
