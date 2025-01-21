package solution7

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func UseLibraryErrgroup() {
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
