package models

import (
	"os"

	"github.com/BurntSushi/toml"
)

var (
	configPath = os.Getenv("STORAGE_CONFIG")
)

// Storage configuration
type Config struct {
	Port int      `toml:"Port"`
	DB   DBConfig `toml:"DB"`
}

type DBConfig struct {
	Host       string `toml:"Host"`
	Port       int    `toml:"Port"`
	Database   string `toml:"Database"`
	Collection string `toml:"Collection"`
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		path = configPath
	}
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	return &cfg, err
}
