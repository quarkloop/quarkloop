package ops

import (
	"encoding/json"
	"errors"

	ops "github.com/quarkloop/quarkloop/pkg/ops/file"
)

type OpCall interface {
	Call(appId, instanceId string, args json.RawMessage) (interface{}, error)
}

type OpCallCatalog struct {
	AppId      string
	InstanceId string
	Name       string
	Args       json.RawMessage
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
		"CreateFile": ops.CreateFile{
			Name: "CreateFile",
		},
	},
}

func FindOp(appId, instanceId, opName string, args json.RawMessage) (*OpCallCatalog, error) {
	val, ok := Opss.Ops[opName]
	if ok {
		catalog := OpCallCatalog{
			AppId:      appId,
			InstanceId: instanceId,
			Name:       opName,
			Args:       args,
		}

		switch val := val.(type) {
		case ops.GetFileById:
			catalog.CallSite = &val
			return &catalog, nil
		case ops.CreateFile:
			catalog.CallSite = &val
			return &catalog, nil
		}

		return nil, errors.New("cannot convert op")
	}

	return nil, errors.New("cannot find op")
}

func (opc *OpCallCatalog) Exec() (interface{}, error) {
	res, err := opc.CallSite.Call(opc.AppId, opc.InstanceId, opc.Args)
	if err != nil {
		return nil, err
	}

	return res, err
}
