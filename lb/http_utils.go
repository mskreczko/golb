package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func copyRequest(r *http.Request, ctx context.Context, targetServer *Server) *http.Request {
	req := r.Clone(ctx)
	req.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	req.URL = &url.URL{
		Scheme: "http",
		Host:   strings.Split(targetServer.Addr, "//")[1],
		Path:   url.QueryEscape(req.RequestURI),
	}
	req.RequestURI = ""
	return req
}

func copyResponse(w http.ResponseWriter, resp *http.Response) {
	w.WriteHeader(resp.StatusCode)
	if resp.Body != nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(body)
	}
}
