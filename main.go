package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	var c Config
	c.GetConfig("config.yaml")

	httpClient := &http.Client{}
	sp := Init(c.Servers)
	ctx := context.Background()

	go sp.HealthcheckAll()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyHandler(w, r, ctx, httpClient, sp)
	})

	log.Printf("Listening on: %s\n", c.GetFullAddress())
	if err := http.ListenAndServe(c.GetFullAddress(), nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, httpClient *http.Client, sp *ServerPool) {
	log.Printf("Incoming request: %s %s %s", r.RemoteAddr, r.Method, r.URL)

	targetServer := sp.getNext()
	if targetServer == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	req := copyRequest(r, ctx, targetServer)
	log.Printf("Redirecting request to: %s", req.URL.String())

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return
	}

	defer resp.Body.Close()
	copyResponse(w, resp)
}
