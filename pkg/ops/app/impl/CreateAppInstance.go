package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/app/db"
	"github.com/quarkloop/quarkloop/pkg/ops/app/model"
)

type CreateAppInstanceArgs struct {
	AppInstance model.AppInstance `json:"appInstance"`
}

func CreateAppInstance(args json.RawMessage) (interface{}, error) {
	var appInstanceArgs CreateAppInstanceArgs
	if err := json.Unmarshal(args, &appInstanceArgs); err != nil {
		return nil, err
	}

	appInstance, err := db.CreateAppInstance(&appInstanceArgs.AppInstance)
	if err != nil {
		return nil, err
	}

	return appInstance, nil
}
