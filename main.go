package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := ":8000"

	log.Println("Hello, Docker!")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Starting server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, `Hello, visitor!`)
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}
