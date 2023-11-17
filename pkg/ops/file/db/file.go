package db

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/quarkloop/quarkloop/pkg/db/api"
	"github.com/quarkloop/quarkloop/pkg/db/client"
	"github.com/quarkloop/quarkloop/pkg/ops/file/model"
)

func GetFileById(projectId, instanceId, fileId string) (*model.File, error) {
	params := url.Values{}
	params.Add("projectId", projectId)
	params.Add("instanceId", instanceId)
	params.Add("fileId", fileId)

	res, err := client.DatabaseClient.Get("http://localhost:3000/api/v1/tables/appFile", &params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var payload api.DatabaseResponsePayload
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

func CreateFile(projectId, instanceId string, file *model.File) (*model.File, error) {
	params := url.Values{}
	params.Add("projectId", projectId)
	params.Add("instanceId", instanceId)

	payload, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}

	res, err := client.DatabaseClient.Post("http://localhost:3000/api/v1/tables/appFile", &params, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		var payload api.DatabaseResponsePayload
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

func UpdateFile(projectId, instanceId string, file *model.File) (*model.File, error) {
	params := url.Values{}
	params.Add("projectId", projectId)
	params.Add("instanceId", instanceId)

	payload, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}

	res, err := client.DatabaseClient.Update("http://localhost:3000/api/v1/tables/appFile", &params, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		var payload api.DatabaseResponsePayload
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

func DeleteFile(projectId, instanceId, fileId string) error {
	params := url.Values{}
	params.Add("projectId", projectId)
	params.Add("instanceId", instanceId)
	params.Add("fileId", fileId)

	res, err := client.DatabaseClient.Delete("http://localhost:3000/api/v1/tables/appFile", &params)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNoContent {
		var payload api.DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return err
		}

		return nil
	}

	return errors.New("delete file failed")
}
