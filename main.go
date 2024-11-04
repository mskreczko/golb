package main

import (
	"log"
	"net/http"
)

func main() {
	var c Config
	c.GetConfig("config.yaml")

	var sp ServerPool
	sp.Init(c.Servers)

	http.HandleFunc("/", proxyHandler)

	log.Printf("Listening on: %s\n", c.GetFullAddress())
	if err := http.ListenAndServe(c.GetFullAddress(), nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %s %s %s", r.RemoteAddr, r.Method, r.URL)
}
