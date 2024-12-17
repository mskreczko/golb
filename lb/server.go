package main

import (
	"net/http"
	"sync"
	"time"
)

type Server struct {
	Addr                string   `json:"addr"`
	HealthcheckEndpoint string   `json:"healthcheckEndpoint"`
	Alive               bool     `json:"alive"`
	LastAlive           JSONTime `json:"lastAlive"`
}

type ServerPool struct {
	Servers             []*Server
	AvailableServers    []*Server
	currentIdx          int
	mu                  sync.Mutex
	healthcheckInterval time.Duration
}

func Init(servers []ServerConfig, healthcheckInterval int) *ServerPool {
	p := &ServerPool{
		Servers:             []*Server{},
		AvailableServers:    []*Server{},
		currentIdx:          0,
		mu:                  sync.Mutex{},
		healthcheckInterval: time.Duration(healthcheckInterval),
	}

	for _, srv := range servers {
		health := p.healthcheck(srv.Addr, srv.HealthcheckEndpoint)
		server := &Server{
			Addr:                srv.Addr,
			HealthcheckEndpoint: srv.HealthcheckEndpoint,
			Alive:               health,
		}
		p.Servers = append(p.Servers, server)
		if health {
			p.AvailableServers = append(p.AvailableServers, server)
		}
	}
	return p
}

func (p *ServerPool) removeFromPool(server Server) {
	for i, srv := range p.AvailableServers {
		if server.Addr == srv.Addr {
			p.AvailableServers = append(p.AvailableServers[:i], p.AvailableServers[i+1:]...)
		}
	}
}

func (p *ServerPool) addToPool(server Server) {
	for _, srv := range p.AvailableServers {
		if server.Addr == srv.Addr {
			return
		}
	}
	p.AvailableServers = append(p.AvailableServers, &server)
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
		for idx := range p.Servers {
			_, err := http.Get(p.Servers[idx].Addr + p.Servers[idx].HealthcheckEndpoint)
			if err != nil {
				p.Servers[idx].Alive = false
				p.removeFromPool(*p.Servers[idx])
			} else {
				p.Servers[idx].Alive = true
				p.Servers[idx].LastAlive = JSONTime(time.Now())
				p.addToPool(*p.Servers[idx])
			}
		}
		time.Sleep(p.healthcheckInterval * time.Second)
	}
}

func (p *ServerPool) getNext() *Server {
	if len(p.AvailableServers) == 0 {
		return nil
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	srv := p.AvailableServers[p.currentIdx%len(p.AvailableServers)]
	p.currentIdx++
	return srv
}
