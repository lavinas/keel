package service

import (
	"os"

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

func (l *LogMock) Error(msg string) {
	l.mtype = "Error"
	l.msg = msg
}

func (l *LogMock) Close() {
}

// Repo Mock
type RepoMock struct {
}

func (r *RepoMock) Create(client *domain.Client) error {
	return nil
}
