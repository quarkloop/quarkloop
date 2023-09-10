package ops

import (
	"encoding/json"
	"errors"

	ops "github.com/quarkloop/quarkloop/pkg/ops/file"
)

type OpCall interface {
	Call(appId, instanceId string, args interface{}) (interface{}, error)
}

type OpCallCatalog struct {
	AppId      string
	InstanceId string
	Name       string
	Args       interface{}
	CallSite   OpCall
}

type Ops struct {
	Ops map[string]interface{}
}

// type OpApp struct {
// 	Name      string
// 	Settings  struct{}
// 	Instances map[string]interface{}
// }

var Opss = Ops{
	Ops: map[string]interface{}{
		"GetFileById": ops.GetFileById{
			Name: "GetFileById",
		},
	},
}

func FindOp(appId, instanceId, opName string, args json.RawMessage) (*OpCallCatalog, error) {
	val, ok := Opss.Ops[opName]
	if ok {
		switch val := val.(type) {
		case ops.GetFileById:
			catalog := OpCallCatalog{
				AppId:      appId,
				InstanceId: instanceId,
				Name:       opName,
				Args:       args,
				CallSite:   &val,
			}
			return &catalog, nil
		}

		return nil, errors.New("cannot convert op")
	}

	return nil, errors.New("cannot find op")
}

func (opc *OpCallCatalog) Exec() (interface{}, error) {
	res, err := opc.CallSite.Call(opc.AppId, opc.InstanceId, 5)
	if err != nil {
		return nil, err
	}

	return res, err
}
