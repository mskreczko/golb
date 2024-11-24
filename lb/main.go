package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var c Config
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()
	c.GetConfig(*configPath)

	httpClient := &http.Client{}
	sp := Init(c.Servers, c.HealthcheckInterval)
	ctx := context.Background()

	srv := NewServer(ctx, httpClient, sp)
	httpServer := &http.Server{
		Addr:    c.GetFullAddress(),
		Handler: srv,
	}

	go sp.HealthcheckAll()

	go func() {
		slog.Info("Listening on: ", "", c.GetFullAddress())
		if err := httpServer.ListenAndServe(); err != nil {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
}

func proxyHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, httpClient *http.Client, sp *ServerPool) {
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
