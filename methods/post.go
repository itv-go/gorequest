package gorequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// POST request with JSON data and decodes the response into the provided object.
func Post[R any](
	url string,
	data interface{},
	responseObj *R,
	headers map[string]string,
	timeout time.Duration,
) (*R, error) {
	client := &http.Client{Timeout: timeout}

	// Marshal the data into JSON
	bodyBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Add additional headers if provided
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		if cerr := Body.Close(); cerr != nil {
			fmt.Printf("Error closing response body: %s\n", cerr)
		}
	}(resp.Body)

	// Check for HTTP status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("invalid response: status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the response into the provided response object
	if err := json.Unmarshal(body, responseObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return responseObj, nil
}
