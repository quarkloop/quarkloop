package api

import (
	"net/url"

	"github.com/quarkloop/quarkloop/pkg/db"
)

func GetFileById(appId, instanceId, fileId string) (interface{}, error) {
	params := url.Values{}
	params.Add("id", fileId)

	res, err := db.HttpClientInstance.Get("http://localhost:3000/api/v1/tables/appFile", &params)
	if err != nil {
		return nil, err
	}

	return res, nil
}
