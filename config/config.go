package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const P8rConfig = "P8RCONFIG"

var (
	// DefaultP8rHome represent P8r home directory.
	DefaultP8rHome = filepath.Join(mustP8rHome(), ".p8r")
	// P8rLogs represents K9s log.
	P8rLogs = filepath.Join(os.TempDir(), fmt.Sprintf("p8r-%s.log", MustP8rUser()))
	// P8rDumpDir represents a directory where K9s screen dumps will be persisted.
	K9sDumpDir = filepath.Join(os.TempDir(), fmt.Sprintf("p8r-screens-%s", MustP8rUser()))
)

type (
	AWSSettings interface {
		CurrentProfile() (string, error)
	}
	// Config tracks p8r configuration options.
	Config struct {
		P8r      *P8r `yaml:"p8r"`
		settings AWSSettings
	}
)

// NewConfig creates a new default config.
func NewConfig(as AWSSettings) *Config {
	return &Config{P8r: NewP8r(), settings: as}
}

func (c *Config) CurrentProfile() string {
	if len(c.P8r.CurrentProfile) == 0 {
		return "default"
	}
	return c.P8r.CurrentProfile
}
