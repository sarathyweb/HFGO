package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	pages := make(chan Page)
	urls := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://yandex.com",
		"https://duckduckgo.com",
	}

	for _, url := range urls {
		go responseSIze(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s : %d\n", page.URL, page.Size)
	}
}

func responseSIze(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	channel <- Page{URL: url, Size: len(body)}
}
