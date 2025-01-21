package solution2

import (
	"fmt"
	"sync"

	"github.com/useinsider/go-pkg/insrequester"
)

func WaitGroups() {
	requester := insrequester.NewRequester().Load()

	var waitGroup sync.WaitGroup

	urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	waitGroup.Add(len(urls))

	for index, url := range urls {
		fmt.Printf("%d. Requesting: %s\n", index, url)
		go func(targetUrl string) {
			defer waitGroup.Done()
			requester.Get(insrequester.RequestEntity{Endpoint: targetUrl})
		}(url)
	}

	waitGroup.Wait() // Wait for goroutines to finish
}
