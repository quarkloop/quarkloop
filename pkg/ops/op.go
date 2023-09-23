package ops

import (
	"encoding/json"
	"errors"

	app "github.com/quarkloop/quarkloop/pkg/ops/app"
	file "github.com/quarkloop/quarkloop/pkg/ops/file"
)

type OpCall interface {
	Call(args json.RawMessage) (interface{}, error)
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
		// App
		"CreateApp": app.CreateApp{
			Name:    "CreateApp",
			Version: "latest",
		},
		"UpdateApp": app.UpdateApp{
			Name:    "UpdateApp",
			Version: "latest",
		},
		"DeleteApp": app.DeleteApp{
			Name:    "DeleteApp",
			Version: "latest",
		},
		// AppInstance
		"CreateAppInstance": app.CreateAppInstance{
			Name:    "CreateAppInstance",
			Version: "latest",
		},
		"UpdateAppInstance": app.UpdateAppInstance{
			Name:    "UpdateAppInstance",
			Version: "latest",
		},
		"DeleteAppInstance": app.DeleteAppInstance{
			Name:    "DeleteAppInstance",
			Version: "latest",
		},
		// File
		"GetFileById": file.GetFileById{
			Name:    "GetFileById",
			Version: "latest",
		},
		"CreateFile": file.CreateFile{
			Name:    "CreateFile",
			Version: "latest",
		},
		"UpdateFile": file.UpdateFile{
			Name:    "UpdateFile",
			Version: "latest",
		},
		"DeleteFile": file.DeleteFile{
			Name:    "DeleteFile",
			Version: "latest",
		},
	},
}

func FindOp(opName string, args json.RawMessage) (*OpCallCatalog, error) {
	val, ok := operations.Ops[opName]
	if ok {
		catalog := OpCallCatalog{
			Name: opName,
			Args: args,
		}

		switch val := val.(type) {
		// App
		case app.CreateApp:
			catalog.CallSite = &val
			return &catalog, nil
		case app.UpdateApp:
			catalog.CallSite = &val
			return &catalog, nil
		case app.DeleteApp:
			catalog.CallSite = &val
			return &catalog, nil
		// AppInstance
		case app.CreateAppInstance:
			catalog.CallSite = &val
			return &catalog, nil
		case app.UpdateAppInstance:
			catalog.CallSite = &val
			return &catalog, nil
		case app.DeleteAppInstance:
			catalog.CallSite = &val
			return &catalog, nil
		// File
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
	res, err := opc.CallSite.Call(opc.Args)
	if err != nil {
		return nil, err
	}

	return res, err
}
