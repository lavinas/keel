package domain

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

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
