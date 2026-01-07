package scraper

import (
	"errors"
	"fmt"
)

// Sentinel errors (known error types)
var (
	ErrInvalidURL  = errors.New("invalid URL")
	ErrTimeout     = errors.New("request timed out")
	ErrRateLimited = errors.New("rate limited")
)

// FetchError is a custom error type with extra context
type FetchError struct {
	URL        string
	StatusCode int
	Err        error
}

// Error implements the error interface
func (e *FetchError) Error() string {
	return fmt.Sprintf("fetch error for URL %s: status code %d: %v", e.URL, e.StatusCode, e.Err)
}

// Unwrap allows errors.Is and errors.As to work with FetchError
func (e *FetchError) Unwrap() error {
	return e.Err
}

// NewFetchError creates a new FetchError instance
func NewFetchError(url string, statusCode int, err error) *FetchError {
	return &FetchError{
		URL:        url,
		StatusCode: statusCode,
		Err:        err,
	}
}
