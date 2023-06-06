// config/config.go
package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	AWS   AWSConfig   `json:"AWS"`
	Email EmailConfig `json:"Email"`
}

type AWSConfig struct {
	Region    string `json:"Region"`
	AccessKey string `json:"AccessKey"`
	SecretKey string `json:"SecretKey"`
}

type EmailConfig struct {
	Sender    string `json:"Sender"`
	Recipient string `json:"Recipient"`
	Subject   string `json:"Subject"`
	CharSet   string `json:"CharSet"`
}

func getCurrentDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

func GetConfig() (Config, error) {

	// print current folder path
	file := "app/config.json"
	getCurrentDirectory()

	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}
