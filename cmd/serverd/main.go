package main

import (
	"log"

	"github.com/aruanurag/preflight-inspector/http"
)

func main() {
	server := http.NewServer("8081")
	err := server.ListenAndServe()
	log.Fatal(err)
}
