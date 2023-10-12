package service

import (
	"encoding/json"
	"strings"
	"errors"
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
	Status string
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
	if c.Status == "invalid" {
		return "", fmt.Errorf("invalid config")
	}
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
	ClientNicknameDuplicityReturn bool
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
	if r.ClientNicknameDuplicityReturn {
		return true, nil
	}
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

// Domain Mock
type DomainMock struct {
	Status string
}
func (d *DomainMock) GetClient() port.Client {
	c := &ClientMock{}
	c.Status = d.Status
	return c
}

func (d *DomainMock) GetClientSet() port.ClientSet {
	c := &ClientSetMock{}
	c.Status = d.Status
	return c
}

// Client Mock
type ClientMock struct {
	Status string
	Id   string
	Name string
	Nick string
	Doc  uint64
	Phone uint64
	Email string
}
func (c *ClientMock) Insert(name, nick string, doc, phone uint64, email string) error {
	if c.Status == "ok" {
		return nil
	}
	if c.Status == "loaderror" {
		return fmt.Errorf("internal error")
	}
	return nil
}
func (c *ClientMock) Load(id, name, nick string, doc, phone uint64, email string) {
	c.Id = id
	c.Name = name
	c.Nick = nick
	c.Doc = doc
	c.Phone = phone
	c.Email = email
}
func (c *ClientMock) LoadById(id string) (bool, error) {
	if c.Status == "findbyid" {
		return true, nil
	}
	if strings.Contains(c.Status, "duplicity") {
		return true, nil
	}
	if strings.Contains(c.Status, "update") {
		return true, nil
	}
	if c.Status == "findbyiderror" {
		return false, errors.New("findbyid error")
	}
	return false, nil
}
func (c *ClientMock) LoadByNick(nick string) (bool, error) {
	if c.Status == "findbynick" {
		return true, nil
	}
	if c.Status == "findbynickerror" {
		return false, errors.New("findbynick error")
	}
	return false, nil
}
func (c *ClientMock) LoadByEmail(email string) (bool, error) {
	if c.Status == "findbyemail" {
		return true, nil
	}
	if c.Status == "findbyemailerror" {
		return false, errors.New("findbyemail error")
	}
	return false, nil
}
func (c *ClientMock) LoadByDoc(doc uint64) (bool, error) {
	if c.Status == "findbydoc" {
		return true, nil
	}
	if c.Status == "findbydocerror" {
		return false, errors.New("findbydoc error")
	}
	return false, nil
}
func (c *ClientMock) LoadByPhone(phone uint64) (bool, error) {
	if c.Status == "findbyphone" {
		return true, nil
	}
	if c.Status == "findbyphoneerror" {
		return false, errors.New("findbyphone error")
	}
	return false, nil
}
func (c *ClientMock) DocumentDuplicity() (bool, error) {
	if c.Status == "duplicitydocument" {
		return true, nil
	}
	if c.Status == "duplicitydocumenterror" {
		return false, errors.New("duplicitydocument error")
	}
	return false, nil
}
func (c *ClientMock) EmailDuplicity() (bool, error) {
	if c.Status == "duplicityemail" {
		return true, nil
	}
	if c.Status == "duplicityemailerror" {
		return false, errors.New("duplicitydocument error")
	}
	return false, nil
}
func (c *ClientMock) NickDuplicity() (bool, error) {
	if c.Status == "duplicitynick" {
		return true, nil
	}
	if c.Status == "duplicitynickerror" {
		return false, errors.New("duplicitydocument error")
	}
	return false, nil
}
func (c *ClientMock) Get() (string, string, string, uint64, uint64, string) {
	return "0", "0", "Test Name", 0, 0, "eeqwewqewqe"
}
func (c *ClientMock) GetFormatted() (string, string, string, string, string, string) {
	return "0", "0", "Test Name", "Test", "00222222222", "eeqwewqewqe"
}
func (c *ClientMock) Save() error {
	if c.Status == "saveerror" {
		return fmt.Errorf("internal error")
	}
	return nil
}
func (c *ClientMock) Update() error {
	if c.Status == "updateerror" {
		return fmt.Errorf("internal error")
	}
	return nil
}

// ClientSet Mock
type ClientSetMock struct {
	Status string
}
func (c *ClientSetMock) Load(page, perPage uint64, name, nick, doc, phone, email string) error {
	if c.Status == "ok" {
		return nil
	}
	if c.Status == "internal" {
		return fmt.Errorf("internal error")
	}
	return nil
}
func (c *ClientSetMock) Append(id, name, nick string, doc, phone uint64, email string) {
}
func (c *ClientSetMock) SetOutput(output port.FindOutputDto) {
}

// FindInputDtoMock
type FindInputDtoMock struct {
	Status string
}
func (f *FindInputDtoMock) Validate() error {
	if f.Status == "ok" {
		return nil
	}
	if f.Status == "invalid" {
		return fmt.Errorf("invalid input")
	}
	return nil
}
func (f *FindInputDtoMock) Get() (string, string, string, string, string, string, string) {
	if f.Status == "blank" {
		return "", "", "", "", "", "", ""
	}
	if f.Status == "invalid" {
		return "a", "b", "c", "d", "e", "f", "g"
	}
	return "0", "0", "Test Name", "Test", "00222222222", "eeqwewqewqe", "eqeqwewewe"
}

type FindOutputDtoMock struct {
}
func (f *FindOutputDtoMock) SetPage(page, perPage uint64) {
}
func (f *FindOutputDtoMock) Append(id, name, nick, doc, phone, email string) {
}
func (f *FindOutputDtoMock) Count() int {
	return 0
}

// Insert input dto mock
type InsertInputDtoMock struct {
	Status string
}
func (i *InsertInputDtoMock) IsBlank() bool {
	return i.Status == "blank"
}
func (i *InsertInputDtoMock) Validate() error {
	if i.Status == "ok" {
		return nil
	}
	if i.Status == "invalid" {
		return fmt.Errorf("invalid input")
	}
	return nil
}
func (i *InsertInputDtoMock) Format() error {
	if i.Status == "ok" {
		return nil
	}
	if i.Status == "invalid" {
		return fmt.Errorf("invalid input")
	}
	return nil
}
func (i *InsertInputDtoMock) Get() (string, string, string, string, string) {
	if i.Status == "blank" {
		return "", "", "", "", ""
	}
	if i.Status == "invalid" {
		return "a", "b", "c", "d", "e"
	}
	return "name", "nick", "doc", "phone", "email"	
}

// Insert output dto mock
type InsertOutputDtoMock struct {
	Status string
	Id    string
	Name  string
	Nick  string
	Doc   string
	Phone string
	Email string
}
func (i *InsertOutputDtoMock) Fill(id, name, nick, doc, phone, email string) {
	i.Id = id
	i.Name = name
	i.Nick = nick
	i.Doc = doc
	i.Phone = phone
	i.Email = email
}
func (i *InsertOutputDtoMock) Get() (string, string, string, string, string, string) {
	return i.Id, i.Name, i.Nick, i.Doc, i.Phone, i.Email
}

// Update input dto mock
type UpdateInputDtoMock struct {
	Status string
}
func (u *UpdateInputDtoMock) Validate() error {
	if u.Status == "ok" {
		return nil
	}
	if u.Status == "invalid" {
		return fmt.Errorf("invalid input")
	}
	return nil
}
func (u *UpdateInputDtoMock) Get() (string, string, string, string, string) {
	if u.Status == "blank" {
		return "", "", "", "", ""
	}
	if u.Status == "invalid" {
		return "a", "b", "c", "d", "e"
	}
	return "name", "nick", "doc", "phone", "email"
}
func (u *UpdateInputDtoMock) Format() error {
	if u.Status == "ok" {
		return nil
	}
	if u.Status == "formaterror" {
		return fmt.Errorf("invalid input")
	}
	return fmt.Errorf("internal error")
}
func (u *UpdateInputDtoMock) IsBlank() bool {
	return u.Status == "blank"
}

// Update output dto mock
type UpdateOutputDtoMock struct {
	Status string
	Id    string
	Name  string
	Nick  string
	Doc   string
	Phone string
	Email string
}
func (u *UpdateOutputDtoMock) Fill(id, name, nick, doc, phone, email string) {
	u.Id = id
	u.Name = name
	u.Nick = nick
	u.Doc = doc
	u.Phone = phone
	u.Email = email
}


