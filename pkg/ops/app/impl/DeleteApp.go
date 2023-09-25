package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
)

type DeleteAppArgs struct {
	Where struct {
		AppId string `json:"appId" binding:"required"`
	} `json:"where" binding:"required"`
}

func DeleteApp(args json.RawMessage) (interface{}, error) {
	var appArgs DeleteAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	err := db.DeleteApp(appArgs.Where.AppId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
