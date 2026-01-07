package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/yinyang2349-star/concurrent-web-scraper/internal/scraper"
)

func main() {
	fmt.Println("Starting web scraper...")

	// Create a new HTTP fetcher with a 10-second timeout
	fetcher := scraper.NewHTTPFetcher(10 * time.Second)

	// Test URLs
	urls := []string{
		"https://example.com",
		"https://www.google.com/404",     // Will return 404
		"",                               // Invalid URL
		"https://github.com/nonexistent", // Will return 404
	}

	// Fetch and print content for each URL
	for _, url := range urls {
		fmt.Printf("📡 Fetching: %s\n", url)

		content, err := fetcher.Fetch(url)
		if err != nil {
			// Check for specific error types
			if errors.Is(err, scraper.ErrInvalidURL) {
				log.Printf("❌ Invalid URL: %s\n", url)
			} else if errors.Is(err, scraper.ErrTimeout) {
				log.Printf("❌ Timeout while fetching: %s\n", url)
			} else {
				var fetchErr *scraper.FetchError
				if errors.As(err, &fetchErr) {
					log.Printf("❌ Fetch error for %s: status code %d\n", fetchErr.URL, fetchErr.StatusCode)
				} else {
					log.Printf("❌ General error fetching %s: %v\n", url, err)
				}
			}

			continue
		}

		fmt.Printf("✅ Fetched content from %s: %d bytes\n", url, len(content))
	}

}
