package main

import (
	"fmt"
	"log"
	"net/http"
)

// homePage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequest()
}
