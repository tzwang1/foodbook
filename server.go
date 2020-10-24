package main

import (
	"fmt"
	"net/http"
	"log"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to example server!")
}

func main() {
	http.HandleFunc("/", exampleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
