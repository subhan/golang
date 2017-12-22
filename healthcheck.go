package main

import (
	"fmt"
	"net/http"
)

func main() {

    links := []string{
        "http://google.com",
        "http://facebook.com",
        "http://news.ycombinator.com",
        "http://amazon.com",
    }

    c := make(chan string)
    for _, url := range links {
        go checkLink(url, c)
    }
    for l := range c{
        go checkLink(l, c)
    }
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if resp.StatusCode != 200 && err != nil {
		fmt.Println(err)
        c <- link
		return
	}
    fmt.Println(link, " up!")
    c <- link
}
