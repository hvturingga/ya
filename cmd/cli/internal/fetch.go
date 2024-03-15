package internal

import (
	"crypto/md5"
	"fmt"
	"github.com/hvturingga/ya/conf"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SubscribeFetch(url string) string {
	name := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	path := filepath.Join(conf.GetSubscribePath(), fmt.Sprintf("%s.json", name))

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	done := make(chan bool)
	var progressBarLength int
	go func() {
		loadingChars := []string{"|", "/", "-", "\\"}
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				progressBar := fmt.Sprintf("\rDownloading %s", loadingChars[i%len(loadingChars)])
				fmt.Print(progressBar)
				progressBarLength = len(progressBar)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	r, err := client.Get(url)
	if err != nil {
		fmt.Printf("Failed to download file: %v\n", err)
		close(done)
		os.Exit(1)
	}
	defer r.Body.Close()

	close(done)
	fmt.Printf("\r%s\r", strings.Repeat(" ", progressBarLength))

	if r.StatusCode != http.StatusOK {
		fmt.Printf("Server returned %d status\n", r.StatusCode)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(body))
		os.Exit(1)
	}

	if _, err = io.Copy(f, r.Body); err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Downloaded to: ", path)
	return path
}
