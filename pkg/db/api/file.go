package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

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
		var payload DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return nil, err
		}

		var file model.File
		if err := json.Unmarshal(payload.Database.File.Records, &file); err != nil {
			return nil, err
		}

		return &file, nil
	}

	return nil, errors.New("file not found")
}

func CreateFile(appId, instanceId string, file *model.File) (*model.File, error) {
	payload, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}

	res, err := client.DatabaseClient.Post("http://localhost:3000/api/v1/tables/appFile", payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		var payload DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return nil, err
		}

		var file model.File
		if err := json.Unmarshal(payload.Database.File.Records, &file); err != nil {
			return nil, err
		}

		return &file, nil
	}

	return nil, errors.New("failed to create file")
}

func UpdateFile(appId, instanceId string, file *model.File) (*model.File, error) {
	payload, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}

	res, err := client.DatabaseClient.Update("http://localhost:3000/api/v1/tables/appFile", payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		var payload DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return nil, err
		}

		var file model.File
		if err := json.Unmarshal(payload.Database.File.Records, &file); err != nil {
			return nil, err
		}

		return &file, nil
	}

	return nil, errors.New("failed to update file")
}

func DeleteFile(appId, instanceId, fileId string) error {
	params := url.Values{}
	params.Add("id", fileId)

	res, err := client.DatabaseClient.Delete("http://localhost:3000/api/v1/tables/appFile", &params)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNoContent {
		var payload DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return err
		}

		return nil
	}

	return errors.New("delete file failed")
}
