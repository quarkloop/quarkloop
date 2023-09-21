package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/db/api"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

type CreateApp struct {
	Name string `json:"name"`
}

type CreateAppArgs struct {
	App model.App `json:"app"`
}

func (op *CreateApp) Call(appId, instanceId string, args json.RawMessage) (interface{}, error) {
	var appArgs CreateAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	app, err := api.CreateApp(&appArgs.App)
	if err != nil {
		return nil, err
	}

	return app, nil
}
