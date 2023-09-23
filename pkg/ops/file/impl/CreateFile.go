package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/file/db"
	"github.com/quarkloop/quarkloop/pkg/ops/file/model"
)

type CreateFileArgs struct {
	AppID      string     `json:"appId" binding:"required"`
	InstanceID string     `json:"instanceId" binding:"required"`
	File       model.File `json:"file" binding:"required"`
}

func CreateFile(args json.RawMessage) (interface{}, error) {
	var fileArgs CreateFileArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	file, err := db.CreateFile(fileArgs.AppID, fileArgs.InstanceID, &fileArgs.File)
	if err != nil {
		return nil, err
	}

	return file, nil
}
