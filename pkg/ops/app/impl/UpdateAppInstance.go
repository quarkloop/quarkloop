package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type UpdateAppInstanceArgs struct {
	Where struct {
		AppId      string `json:"appId" binding:"required"`
		InstanceId string `json:"instanceId" binding:"required"`
	} `json:"where" binding:"required"`
	Data struct {
		AppInstance model.AppInstance `json:"appInstance" binding:"required"`
	} `json:"data" binding:"required"`
	Select *struct {
		Id   *bool `json:"id"`
		Name *bool `json:"name"`
	} `json:"select"`
}

func UpdateAppInstance(args json.RawMessage) (interface{}, error) {
	var appInstanceArgs UpdateAppInstanceArgs
	if err := json.Unmarshal(args, &appInstanceArgs); err != nil {
		return nil, err
	}

	appInstance, err := db.UpdateAppInstance(&appInstanceArgs.Data.AppInstance)
	if err != nil {
		return nil, err
	}

	return appInstance, nil
}
