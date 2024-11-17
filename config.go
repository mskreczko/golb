package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
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

func (c *Config) GetConfig(filename string) {
	yamlFile, err := os.ReadFile(filename)
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
	if c.ListeningAddr == ":" {
		return fmt.Sprintf(":%s", c.ListeningPort)
	}
	return fmt.Sprintf("%s:%s", c.ListeningAddr, c.ListeningPort)
}
