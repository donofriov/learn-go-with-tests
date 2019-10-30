package racer

import (
	"net/http"
	"time"
)

// Racer function takes in two URLs
// returns the one with the fastest response time
func Racer(firstURL, secondURL string) (winner string) {
	startFirst := time.Now()
	http.Get(firstURL)
	firstDuration := time.Since(startFirst)

	startSecond := time.Now()
	http.Get(secondURL)
	secondDuration := time.Since(startSecond)

	if firstDuration < secondDuration {
		return firstURL
	}

	return secondURL
}
