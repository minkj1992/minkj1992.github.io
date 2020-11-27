package main

import (
	"fmt"

	"github.com/minkj1992/go_nomad/urlchecker/urlchecker"
)

func main() {
	results := make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	urlchecker.HitURL(urls, results)
	for url, result := range results {
		fmt.Println(url, result)
	}
}
