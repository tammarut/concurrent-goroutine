package solution3

import (
	"fmt"

	"github.com/useinsider/go-pkg/insrequester"
)

func UseChannels() {
	requester := insrequester.NewRequester().Load()

	urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	channel := make(chan string, len(urls))

	for index, url := range urls {
		fmt.Printf("%d. Requesting: %s\n", index, url)
		go func(targetUrl string) {
			response, _ := requester.Get(insrequester.RequestEntity{Endpoint: targetUrl})
			channel <- fmt.Sprintf("%s %d", targetUrl, response.StatusCode)
		}(url)
	}

	for range urls {
		response := <-channel
		fmt.Println(response)
	}
}
