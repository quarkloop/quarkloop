package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type DeleteAppInstanceArgs struct {
	AppID       string            `json:"appId" binding:"required"`
	AppInstance model.AppInstance `json:"appInstance"`
}

func DeleteAppInstance(args json.RawMessage) (interface{}, error) {
	var appInstanceArgs DeleteAppInstanceArgs
	if err := json.Unmarshal(args, &appInstanceArgs); err != nil {
		return nil, err
	}

	err := db.DeleteAppInstance(appInstanceArgs.AppInstance.AppId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
