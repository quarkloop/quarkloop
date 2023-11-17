package impl

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/ops/file/db"
)

type GetFileByIdArgs struct {
	Where struct {
		AppId      string `json:"projectId" binding:"required"`
		InstanceId string `json:"instanceId" binding:"required"`
		FileId     string `json:"fileId" binding:"required"`
	} `json:"where" binding:"required"`
	Select *struct {
		Id   *bool `json:"id"`
		Name *bool `json:"name"`
	} `json:"select"`
}

func GetFileById(args json.RawMessage) (interface{}, error) {
	var fileArgs GetFileByIdArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	file, err := db.GetFileById(fileArgs.Where.AppId, fileArgs.Where.InstanceId, fileArgs.Where.FileId)
	if err != nil {
		return nil, err
	}

	return file, nil
}
