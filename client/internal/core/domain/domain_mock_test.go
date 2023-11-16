package domain

import (
	"strconv"

	"github.com/lavinas/keel/client/internal/core/port"
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

type FindOutputDtoMock struct {
	Page    uint64
	PerPage uint64
	Clients []Client
}

func (f *FindOutputDtoMock) SetPage(page, perPage uint64) {
	f.Page = page
	f.PerPage = perPage
}

func (f *FindOutputDtoMock) Append(id, name, nick, doc, phone, email string) {
	client := NewClient(&RepoMock{})
	idoc, _ := strconv.ParseUint(doc, 10, 64)
	iphone, _ := strconv.ParseUint(phone, 10, 64)
	client.Load(id, name, nick, idoc, iphone, email)
	f.Clients = append(f.Clients, *client)
}

func (f *FindOutputDtoMock) Count() int {
	return len(f.Clients)
}
