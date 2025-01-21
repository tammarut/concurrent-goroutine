package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/useinsider/go-pkg/insrequester"
	"golang.org/x/sync/errgroup"
)

func main() {
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 1: Basic Goroutines
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	requester := insrequester.NewRequester().Load()

	urls := []string{"http://example.com", "http://example.org", "http://example.net"}
	for _, url := range urls {
		fmt.Println("Requesting", url)
		go requester.Get(insrequester.RequestEntity{Endpoint: url})
	}
	time.Sleep(2 * time.Second) // Wait for goroutines to finish

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 2: WaitGroups
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// requester := insrequester.NewRequester().Load()

	// var waitGroup sync.WaitGroup

	// urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	// waitGroup.Add(len(urls))

	// for index, url := range urls {
	// 	fmt.Printf("%d. Requesting: %s\n", index, url)
	// 	go func(targetUrl string) {
	// 		defer waitGroup.Done()
	// 		requester.Get(insrequester.RequestEntity{Endpoint: targetUrl})
	// 	}(url)
	// }

	// waitGroup.Wait() // Wait for goroutines to finish

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 3: Channels
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// requester := insrequester.NewRequester().Load()

	// urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	// channel := make(chan string, len(urls))

	// for index, url := range urls {
	// 	fmt.Printf("%d. Requesting: %s\n", index, url)
	// 	go func(targetUrl string) {
	// 		response, _ := requester.Get(insrequester.RequestEntity{Endpoint: targetUrl})
	// 		channel <- fmt.Sprintf("%s %d", targetUrl, response.StatusCode)
	// 	}(url)
	// }

	// for range urls {
	// 	response := <-channel
	// 	fmt.Println(response)
	// }

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 4: Worker Pool
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// requester := insrequester.NewRequester().Load()

	// urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	// const numberOfWorkers = 2 // Number of workers

	// jobChannel := make(chan Job, len(urls))
	// responseChannel := make(chan *http.Response, len(urls))

	// var waitGroup sync.WaitGroup

	// // Start workers
	// for currentWorker := 0; currentWorker < numberOfWorkers; currentWorker++ {
	// 	go workerDo(requester, jobChannel, responseChannel, &waitGroup)
	// }

	// // Sending jobs to thr worker pool
	// waitGroup.Add(len(urls))
	// for _, url := range urls {
	// 	jobChannel <- Job{URL: url} // send job via channel
	// }
	// close(jobChannel)
	// waitGroup.Wait()

	// // Collecting responses
	// for _, url := range urls {
	// 	response := <-responseChannel
	// 	fmt.Printf("%s %d\n", url, response.StatusCode)
	// }

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 5: Limiting Goroutines with channels
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// requester := insrequester.NewRequester().Load()
	// urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	// maxConcurrency := 2 // Limit number of concurrent requests

	// limiterChannel := make(chan struct{}, maxConcurrency)

	// for _, url := range urls {
	// 	limiterChannel <- struct{}{} // Acquire a token, waits here for token releases from the limiterChannel
	// 	go func(targetUrl string) {
	// 		defer func() { <-limiterChannel }() // Release a token
	// 		response, _ := requester.Get(insrequester.RequestEntity{Endpoint: targetUrl})
	// 		fmt.Printf("%s %d\n", url, response.StatusCode)
	// 	}(url)
	// }

	// // Wait for all goroutines to finish
	// for i := 0; i < cap(limiterChannel); i++ {
	// 	limiterChannel <- struct{}{}
	// }

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 6: Limiting Goroutines with sync/semaphore
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// requester := insrequester.NewRequester().Load()
	// urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	// maxConcurrency := 2 // Limit number of concurrent requests

	// sem := semaphore.NewWeighted(int64(maxConcurrency))
	// ctx := context.Background()

	// for _, url := range urls {
	// 	// Acquire a semaphore weight before starting a goroutine
	// 	sem.Acquire(ctx, 1) // Acquire a token

	// 	go func(targetUrl string) {
	// 		defer sem.Release(1) // Release the semaphore weight on completion
	// 		response, _ := requester.Get(insrequester.RequestEntity{Endpoint: targetUrl})
	// 		fmt.Printf("%s %d\n", url, response.StatusCode)
	// 	}(url)
	// }

	// // Wait for all goroutines to release their semaphore weights
	// if err := sem.Acquire(ctx, int64(maxConcurrency)); err != nil {
	// 	fmt.Printf("Failed to acquire semaphore while waiting: %v\n", err)
	// }

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 7: Library errgroup
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	g, ctx := errgroup.WithContext(context.Background())

	myUrls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	for _, url := range myUrls {
		// Lanch a goroutine for each URL
		g.Go(func() error {
			// Replace with actual HTTP request logic
			result, err := fetch(ctx, url)
			fmt.Println(result)
			return err
		})
	}

	// Wait for all requests to finish
	if err := g.Wait(); err != nil {
		log.Printf("Error occurred: %v", err)
	}
}

// Job is a struct for Solution 4
// type Job struct {
// 	URL string
// }

// func workerDo(requester *insrequester.Request, jobs <-chan Job, resultChannel chan<- *http.Response, wg *sync.WaitGroup) {
// 	for job := range jobs {
// 		response, _ := requester.Get(insrequester.RequestEntity{Endpoint: job.URL})
// 		resultChannel <- response
// 		wg.Done()
// 	}
// }

// fetch is a function for Solution 7
func fetch(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result := fmt.Sprintf("%s %d", url, resp.StatusCode)
	return result, nil
}
