package solution1

import (
	"fmt"
	"time"

	"github.com/useinsider/go-pkg/insrequester"
)

func SimpleGoroutine() {
	requester := insrequester.NewRequester().Load()

	urls := []string{"http://example.com", "http://example.org", "http://example.net"}
	for _, url := range urls {
		fmt.Println("Requesting", url)
		go requester.Get(insrequester.RequestEntity{Endpoint: url})
	}
	time.Sleep(2 * time.Second) // Wait for goroutines to finish
}
