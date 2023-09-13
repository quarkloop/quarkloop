package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/db/api"
)

type DeleteFile struct {
	Name string `json:"name"`
}

type DeleteFileArgs struct {
	FileId string `json:"fileId"`
}

func (op *DeleteFile) Call(appId, instanceId string, args json.RawMessage) (interface{}, error) {
	var fileArgs DeleteFileArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	file, err := api.DeleteFile(appId, instanceId, fileArgs.FileId)
	if err != nil {
		return nil, err
	}

	return file, nil
}
