package api

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/db"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

func GetFileById(appId, instanceId string) (interface{}, error) {
	payload := model.File{
		AppId:      appId,
		InstanceId: instanceId,
		Enable:     false,
	}

	marshalled, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	res, err := db.HttpClientInstance.Get("http://localhost:3000/api/v1/tables/appFile", marshalled, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
