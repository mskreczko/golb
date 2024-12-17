package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func NewServer(ctx context.Context, httpClient *http.Client, sp *ServerPool) http.Handler {
	mux := http.NewServeMux()
	var handler http.Handler = mux
	handler = LogHttp(handler)
	handler = CORS(handler)
	handler = PrometheusMiddleware(handler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyHandler(w, r, ctx, httpClient, sp)
	})
	mux.Handle("/metrics", promhttp.Handler())

	return handler
}
