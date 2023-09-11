package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/db/api"
)

type GetFileById struct {
	Name string `json:"name"`
}

type GetFileByIdArgs struct {
	FileId string `json:"fileId"`
}

func (op *GetFileById) Call(appId, instanceId string, args json.RawMessage) (interface{}, error) {
	var fileArgs GetFileByIdArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	res, err := api.GetFileById(appId, instanceId, fileArgs.FileId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
