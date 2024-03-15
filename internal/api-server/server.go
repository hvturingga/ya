package api_server

import (
	"context"
	"fmt"
	"github.com/hvturingga/ya/conf"
	"net/http"
	"time"
)

func WaitForAPIServer() bool {
	maxAttempts := 20
	waitInterval := 2 * time.Second

	cli := &http.Client{
		Timeout: 10 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(maxAttempts)*waitInterval)
	defer cancel()

	ticker := time.NewTicker(waitInterval)
	defer ticker.Stop()

	for attempt := 1; ; attempt++ {
		select {
		case <-ctx.Done():
			return false
		case <-ticker.C:
			addr := fmt.Sprintf("http://%s", conf.ClashAPI)
			r, err := cli.Get(addr)
			if err == nil && r.StatusCode == http.StatusOK {
				r.Body.Close()
				return true
			}
			if attempt < maxAttempts {
				fmt.Printf("Connection to API server failed on attempt %d, waiting %s before retrying...\n", attempt, waitInterval)
			}
		}
	}
}
