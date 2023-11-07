package mysql

import (
	"os"
)

// ConfigMock is a mock for config
type ConfigMock struct {
}

func (c *ConfigMock) Get(key string) string {
	return os.Getenv(key)
}

func (c *ConfigMock) Set(key, value string) {
	os.Setenv(key, value)
}
