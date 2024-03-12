package main

import (
	"github.com/aruanurag/preflight-inspector/http"
	log "k8s.io/klog/v2"
)

func main() {
	server := http.NewServer("8081")
	err := server.ListenAndServe()
	log.Fatal(err)
	log.Infof("Starting server on port: %s", 8081)
	if err := server.ListenAndServe(); err != nil {
		log.Errorf("Failed to listen and serve: %v", err)
	}
}
