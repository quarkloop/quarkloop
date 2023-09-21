package ops

import (
	"encoding/json"
	"errors"

	app "github.com/quarkloop/quarkloop/pkg/ops/app"
	file "github.com/quarkloop/quarkloop/pkg/ops/file"
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

var operations = Ops{
	Ops: map[string]interface{}{
		"CreateApp": app.CreateApp{
			Name: "CreateApp",
		},
		"UpdateApp": app.UpdateApp{
			Name: "UpdateApp",
		},
		"DeleteApp": app.DeleteApp{
			Name: "DeleteApp",
		},
		"GetFileById": file.GetFileById{
			Name: "GetFileById",
		},
		"CreateFile": file.CreateFile{
			Name: "CreateFile",
		},
		"UpdateFile": file.UpdateFile{
			Name: "UpdateFile",
		},
		"DeleteFile": file.DeleteFile{
			Name: "DeleteFile",
		},
	},
}

func FindOp(appId, instanceId, opName string, args json.RawMessage) (*OpCallCatalog, error) {
	val, ok := operations.Ops[opName]
	if ok {
		catalog := OpCallCatalog{
			AppId:      appId,
			InstanceId: instanceId,
			Name:       opName,
			Args:       args,
		}

		switch val := val.(type) {
		case app.CreateApp:
			catalog.CallSite = &val
			return &catalog, nil
		case app.UpdateApp:
			catalog.CallSite = &val
			return &catalog, nil
		case app.DeleteApp:
			catalog.CallSite = &val
			return &catalog, nil
		case file.GetFileById:
			catalog.CallSite = &val
			return &catalog, nil
		case file.CreateFile:
			catalog.CallSite = &val
			return &catalog, nil
		case file.UpdateFile:
			catalog.CallSite = &val
			return &catalog, nil
		case file.DeleteFile:
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
