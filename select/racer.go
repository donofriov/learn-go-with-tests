package racer

import (
	"net/http"
	"time"
)

// Racer function takes in two URLs
// returns the one with the fastest response time
func Racer(firstURL, secondURL string) (winner string) {
	firstDuration := measureResponseTime(firstURL)
	secondDuration := measureResponseTime(secondURL)

	if firstDuration < secondDuration {
		return firstURL
	}

	return secondURL
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
