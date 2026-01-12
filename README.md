# 🚀 Concurrent Web Scraper

A production-ready web scraper built with Go, focusing on clean architecture, proper error handling, and context management.

## 📚 Learning Project

This is Week 1 of my journey to become an expert software engineer in 2026. The goal is to master Go fundamentals, interfaces, error handling, and the context package.

## ✨ Features

- ✅ Interface-based design for testability
- ✅ Custom error types with error wrapping
- ✅ Context support for cancellationn and timeouts
- ✅ Multiple URL scraping with result tracking
- ✅ Duration and timestamp tracking
- 🔄 Concurrency support (coming in Week 2)

## 🏗️ Architecture

```
concurrent-web-scraper/
├── cmd/
│   └── scraper/
│       └── main.go          # Entry point
├── internal/
│   └── scraper/
│       ├── fetcher.go       # HTTP fetcher implementation
│       ├── errors.go        # Custom error types
│       ├── result.go        # Result data structure
│       └── scraper.go       # Scraper orchestrator
├── pkg/                     # (future:  reusable packages)
├── test/                    # (future: tests)
├── go.mod
├── go.sum
├── README.md
└── LEARNING.md
```

## 🚀 Usage

```bash
# Clone repository
git clone https://github.com/yinyang2349-star/concurrent-web-scraper.git
cd concurrent-web-scraper

# Run scraper
go run cmd/scraper/main.go
```

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/yinyang2349-star/concurrent-web-scraper/internal/scraper"
)

func main() {
	fmt.Println("Starting web scraper...")
	fmt.Println("🚀 Day 5 - Multiple URLs")

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

```

## 🛠️ Tech Stack

- **Language:** Go 1.21+
- **Standard Library:** net/http, context, errors, io

## 📈 Roadmap

- [x] Week 1: Fundamentals (interfaces, errors, context)
- [ ] Week 2: Concurrency (goroutines, channels, worker pools)
- [ ] Week 3: Testing (unit tests, table-driven tests, mocks)
- [ ] Week 4: CLI & Features (Cobra, rate limiting, output formats)

## 🎓 Learning Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com)
- [Go Blog - Error Handling](https://go.dev/blog/error-handling-and-go)
- [Go Blog - Context](https://go.dev/blog/context)

## 📝 Development Log

See [LEARNING.md](LEARNING.md) for daily learning notes and progress.

## 👤 Author

**yinyang2349-star**

- GitHub: [@yinyang2349-star](https://github.com/yinyang2349-star)
- Learning Journey: [2026 Expert Software Engineer Roadmap]

## 📄 License

MIT License - feel free to use for learning!

---

**⭐ Star this repo if you're also on a learning journey!**
