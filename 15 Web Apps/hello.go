package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello Web")
}

func tamilHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Vanakkam Web")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste Web")
}

func main() {
	http.HandleFunc("/tamil", tamilHandler)
	http.HandleFunc("/english", englishHandler)
	http.HandleFunc("hindi", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
