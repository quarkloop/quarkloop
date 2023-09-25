package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type CreateAppInstanceArgs struct {
	Data struct {
		AppInstance model.AppInstance `json:"appInstance" binding:"required"`
	} `json:"data" binding:"required"`
	Select *struct {
		Id   *bool `json:"id"`
		Name *bool `json:"name"`
	} `json:"select"`
}

func CreateAppInstance(args json.RawMessage) (interface{}, error) {
	var appInstanceArgs CreateAppInstanceArgs
	if err := json.Unmarshal(args, &appInstanceArgs); err != nil {
		return nil, err
	}

	appInstance, err := db.CreateAppInstance(&appInstanceArgs.Data.AppInstance)
	if err != nil {
		return nil, err
	}

	return appInstance, nil
}
