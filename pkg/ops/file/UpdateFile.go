package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/file/db"
	"github.com/quarkloop/quarkloop/pkg/ops/file/model"
)

type UpdateFile struct {
	Name string `json:"name"`
}

type UpdateFileArgs struct {
	AppID      string     `json:"appId" binding:"required"`
	InstanceID string     `json:"instanceId" binding:"required"`
	File       model.File `json:"file" binding:"required"`
}

func (op *UpdateFile) Call(appId, instanceId string, args json.RawMessage) (interface{}, error) {
	var fileArgs UpdateFileArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	file, err := db.UpdateFile(fileArgs.AppID, fileArgs.InstanceID, &fileArgs.File)
	if err != nil {
		return nil, err
	}

	return file, nil
}
