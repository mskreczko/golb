package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var PromTotalRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Number of requests",
}, []string{"path"})

var PromRequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "http_request_latency_seconds",
	Help:    "Latency of HTTP requests",
	Buckets: []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
})

type PromCollector struct {
	serversAliveDesc *prometheus.Desc
}

func (c *PromCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.serversAliveDesc
}

func (c *PromCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.serversAliveDesc, prometheus.GaugeValue, float64(len(ServerPoolObj.AvailableServers)))
}

func NewPromServersAliveCollector() *PromCollector {
	return &PromCollector{
		serversAliveDesc: prometheus.NewDesc("http_servers_alive", "Number of currently active servers", nil, nil),
	}
}

func PromRegister() {
	prometheus.Register(PromTotalRequests)
	prometheus.Register(NewPromServersAliveCollector())
	prometheus.Register(PromRequestLatency)
}
