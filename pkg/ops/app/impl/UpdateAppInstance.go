package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type UpdateAppInstanceArgs struct {
	AppID       string            `json:"appId" binding:"required"`
	AppInstance model.AppInstance `json:"appInstance"`
}

func UpdateAppInstance(args json.RawMessage) (interface{}, error) {
	var appInstanceArgs UpdateAppInstanceArgs
	if err := json.Unmarshal(args, &appInstanceArgs); err != nil {
		return nil, err
	}

	appInstance, err := db.UpdateAppInstance(&appInstanceArgs.AppInstance)
	if err != nil {
		return nil, err
	}

	return appInstance, nil
}
