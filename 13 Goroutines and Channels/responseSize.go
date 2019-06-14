package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSIze("https://google.com")
	time.Sleep(time.Second)
	go responseSIze("https://yahoo.com")
	time.Sleep(time.Second)
	go responseSIze("https://bing.com")
	time.Sleep(time.Second)
	go responseSIze("https://yandex.com")
	time.Sleep(time.Second)
	go responseSIze("https://duckduckgo.com")
	time.Sleep(time.Second)
}

func responseSIze(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(len(body))
}
