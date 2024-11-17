package main

import (
	"net/http"
	"sync"
	"time"
)

type Server struct {
	addr                string
	healthcheckEndpoint string
	alive               bool
}

type ServerPool struct {
	servers             []*Server
	availableServers    []*Server
	currentIdx          int
	mu                  sync.Mutex
	healthcheckInterval time.Duration
}

func Init(servers []ServerConfig, healthcheckInterval int) *ServerPool {
	p := &ServerPool{
		servers:             []*Server{},
		availableServers:    []*Server{},
		currentIdx:          0,
		mu:                  sync.Mutex{},
		healthcheckInterval: time.Duration(healthcheckInterval),
	}

	for _, srv := range servers {
		health := p.healthcheck(srv.Addr, srv.HealthcheckEndpoint)
		server := &Server{
			addr:                srv.Addr,
			healthcheckEndpoint: srv.HealthcheckEndpoint,
			alive:               health,
		}
		p.servers = append(p.servers, server)
		if health {
			p.availableServers = append(p.availableServers, server)
		}
	}
	return p
}

func (p *ServerPool) removeFromPool(server Server) {
	for i, srv := range p.availableServers {
		if server.addr == srv.addr {
			p.availableServers = append(p.availableServers[:i], p.availableServers[i+1:]...)
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

func (p *ServerPool) healthcheck(addr string, healthcheckEndpoint string) bool {
	_, err := http.Get(addr + healthcheckEndpoint)
	if err != nil {
		return false
	}
	return true
}

func (p *ServerPool) HealthcheckAll() {
	for {
		for idx := range p.servers {
			_, err := http.Get(p.servers[idx].addr + p.servers[idx].healthcheckEndpoint)
			if err != nil {
				p.removeFromPool(*p.servers[idx])
			} else {
				p.addToPool(*p.servers[idx])
			}
		}
		time.Sleep(p.healthcheckInterval * time.Second)
	}
}

func (p *ServerPool) getNext() *Server {
	if len(p.availableServers) == 0 {
		return nil
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	srv := p.availableServers[p.currentIdx%len(p.availableServers)]
	p.currentIdx++
	return srv
}
