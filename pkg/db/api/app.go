package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/quarkloop/quarkloop/pkg/db/client"
	"github.com/quarkloop/quarkloop/pkg/db/model"
)

func CreateApp(app *model.App) (*model.App, error) {
	payload, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}

	res, err := client.DatabaseClient.Post("http://localhost:3000/api/v1/tables/app", nil, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		var payload DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return nil, err
		}

		var app model.App
		if err := json.Unmarshal(payload.Database.App.Records, &app); err != nil {
			return nil, err
		}

		return &app, nil
	}

	return nil, errors.New("failed to create app")
}

func UpdateApp(app *model.App) (*model.App, error) {
	payload, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}

	res, err := client.DatabaseClient.Update("http://localhost:3000/api/v1/tables/app", nil, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		var payload DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return nil, err
		}

		var app model.App
		if err := json.Unmarshal(payload.Database.App.Records, &app); err != nil {
			return nil, err
		}

		return &app, nil
	}

	return nil, errors.New("failed to update app")
}

func DeleteApp(appId string) error {
	q := url.Values{}
	q.Add("id", appId)

	res, err := client.DatabaseClient.Delete("http://localhost:3000/api/v1/tables/app", &q)
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

	return errors.New("delete app failed")
}

func CreateAppInstance(instance *model.AppInstance) (*model.AppInstance, error) {
	payload, err := json.Marshal(instance)
	if err != nil {
		return nil, err
	}

	res, err := client.DatabaseClient.Post("http://localhost:3000/api/v1/tables/appInstance", nil, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		var payload DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return nil, err
		}

		var instance model.AppInstance
		if err := json.Unmarshal(payload.Database.AppInstance.Records, &instance); err != nil {
			return nil, err
		}

		return &instance, nil
	}

	return nil, errors.New("failed to create app instance")
}

func UpdateAppInstance(instance *model.AppInstance) (*model.AppInstance, error) {
	payload, err := json.Marshal(instance)
	if err != nil {
		return nil, err
	}

	res, err := client.DatabaseClient.Update("http://localhost:3000/api/v1/tables/appInstance", nil, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		var payload DatabaseResponsePayload
		if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
			return nil, err
		}

		var instance model.AppInstance
		if err := json.Unmarshal(payload.Database.AppInstance.Records, &instance); err != nil {
			return nil, err
		}

		return &instance, nil
	}

	return nil, errors.New("failed to update app instance")
}
