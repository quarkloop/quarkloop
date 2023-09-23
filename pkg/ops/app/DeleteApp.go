package ops

import (
	"encoding/json"
	"errors"

	"github.com/quarkloop/quarkloop/pkg/ops/app/impl"
)

type DeleteApp struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (op *DeleteApp) Call(args json.RawMessage) (interface{}, error) {
	if op.Version == "latest" {
		val, err := impl.DeleteApp(args)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return nil, errors.New("failed to call op")
}
