package config

import (
	"os"
)

type Config struct {
	APIKey     string
	APIBaseURL string
}

func LoadConfig() *Config {
	return &Config{
		APIKey:     os.Getenv("OPENAI_API_KEY"),
		APIBaseURL: "https://api.openai.com/v1/engines/gpt-3.5-turbo/completions", // Updated model endpoint
	}
}
