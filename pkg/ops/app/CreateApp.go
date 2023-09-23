package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type CreateApp struct {
	Name string `json:"name"`
}

type CreateAppArgs struct {
	App model.App `json:"app"`
}

func (op *CreateApp) Call(args json.RawMessage) (interface{}, error) {
	var appArgs CreateAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	app, err := db.CreateApp(&appArgs.App)
	if err != nil {
		return nil, err
	}

	return app, nil
}
