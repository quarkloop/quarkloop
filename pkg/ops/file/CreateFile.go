package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/db/api"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

type CreateFile struct {
	Name string `json:"name"`
}

type CreateFileArgs struct {
	File model.File `json:"file"`
}

func (op *CreateFile) Call(appId, instanceId string, args json.RawMessage) (interface{}, error) {
	var fileArgs CreateFileArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	file, err := api.CreateFile(appId, instanceId, &fileArgs.File)
	if err != nil {
		return nil, err
	}

	return file, nil
}
