package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/quarkloop/quarkloop/pkg/db"
	"github.com/quarkloop/quarkloop/pkg/db/client"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

func GetFileById(appId, instanceId, fileId string) (*model.File, error) {
	params := url.Values{}
	params.Add("id", fileId)

	res, err := client.DatabaseClient.Get("http://localhost:3000/api/v1/tables/appFile", &params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var file model.File

		var payload db.DatabaseResponsePayload
		decodeErr := json.NewDecoder(res.Body).Decode(&payload)
		if decodeErr != nil {
			return nil, decodeErr
		}

		val, ok := payload.Database.File.Records.(model.File)
		if ok {
			file = val
			return &file, nil
		}
	}

	return nil, errors.New("file not found")
}
