package model

type AppInstance struct {
	AppId      string `json:"appId"`
	InstanceId string `json:"instanceId"`
	ID         string `json:"id,omitempty"`
	Name       string `json:"name"`
}
