package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
)

type DeleteAppInstanceArgs struct {
	Where struct {
		AppId      string `json:"projectId" binding:"required"`
		InstanceId string `json:"instanceId" binding:"required"`
	} `json:"where" binding:"required"`
}

func DeleteAppInstance(args json.RawMessage) (interface{}, error) {
	var appInstanceArgs DeleteAppInstanceArgs
	if err := json.Unmarshal(args, &appInstanceArgs); err != nil {
		return nil, err
	}

	err := db.DeleteAppInstance(appInstanceArgs.Where.InstanceId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
