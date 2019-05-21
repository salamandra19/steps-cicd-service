package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", steps)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func steps(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good Luck")
}
