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

func (r *RepoMock) Save(client port.Client) error {
	r.client = client
	return nil
}

func (r *RepoMock) DocumentDuplicity(document uint64, id string) (bool, error) {
	if r.ClientDocumentDuplicityReturn {
		return true, nil
	}
	return false, nil
}

func (r *RepoMock) EmailDuplicity(email, id string) (bool, error) {
	if r.ClientEmailDuplicityReturn {
		return true, nil
	}
	return false, nil
}

func (r *RepoMock) NickDuplicity(nick, id string) (bool, error) {
	return false, nil
}

func (r *RepoMock) LoadSet(page, perPage uint64, name, nick, doc, phone, email string, set port.ClientSet) error {
	return nil
}

func (r *RepoMock) GetById(id string, client port.Client) (bool, error) {
	return false, nil
}

func (r *RepoMock) GetByNick(nick string, client port.Client) (bool, error) {
	return false, nil
}

func (r *RepoMock) GetByEmail(email string, client port.Client) (bool, error) {
	return false, nil
}

func (r *RepoMock) GetByDoc(doc uint64, client port.Client) (bool, error) {
	return false, nil
}

func (r *RepoMock) GetByPhone(phone uint64, client port.Client) (bool, error) {
	return false, nil
}

func (r *RepoMock) Update(client port.Client) error {
	return nil
}

func (r *RepoMock) Close() error {
	return nil
}
