/*
307832097 Roy Kosary
311319313 Nir Agam
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds the application settings.
type Config struct {
	Producers      []ProducerConfig `json:"Producers"`
	Zones          []ZoneConfig     `json:"Zones"`
	ZoneQueueSize  int              `json:"ZoneQueueSize"`
	HTTPServerPort int              `json:"HTTPServerPort"`
}

// ProducerConfig describes each producer's settings.
type ProducerConfig struct {
	Restaurant string `json:"Restaurant"`
	FoodType   string `json:"FoodType"`
	Orders     int    `json:"Orders"`
	QueueSize  int    `json:"QueueSize"`
}

// ZoneConfig describes each zone's settings.
type ZoneConfig struct {
	Name    string `json:"Name"`
	Workers int    `json:"Workers"`
}

// LoadConfig reads settings from a JSON file.
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("unable to parse config file: %v", err)
	}
	return &config, nil
}
