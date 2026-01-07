package scraper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Fetcher defines the interface for fetching web content.
type Fetcher interface {
	Fetch(url string) (string, error)
}

// HTTPFetcher implements the Fetcher interface using HTTP.
type HTTPFetcher struct {
	Client *http.Client
}

// NewHTTPFetcher creates a new instance of HTTPFetcher with a timeout.
func NewHTTPFetcher(timeout time.Duration) *HTTPFetcher {
	return &HTTPFetcher{
		Client: &http.Client{Timeout: timeout},
	}
}

// Fetch retrieves the content of the given URL.
// Update fetch method to use custom errors
func (f *HTTPFetcher) Fetch(url string) (string, error) {
	if url == "" {
		return "", ErrInvalidURL
	}

	resp, err := f.Client.Get(url)
	if err != nil {
		// Check if it's a timeout
		if errors.Is(err, context.DeadlineExceeded) {
			return "", ErrTimeout
		}
		return "", fmt.Errorf("failed to fetch %s: %w", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", NewFetchError(url, resp.StatusCode, fmt.Errorf("bad status code"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body for %s: %w", url, err)
	}

	return string(body), nil
}
