package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type CreateAppArgs struct {
	Data struct {
		Project model.Project `json:"app" binding:"required"`
	} `json:"data" binding:"required"`
	Select *struct {
		Id   *bool `json:"id"`
		Name *bool `json:"name"`
	} `json:"select"`
}

func CreateApp(args json.RawMessage) (interface{}, error) {
	var appArgs CreateAppArgs
	if err := json.Unmarshal(args, &appArgs); err != nil {
		return nil, err
	}

	app, err := db.CreateApp(&appArgs.Data.Project)
	if err != nil {
		return nil, err
	}

	return app, nil
}
