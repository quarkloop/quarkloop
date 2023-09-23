package ops

import (
	"encoding/json"
	"errors"

	"github.com/quarkloop/quarkloop/pkg/ops/app/impl"
)

type DeleteAppInstance struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (op *DeleteAppInstance) Call(args json.RawMessage) (interface{}, error) {
	if op.Version == "latest" {
		val, err := impl.DeleteAppInstance(args)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return nil, errors.New("failed to call op")
}
