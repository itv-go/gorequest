package gorequest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Get makes a GET request to the specified URL, adds optional headers, and decodes the JSON response into the provided object.
// Usage:
//
//	var result YourType
//	if _, err := Get("https://api.example.com/data", &result, nil, 10*time.Second); err != nil {
//	    log.Println(err)
//	}
func Get[T any](url string, obj *T, headers map[string]string, timeout time.Duration) (*T, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		if cerr := Body.Close(); cerr != nil {
			fmt.Printf("Error closing response body: %s\n", cerr)
		}
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("invalid response: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, obj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return obj, nil
}
