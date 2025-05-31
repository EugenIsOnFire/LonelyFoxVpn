package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, anonymous from %s", request.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Запуск сервера")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
