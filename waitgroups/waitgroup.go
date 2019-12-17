package main

import (
	"fmt"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	var urls = []string{
		"http://www.golang.org/",
		"https://www.facebook.com",
		"https://www.google.com",
	}

	for _, url := range urls {
		go func(url string) {
			http.Get(url)
			fmt.Println("Fetched")
		}(url)
	}
}
