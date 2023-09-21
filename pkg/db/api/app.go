package api

import (
	"encoding/json"
	"errors"
	"net/http"

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
