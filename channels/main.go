package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://jacobnelson.tech",
	}

	sc := make(chan string)

	for _, link := range links {
		go checkLink(link, sc)
	}

	for l := range sc {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, sc)
		}(l)
	}
}

func checkLink(link string, sc chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		sc <- link
		return
	}
	fmt.Println(link, "is up!")
	sc <- link
}
