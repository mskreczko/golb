package main

import (
	"log"
	"net/http"
)

func main() {
	var c Config
	c.GetConfig("config.yaml")

	log.Printf("Listening on: %s\n", c.GetFullAddress())
	err := http.ListenAndServe(c.GetFullAddress(), nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
