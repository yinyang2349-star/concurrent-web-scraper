package scraper

import "time"

// Result represents the result of a fetch operation
type Result struct {
	URL       string
	Content   string
	Error     error
	Duration  time.Duration
	Timestamp time.Time
}

// NewResult creates a new Result instance
func NewResult(url, content string, err error, duration time.Duration) *Result {
	return &Result{
		URL:       url,
		Content:   content,
		Error:     err,
		Duration:  duration,
		Timestamp: time.Now(),
	}
}

// Success returns true if the fetch was successful
func (r *Result) Success() bool {
	return r.Error == nil
}
