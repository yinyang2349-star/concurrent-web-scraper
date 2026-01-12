package main

import (
	"context"
	"fmt"
	"time"

	"github.com/yinyang2349-star/concurrent-web-scraper/internal/scraper"
)

func main() {
	fmt.Println("Starting web scraper...")
	fmt.Println("🚀 Day 5 - Multiple URLs\n")

	// Create scraper instance
	fetcher := scraper.NewHTTPFetcher()
	scr := scraper.NewScraper(fetcher)

	// URLs to scrape
	urls := []string{
		"https://example.com",
		"https://httpstat.us/200",
		"https://httpstat.us/404",
		"https://httpstat.us/500",
		"https://thisurldoesnotexist.tld", // Invalid URL to test error handling
	}

	// Scrape with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Printf("Scraping %d URLs with 30 seconds timeout...\n\n", len(urls))
	start := time.Now()

	results := scr.Scrape(ctx, urls)

	totalDuration := time.Since(start)

	// Print results
	successCount := 0
	for _, result := range results {
		if result.Success() {
			fmt.Printf("✅ Fetched %s in %v (Content Length: %d bytes)\n", result.URL, result.Duration, len(result.Content))
			successCount++
		} else {
			fmt.Printf("❌ Failed to fetch %s in %v (Error: %v)\n", result.URL, result.Duration, result.Error)

		}
	}

	// Summary
	fmt.Println("-------------------------------------")
	fmt.Printf("\n--- Summary ---")
	fmt.Printf("\nTotal URLs: %d", len(urls))
	fmt.Printf("\nSuccessful fetches: %d", successCount)
	fmt.Printf("\nFailed fetches: %d", len(urls)-successCount)
	fmt.Printf("\nTotal duration: %v\n", totalDuration)

}
