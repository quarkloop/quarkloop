package ops

import (
	"encoding/json"
	"errors"

	"github.com/quarkloop/quarkloop/pkg/ops/app/impl"
)

type CreateApp struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (op *CreateApp) Call(args json.RawMessage) (interface{}, error) {
	if op.Version == "latest" {
		val, err := impl.CreateApp(args)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return nil, errors.New("failed to call op")
}
