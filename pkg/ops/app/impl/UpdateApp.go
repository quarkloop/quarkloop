package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type UpdateAppArgs struct {
	AppID string    `json:"appId" binding:"required"`
	App   model.App `json:"app"`
}

func UpdateApp(args json.RawMessage) (interface{}, error) {
	var appArgs UpdateAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	app, err := db.UpdateApp(&appArgs.App)
	if err != nil {
		return nil, err
	}

	return app, nil
}
