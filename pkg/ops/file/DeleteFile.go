package ops

import (
	"encoding/json"
	"errors"

	"github.com/quarkloop/quarkloop/pkg/ops/file/impl"
)

type DeleteFile struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (op *DeleteFile) Call(args json.RawMessage) (interface{}, error) {
	if op.Version == "latest" {
		val, err := impl.DeleteFile(args)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return nil, errors.New("failed to call op")
}
