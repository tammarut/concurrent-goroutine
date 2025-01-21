package solution6

import (
	"context"
	"fmt"

	"github.com/useinsider/go-pkg/insrequester"
	"golang.org/x/sync/semaphore"
)

func LimitGoroutinesWithSemaphore() {
	requester := insrequester.NewRequester().Load()
	urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	maxConcurrency := 2 // Limit number of concurrent requests

	sem := semaphore.NewWeighted(int64(maxConcurrency))
	ctx := context.Background()

	for _, url := range urls {
		// Acquire a semaphore weight before starting a goroutine
		sem.Acquire(ctx, 1) // Acquire a token

		go func(targetUrl string) {
			defer sem.Release(1) // Release the semaphore weight on completion
			response, _ := requester.Get(insrequester.RequestEntity{Endpoint: targetUrl})
			fmt.Printf("%s %d\n", url, response.StatusCode)
		}(url)
	}

	// Wait for all goroutines to release their semaphore weights
	if err := sem.Acquire(ctx, int64(maxConcurrency)); err != nil {
		fmt.Printf("Failed to acquire semaphore while waiting: %v\n", err)
	}
}
