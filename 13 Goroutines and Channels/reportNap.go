package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println("Sleeping", name)
		time.Sleep(time.Second)
	}
	fmt.Println("Waking up", name)
}

func send(myChannel chan string) {
	reportNap("Sending goroutine", 2)
	fmt.Println("***Sending value***")
	myChannel <- "a"
	fmt.Println("***Sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("Receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
