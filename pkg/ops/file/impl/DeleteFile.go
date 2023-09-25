package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/file/db"
)

type DeleteFileArgs struct {
	Where struct {
		AppId      string `json:"appId" binding:"required"`
		InstanceId string `json:"instanceId" binding:"required"`
		FileId     string `json:"fileId" binding:"required"`
	} `json:"where" binding:"required"`
}

func DeleteFile(args json.RawMessage) (interface{}, error) {
	var fileArgs DeleteFileArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	err := db.DeleteFile(fileArgs.Where.AppId, fileArgs.Where.InstanceId, fileArgs.Where.FileId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
