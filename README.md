# GoRequest: Simplified HTTP Requests in Go

## Overview

`gorequest` is a lightweight Go package that simplifies HTTP interactions by providing generic methods for making GET and POST requests. The package handles common tasks like JSON serialization/deserialization, header management, and error handling with a clean, type-safe approach.

## Features

- Generic type support for JSON requests and responses
- Simple GET and POST request methods
- Automatic JSON marshaling and unmarshaling
- Configurable request headers and timeouts
- Comprehensive error handling
- Type-safe request and response handling

## Installation

Install the package using Go modules:

```bash
go get github.com/itv-go/gorequest
```

## Usage Examples

### GET Request

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/itv-go/gorequest"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    var user User
    
    // Basic GET request
    _, err := gorequest.Get(
        "https://api.example.com/users/1", 
        &user, 
        nil,           // No additional headers
        10*time.Second // Request timeout
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("User: %+v\n", user)

    // GET request with custom headers
    headers := map[string]string{
        "Authorization": "Bearer YOUR_TOKEN",
    }
    var users []User
    _, err = gorequest.Get(
        "https://api.example.com/users", 
        &users, 
        headers, 
        15*time.Second
    )
    if err != nil {
        log.Fatal(err)
    }
}
```

### POST Request

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/itv-go/gorequest"
)

type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateUserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // Create user request
    requestData := CreateUserRequest{
        Name:  "John Doe",
        Email: "john@example.com",
    }

    var response CreateUserResponse
    
    // Basic POST request
    _, err := gorequest.Post(
        "https://api.example.com/users",     // URL
        requestData,                         // Request payload
        &response,                           // Response object
        nil,                                 // No additional headers
        10*time.Second                       // Request timeout
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Created User: %+v\n", response)

    // POST request with custom headers
    headers := map[string]string{
        "Authorization": "Bearer YOUR_TOKEN",
        "Content-Type":  "application/json",
    }
    _, err = gorequest.Post(
        "https://api.example.com/users", 
        requestData, 
        &response, 
        headers, 
        15*time.Second
    )
    if err != nil {
        log.Fatal(err)
    }
}
```

## Function Signatures

### GET Request
```go
func Get[T any](
    url string,                   // Target URL
    obj *T,                       // Pointer to result struct
    headers map[string]string,    // Optional request headers
    timeout time.Duration         // Request timeout duration
) (*T, error)
```

### POST Request
```go
func Post[R any](
    url string,                   // Target URL
    data interface{},             // Request payload
    responseObj *R,               // Pointer to response struct
    headers map[string]string,    // Optional request headers
    timeout time.Duration         // Request timeout duration
) (*R, error)
```

## Error Handling

Both methods return detailed errors for:
- Request creation failures
- JSON marshaling/unmarshaling errors
- Network request errors
- Non-2xx HTTP status codes

## Important Notes

- Supports any JSON-marshalable/unmarshalable structs
- Timeout is mandatory to prevent hanging requests
- Headers are optional
- Returns both the populated object and an error
- Automatically sets `Content-Type: application/json` for POST requests
- Validates HTTP status codes (200-299 range)
