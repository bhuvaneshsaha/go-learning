package config

import "log"

type AppConfig struct {
	UseCache bool
	InfoLog  *log.Logger
}
