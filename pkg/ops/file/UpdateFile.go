package ops

import (
	"encoding/json"

	"github.com/quarkloop/quarkloop/pkg/db/api"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

type UpdateFile struct {
	Name string `json:"name"`
}

type UpdateFileArgs struct {
	File model.File `json:"file"`
}

func (op *UpdateFile) Call(appId, instanceId string, args json.RawMessage) (interface{}, error) {
	var fileArgs UpdateFileArgs
	if err := json.Unmarshal(args, &fileArgs); err != nil {
		return nil, err
	}

	file, err := api.UpdateFile(appId, instanceId, &fileArgs.File)
	if err != nil {
		return nil, err
	}

	return file, nil
}
