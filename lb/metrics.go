package main

import (
	"encoding/json"
	"net/http"
)

func (p *ServerPool) GetAllServers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(p.Servers)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
