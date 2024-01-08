package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

type APIResponse struct {
	Data       string
	StatusCode int
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &APIResponse{
		Data:       string(body),
		StatusCode: resp.StatusCode,
	}, nil
}
