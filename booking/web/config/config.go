package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache     bool                `json:"use_cache"`
	InfoLog      *log.Logger         `json:"info_log"`
	InProduction bool                `json:"in_production"`
	Sessions     *scs.SessionManager `json:"sessions"`
}

var (
	config     *AppConfig
	configOnce sync.Once
)

// LoadConfig loads the configuration from a JSON file
func LoadConfig(filePath string) error {
	// Read the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Parse the JSON data into the AppConfig struct
	var appConfig AppConfig
	err = json.Unmarshal(data, &appConfig)
	if err != nil {
		return err
	}

	config = &appConfig
	return nil
}

// GetConfig returns the application configuration
func GetConfig() *AppConfig {
	configOnce.Do(func() {
		if config == nil {
			log.Fatal("Configuration has not been loaded.")
		}
	})
	return config
}
