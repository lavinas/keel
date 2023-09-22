package service

import (
	"os"
	"encoding/json"

	"github.com/lavinas/keel/internal/client/core/domain"
)

// Log Mock
type LogMock struct {
	mtype string
	msg string
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

func (l *LogMock) Errorf (input any, err error) {
	b, _ := json.Marshal(input)
	l.Error(err.Error() + " | " + string(b))
}

func (l *LogMock) Close() {
}

// Repo Mock
type RepoMock struct {
	client *domain.Client
}

func (r *RepoMock) Create(client *domain.Client) error {
	r.client = client
	return nil
}
