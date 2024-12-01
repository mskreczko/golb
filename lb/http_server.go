package main

import (
	"context"
	"net/http"
)

func NewServer(ctx context.Context, httpClient *http.Client, sp *ServerPool) http.Handler {
	mux := http.NewServeMux()
	var handler http.Handler = mux
	handler = LogHttp(handler)
	handler = CORS(handler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyHandler(w, r, ctx, httpClient, sp)
	})
	mux.HandleFunc("/metrics/servers", func(w http.ResponseWriter, r *http.Request) {
		sp.GetAllServers(w)
	})

	return handler
}
