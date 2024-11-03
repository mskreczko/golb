package main

import (
	"log"
	"net/http"
)

func main() {
	var c Config
	c.GetConfig("config.yaml")

	http.HandleFunc("/", proxyHandler)

	log.Printf("Listening on: %s\n", c.GetFullAddress())
	if err := http.ListenAndServe(c.GetFullAddress(), nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %s %s %s", r.RemoteAddr, r.Method, r.URL)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
