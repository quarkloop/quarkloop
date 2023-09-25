package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type UpdateAppArgs struct {
	Where struct {
		AppId string `json:"appId" binding:"required"`
	} `json:"where" binding:"required"`
	Data struct {
		App model.App `json:"app" binding:"required"`
	} `json:"data" binding:"required"`
	Select *struct {
		Id   *bool `json:"id"`
		Name *bool `json:"name"`
	} `json:"select"`
}

func UpdateApp(args json.RawMessage) (interface{}, error) {
	var appArgs UpdateAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	app, err := db.UpdateApp(&appArgs.Data.App)
	if err != nil {
		return nil, err
	}

	return app, nil
}
