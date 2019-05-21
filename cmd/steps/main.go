package main

import (
	"fmt"
	"net/http"

	"github.com/powerman/structlog"
)

func main() {
	log := structlog.New()
	http.HandleFunc("/", steps)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func steps(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good Luck")
}
