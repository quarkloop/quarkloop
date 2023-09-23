package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/file/db"
)

type DeleteFile struct {
	Name string `json:"name"`
}

type DeleteFileArgs struct {
	AppID      string `json:"appId" binding:"required"`
	InstanceID string `json:"instanceId" binding:"required"`
	FileId     string `json:"fileId" binding:"required"`
}

func (op *DeleteFile) Call(args json.RawMessage) (interface{}, error) {
	var fileArgs DeleteFileArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	err := db.DeleteFile(fileArgs.AppID, fileArgs.InstanceID, fileArgs.FileId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
