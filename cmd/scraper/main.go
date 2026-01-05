package main

import (
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
		"https://golang.org",
	}

	// Fetch and print content for each URL
	for _, url := range urls {
		content, err := fetcher.Fetch(url)
		if err != nil {
			log.Printf("Error fetching %s: %v", url, err)
			continue
		}
		fmt.Printf("Content fetched from %s:\n%s\n", url, content[:100]) // Print first 100 characters
	}
	fmt.Println("Web scraper finished.")
}
