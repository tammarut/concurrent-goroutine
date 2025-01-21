package solution4

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/useinsider/go-pkg/insrequester"
)

func WorkerPool() {
	requester := insrequester.NewRequester().Load()

	urls := []string{"https://api.jikan.moe/v4/random/users", "https://api.jikan.moe/v4/random/anime", "https://api.jikan.moe/v4/random/manga"}
	const numberOfWorkers = 2 // Number of workers

	jobChannel := make(chan Job, len(urls))
	responseChannel := make(chan *http.Response, len(urls))

	var waitGroup sync.WaitGroup

	// Start workers
	for currentWorker := 0; currentWorker < numberOfWorkers; currentWorker++ {
		go workerDo(requester, jobChannel, responseChannel, &waitGroup)
	}

	// Sending jobs to thr worker pool
	waitGroup.Add(len(urls))
	for _, url := range urls {
		jobChannel <- Job{URL: url} // send job via channel
	}
	close(jobChannel)
	waitGroup.Wait()

	// Collecting responses
	for _, url := range urls {
		response := <-responseChannel
		fmt.Printf("%s %d\n", url, response.StatusCode)
	}
}

// Job is a struct for Solution 4
type Job struct {
	URL string
}

func workerDo(requester *insrequester.Request, jobs <-chan Job, resultChannel chan<- *http.Response, wg *sync.WaitGroup) {
	for job := range jobs {
		response, _ := requester.Get(insrequester.RequestEntity{Endpoint: job.URL})
		resultChannel <- response
		wg.Done()
	}
}
