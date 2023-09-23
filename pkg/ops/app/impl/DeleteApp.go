package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type DeleteAppArgs struct {
	App model.App `json:"app"`
}

func DeleteApp(args json.RawMessage) (interface{}, error) {
	var appArgs DeleteAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	err := db.DeleteApp(appArgs.App.AppId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
