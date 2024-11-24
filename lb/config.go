package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"
)

type ServerConfig struct {
	Addr                string `yaml:"addr"`
	HealthcheckEndpoint string `yaml:"healthcheck"`
}

type Config struct {
	ListeningAddr       string         `yaml:"listening_addr,omitempty"`
	ListeningPort       string         `yaml:"listening_port,omitempty"`
	HealthcheckInterval int            `yaml:"healthcheck_interval,omitempty"`
	Servers             []ServerConfig `yaml:"servers"`
}

func (c *Config) GetConfig(configPath string) {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	if c.ListeningAddr == "" {
		c.ListeningAddr = ":"
	}
	if c.ListeningPort == "" {
		c.ListeningPort = "80"
	}
	if c.HealthcheckInterval == 0 {
		c.HealthcheckInterval = 5
	}
}

func (c *Config) GetFullAddress() string {
	return net.JoinHostPort(c.ListeningAddr, c.ListeningPort)
}
