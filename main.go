package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// APIResponse contains the body and status code of the API response.
type APIResponse struct {
	// Data is the body of the API response.
	Data string

	// StatusCode is the HTTP status code of the API response.
	StatusCode int
}

// fetchAPI makes a GET request to the specified URL and returns the response body and status code.
//
// The function uses a context to limit the request time and cancel waiting after the specified timeout.
// If an error occurs during the request, it returns nil and the error.
// If the timeout is exceeded, it returns nil and context.DeadlineExceeded.
//
// Parameters:
//   ctx:      The context to use for the request. Used for cancellation and timeout.
//   url:      The URL to make the GET request to.
//   timeout:  The maximum time allowed for the request, including establishing a connection, sending the request, and receiving the response.
//
// Returns:
//   *APIResponse: A structure containing the response body and status code if the request is successful.
//   error:        An error value indicating any issues encountered during the request.
//
// Example:
//   ctx := context.Background()
//   url := "https://example.com"
//   timeout := 5 * time.Second
//   apiResponse, err := fetchAPI(ctx, url, timeout)
//   if err != nil {
//       fmt.Printf("Error: %v\n", err)
//       return
//   }
//   fmt.Printf("StatusCode: %d\n", apiResponse.StatusCode)
//   fmt.Printf("Data: %s\n", apiResponse.Data)
func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Use WithTimeout to limit the request time
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Attach the context to the request
	req = req.WithContext(ctx)

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Return the APIResponse structure with the response data and status code
	return &APIResponse{
		Data:       string(body),
		StatusCode: resp.StatusCode,
	}, nil
}

func main() {
	// Example usage of fetchAPI function
	ctx := context.Background()
	url := "https://example.com"
	timeout := 5 * time.Second

	apiResponse, err := fetchAPI(ctx, url, timeout)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("StatusCode: %d\n", apiResponse.StatusCode)
	fmt.Printf("Data: %s\n", apiResponse.Data)
}

