package ops

import (
	"github.com/quarkloop/quarkloop/pkg/db/api"
)

type GetFileById struct {
	Name string `json:"name"`
}

func (op *GetFileById) Call(appId, instanceId string, args interface{}) (interface{}, error) {
	res, err := api.GetFileById(appId, instanceId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
