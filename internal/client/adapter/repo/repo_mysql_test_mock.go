package repo

import (
	"fmt"
)

// Config Mock
type ConfigMock struct {
	ConnectionOk bool
}

var ConfigFields = map[string]string{
	"host":      "127.0.0.1",
	"port":      "3306",
	"user":      "root",
	"pass":      "root",
	"dbname":    "keel_client",
	"pool_size": "3",
}

func (c *ConfigMock) GetField(group string, field string) (string, error) {
	if group == "mysql" {
		r, ok := ConfigFields[field]
		if !ok {
			return "", fmt.Errorf("field %s not found", field)
		}
		return r, nil
	}
	return "", nil
}

func (c *ConfigMock) GetGroup(group string) (map[string]interface{}, error) {
	var r map[string]interface{}
	return r, nil
}
