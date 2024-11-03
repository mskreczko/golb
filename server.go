package main

import (
	"net/http"
	"sync"
)

type Server struct {
	addr                string
	healthcheckEndpoint string
	alive               bool
}

type ServerPool struct {
	servers          []*Server
	availableServers []*Server
	currentIdx       int
	mu               sync.Mutex
}

func (p *ServerPool) removeFromPool(server Server) {
	for i, srv := range p.availableServers {
		if server.addr == srv.addr {
			p.servers = append(p.servers[:i], p.servers[i+1:]...)
		}
	}
}

func (p *ServerPool) addToPool(server Server) {
	for _, srv := range p.availableServers {
		if server.addr == srv.addr {
			return
		}
	}
	p.availableServers = append(p.availableServers, &server)
}

func (p *ServerPool) healthcheck() {
	for idx := range p.availableServers {
		_, err := http.Get(p.servers[idx].healthcheckEndpoint)
		if err != nil {
			p.removeFromPool(*p.servers[idx])
		} else {
			p.addToPool(*p.servers[idx])
		}
	}
}

func (p *ServerPool) getNext() *Server {
	p.mu.Lock()
	defer p.mu.Unlock()
	srv := p.availableServers[p.currentIdx%len(p.availableServers)]
	p.currentIdx++
	return srv
}
