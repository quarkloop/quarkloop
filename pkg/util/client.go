package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func NewRequest() []byte {
	requestURL := fmt.Sprintf("http://localhost:%d", 3000)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return resBody
}

type DatabaseHttpClient struct{}

var DatabaseClient = DatabaseHttpClient{}

func (c *DatabaseHttpClient) Get(path string, queryParams *url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", path, nil)
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

func (c *DatabaseHttpClient) Post(path string, queryParams *url.Values, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
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

func (c *DatabaseHttpClient) Update(path string, queryParams *url.Values, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("PUT", path, bytes.NewBuffer(body))
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

func (c *DatabaseHttpClient) Delete(path string, queryParams *url.Values) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", path, nil)
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
