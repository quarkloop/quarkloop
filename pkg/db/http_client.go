package db

import (
	"bytes"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
}

var HttpClientInstance = HttpClient{}

func (c *HttpClient) Get(path string, buf []byte, queryParams *url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", path, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if queryParams != nil {
		req.URL.RawQuery = queryParams.Encode()
	}

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *HttpClient) Post() (interface{}, error) {
	return nil, nil
}

func (c *HttpClient) Update() (interface{}, error) {
	return nil, nil
}

func (c *HttpClient) Delete() (interface{}, error) {
	return nil, nil
}
