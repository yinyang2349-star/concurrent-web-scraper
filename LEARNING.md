# Learning Notes - Week 1

## Day 2: Interfaces

### Key Concepts:

- Interfaces are satisfied implicitly (no "implements" keyword)
- Accept interfaces, return concrete types
- Empty interface `interface{}` can hold any value
- Type assertion: `value.(Type)`

### Questions:

- Q: When to use interfaces?
- A: When you want flexibility and testability

### Code Examples:

```go
package scraper

import (
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
func (hf *HTTPFetcher) Fetch(url string) (string, error) {
	resp, err := hf.Client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 response for URL %s: %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body for URL %s: %w", url, err)
	}
	return string(body), nil
}

```

## Day 3: Error Handling

### Key Concepts:

- Errors are values (implements `error` interface)
- `fmt.Errorf` with `%w` for error wrapping
- `errors.Is()` for error comparison
- `errors.As()` for type assertion
- Custom error types for rich error information

### Best Practices:

- Always wrap errors with context
- Use sentinel errors for known error types
- Custom errors when you need extra data
