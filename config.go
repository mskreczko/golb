package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Addr                string `yaml:"addr"`
	HealthcheckEndpoint string `yaml:"healthcheck"`
}

type Config struct {
	Servers []ServerConfig `yaml:"servers"`
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
}
