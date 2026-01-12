package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yinyang2349-star/concurrent-web-scraper/internal/scraper"
)

func main() {
	fmt.Println("Starting web scraper...")
	fmt.Println("🚀 Day 4 - Context Package\n")

	fetcher := scraper.NewHTTPFetcher()

	// Test 1: Normal fetch with context
	fmt.Println("Test 1: Normal fetch with context")
	ctx := context.Background()
	content, err := fetcher.Fetch(ctx, "https://example.com")
	if err != nil {
		log.Printf("❌ error fetching URL: %v\n", err)
	} else {
		fmt.Printf("✅ successfully fetched content length: %d bytes\n\n", len(content))
	}

	// Test 2: Fetch with timeout context
	fmt.Println("Test 2: Fetch with 1 second timeout context")
	ctx2, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// This might take longer than 1 second depending on network conditions
	content, err = fetcher.Fetch(ctx2, "https://httpstat.us/200? sleep=2000")
	if err != nil {
		log.Printf("❌ Timeout as expected: %v\n", err)
	} else {
		fmt.Printf("✅ successfully fetched content length: %d bytes\n\n", len(content))
	}

	// Test 3: Manual cancellation
	fmt.Println("Test 3: Manual cancellation")
	ctx3, cancel3 := context.WithCancel(context.Background())

	// Cancel after 100 milliseconds
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("⏳ Canceling the request...")
		cancel3()
	}()

	content, err = fetcher.Fetch(ctx3, "https://httpstat.us/200? sleep=5000")
	if err != nil {
		log.Printf("❌ Request canceled as expected: %v\n", err)
	} else {
		fmt.Printf("✅ successfully fetched content length: %d bytes\n\n", len(content))
	}
}
