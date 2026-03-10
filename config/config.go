package config

import "os"

type Config struct {
	ServerName string
	Version    string
	LogLevel   string
}

func Load() *Config {
	return &Config{
		ServerName: getEnv("SERVER_NAME", "zsv-mcp"),
		Version:    getEnv("VERSION", "v1.0.0"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
