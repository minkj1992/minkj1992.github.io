package urlchecker

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFailed = errors.New("Request Failed")
)

type result struct {
	url    string
	status string
}

func hitURL(url string, c chan<- result) {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}

// HitURL hits urls
func HitURL(urls []string, results map[string]string) {
	c := make(chan result)
	n := len(urls)

	for _, url := range urls {
		go hitURL(url, c)
	}

	// wait in here
	for i := 0; i < n; i++ {
		r := <-c
		results[r.url] = r.status
	}
}
