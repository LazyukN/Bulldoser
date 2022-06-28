package main

import (
	"fmt"
	"log"
	"net/http"
)

var requestCounter int

func main() {
	http.HandleFunc("/ping/", ping)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
	requestCounter++
	fmt.Println(requestCounter)
	fmt.Fprintf(w, "ping\n")
}
