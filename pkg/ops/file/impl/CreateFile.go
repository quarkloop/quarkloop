package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/file/db"
	"github.com/quarkloop/quarkloop/pkg/ops/file/model"
)

type CreateFileArgs struct {
	Where struct {
		AppId      string `json:"projectId" binding:"required"`
		InstanceId string `json:"instanceId" binding:"required"`
	} `json:"where" binding:"required"`
	Data struct {
		File model.File `json:"file" binding:"required"`
	} `json:"data" binding:"required"`
	Select *struct {
		Id   *bool `json:"id"`
		Name *bool `json:"name"`
	} `json:"select"`
}

func CreateFile(args json.RawMessage) (interface{}, error) {
	var fileArgs CreateFileArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	file, err := db.CreateFile(fileArgs.Where.AppId, fileArgs.Where.InstanceId, &fileArgs.Data.File)
	if err != nil {
		return nil, err
	}

	return file, nil
}
