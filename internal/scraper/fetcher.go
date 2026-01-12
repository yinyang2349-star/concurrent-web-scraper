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
	Fetch(ctx context.Context, url string) (string, error)
}

// HTTPFetcher implements the Fetcher interface using HTTP.
type HTTPFetcher struct {
	Client *http.Client
}

// NewHTTPFetcher creates a new instance of HTTPFetcher with a timeout.
func NewHTTPFetcher() *HTTPFetcher {
	return &HTTPFetcher{
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

// Fetch retrieves the content of the given URL.
// Update fetch method to use custom errors
func (f *HTTPFetcher) Fetch(ctx context.Context, url string) (string, error) {
	if url == "" {
		return "", ErrInvalidURL
	}

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Execute the request
	resp, err := f.Client.Do(req)
	if err != nil {
		// Check if context was canceled or timed out
		if errors.Is(err, context.Canceled) {
			return "", fmt.Errorf("request canceled: %w", err)
		}

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

	// Previous implementation without context handling

	// resp, err := f.Client.Get(url)
	// if err != nil {
	// 	// Check if it's a timeout
	// 	if errors.Is(err, context.DeadlineExceeded) {
	// 		return "", ErrTimeout
	// 	}
	// 	return "", fmt.Errorf("failed to fetch %s: %w", url, err)
	// }

	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	return "", NewFetchError(url, resp.StatusCode, fmt.Errorf("bad status code"))
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to read body for %s: %w", url, err)
	// }

	// return string(body), nil
}
