package pkg

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server     Server     `yaml:"server"`
	GrpcServer GrpcServer `yaml:"grpc-server"`
	Database   Database   `yaml:"database"`
}

type Server struct {
	Port int `yaml:"port"`
}
type GrpcServer struct {
	Port int `yaml:"port"`
}

type Database struct {
	Username   string `yaml:"username"`
	Host       string `yaml:"host"`
	DbName     string `yaml:"database"`
	Port       int    `yaml:"port"`
	SslMode    string `yaml:"sslmode"`
	MaxRetries int    `yaml:"max_retries"`
}

func (c *Config) LoadConfig(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Failed to read config file", "file", filePath, "error", err)
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		slog.Error("Failed to unmarshal config YAML", "error", err)
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
