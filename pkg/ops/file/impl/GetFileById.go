package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/file/db"
)

type GetFileByIdArgs struct {
	AppID      string `json:"appId" binding:"required"`
	InstanceID string `json:"instanceId" binding:"required"`
	FileId     string `json:"fileId" binding:"required"`
}

func GetFileById(args json.RawMessage) (interface{}, error) {
	var fileArgs GetFileByIdArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	file, err := db.GetFileById(fileArgs.AppID, fileArgs.InstanceID, fileArgs.FileId)
	if err != nil {
		return nil, err
	}

	return file, nil
}
