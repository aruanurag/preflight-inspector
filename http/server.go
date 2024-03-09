package http

import "net/http"

func NewServer(port string) *http.Server {

	//routers
	mux := http.NewServeMux()
	mux.Handle("/healthz", healthz())
	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
