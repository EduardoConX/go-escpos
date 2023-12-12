package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port    = 8083
	logging = true
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
