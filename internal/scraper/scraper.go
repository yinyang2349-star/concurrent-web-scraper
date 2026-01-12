package scraper

import (
	"context"
	"time"
)

// Scraper orchestrates fetching multiple URLs
type Scraper struct {
	fetcher Fetcher
}

// NewScraper creates a new scraper instance
func NewScraper(fetcher Fetcher) *Scraper {
	return &Scraper{
		fetcher: fetcher,
	}
}

// Scrape fetches multiple URLs and returns results
func (s *Scraper) Scrape(ctx context.Context, urls []string) []*Result {
	results := make([]*Result, 0, len(urls))

	for _, url := range urls {
		start := time.Now()
		content, err := s.fetcher.Fetch(ctx, url)
		duration := time.Since(start)

		result := NewResult(url, content, err, duration)
		results = append(results, result)
	}

	return results
}
