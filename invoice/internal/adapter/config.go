package adapter

import "os"

// Config is the configuration handler for the application
type Config struct {
}

// NewConfig creates a new configuration handler
func NewConfig() *Config {
	return &Config{}
}

// Get returns the value of the environment variable
func (c *Config) Get(param string) string {
	return os.Getenv(param)
}
