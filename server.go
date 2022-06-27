package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping/", ping)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping!")
	fmt.Fprintf(w, "ping\n")
}
