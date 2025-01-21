package solution5

import (
	"fmt"

	"github.com/useinsider/go-pkg/insrequester"
)

func LimitGoroutinesWithChannel() {
	requester := insrequester.NewRequester().Load()
	urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	maxConcurrency := 2 // Limit number of concurrent requests

	limiterChannel := make(chan struct{}, maxConcurrency)

	for _, url := range urls {
		limiterChannel <- struct{}{} // Acquire a token, waits here for token releases from the limiterChannel
		go func(targetUrl string) {
			defer func() { <-limiterChannel }() // Release a token
			response, _ := requester.Get(insrequester.RequestEntity{Endpoint: targetUrl})
			fmt.Printf("%s %d\n", url, response.StatusCode)
		}(url)
	}

	// Wait for all goroutines to finish
	for i := 0; i < cap(limiterChannel); i++ {
		limiterChannel <- struct{}{}
	}
}
