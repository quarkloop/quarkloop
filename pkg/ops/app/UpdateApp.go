package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/db/api"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

type UpdateApp struct {
	Name string `json:"name"`
}

type UpdateAppArgs struct {
	App model.App `json:"app"`
}

func (op *UpdateApp) Call(appId, instanceId string, args json.RawMessage) (interface{}, error) {
	var appArgs UpdateAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	app, err := api.UpdateApp(&appArgs.App)
	if err != nil {
		return nil, err
	}

	return app, nil
}
