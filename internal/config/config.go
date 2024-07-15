// Package config is a package for loading configuration
package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Configuration struct {
	Bank BankConfig
}

type BankConfig struct {
	PosURL string `yaml:"pos_url" env:"BANK_POS_URL" envDefault:"http://localhost:8080/payments"`
}

// Config Global variable to hold the configuration
var Config Configuration

// LoadConfig initializes the configuration with default values and any environment variables
func LoadConfig() {
	// Parse environment variables and set defaults
	if err := env.Parse(&Config); err != nil {
		fmt.Printf("Failed to parse environment variables: %+v\n", err)
	}
}
