package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lavinas/keel/internal/client/core/port"
)

// Log Mock
type LogMock struct {
	mtype string
	msg   string
}

func (l *LogMock) GetFile() *os.File {
	return nil
}

func (l *LogMock) Info(msg string) {
	l.mtype = "Info"
	l.msg = msg
}

func (l *LogMock) Infof(input any, message string) {
	l.mtype = "Info"
	b, _ := json.Marshal(input)
	l.Info(message + " | " + string(b))
}

func (l *LogMock) Error(msg string) {
	l.mtype = "Error"
	l.msg = msg
}

func (l *LogMock) Errorf(input any, err error) {
	b, _ := json.Marshal(input)
	l.Error(err.Error() + " | " + string(b))
}

func (l *LogMock) Close() {
}

// Config Mock
type ConfigMock struct {
}

var ConfigFields = map[string]string{
	"host":      "127.0.0.1",
	"port":      "3306",
	"user":      "root",
	"pass":      "pwd22Adm",
	"dbname":    "cbs_client",
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

// Repo Mock
type RepoMock struct {
	client                        port.Client
	ClientDocumentDuplicityReturn bool
	ClientEmailDuplicityReturn    bool
}

func (r *RepoMock) ClientSave(client port.Client) error {
	r.client = client
	return nil
}

func (r *RepoMock) ClientDocumentDuplicity(document uint64) (bool, error) {
	if r.ClientDocumentDuplicityReturn {
		return true, nil
	}
	return false, nil
}

func (r *RepoMock) ClientEmailDuplicity(email string) (bool, error) {
	if r.ClientEmailDuplicityReturn {
		return true, nil
	}
	return false, nil
}

func (r *RepoMock) ClientLoad(set port.ClientSet) error {
	return nil
}
