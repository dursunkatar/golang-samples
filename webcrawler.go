package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.google.com/"

	urls := make(chan string)
	finish := make(chan bool)
	go crawl(url, urls, finish)

	for {
		select {
		case bu := <-urls:
			fmt.Println(bu)
		case <-finish:
			fmt.Println("Bitti")
			break
		}
	}

	close(urls)
}

func crawl(url string, urls chan string, finish chan bool) {
	res, _ := http.Get(url)
	defer res.Body.Close()
	z := html.NewTokenizer(res.Body)
	for {
		tt := z.Next()
		switch {
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				for _, a := range t.Attr {
					if a.Key == "href" {
						urls <- a.Val
						break
					}
				}
			}
		}
	}
	finish <- true
}
