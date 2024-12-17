package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var PromTotalRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Number of requests",
}, []string{"path"})

var PromServersAlive = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "http_servers_alive",
	Help: "Number of currently active servers",
})

type PromServersAliveCollector struct {
	gaugeDesc *prometheus.Desc
}

func (c *PromServersAliveCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.gaugeDesc
}

func (c *PromServersAliveCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.gaugeDesc, prometheus.GaugeValue, float64(len(ServerPoolObj.AvailableServers)))
}

func NewPromServersAliveCollector() *PromServersAliveCollector {
	return &PromServersAliveCollector{
		gaugeDesc: prometheus.NewDesc("http_servers_alive", "Number of currently active servers", nil, nil),
	}
}

func PromRegister() {
	prometheus.Register(PromTotalRequests)
	prometheus.Register(NewPromServersAliveCollector())
}
