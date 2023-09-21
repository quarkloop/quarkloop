package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/db/api"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

type DeleteApp struct {
	Name string `json:"name"`
}

type DeleteAppArgs struct {
	App model.App `json:"app"`
}

func (op *DeleteApp) Call(appId, instanceId string, args json.RawMessage) (interface{}, error) {
	var appArgs DeleteAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	err := api.DeleteApp(appArgs.App.AppId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
