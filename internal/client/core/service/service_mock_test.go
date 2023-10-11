package service

import (
	"encoding/json"
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

type ClientMock struct {
	Status string
}

func (c *ClientMock) Insert(name, nick string, doc, phone uint64, email string) error {
	if c.Status == "ok" {
		return nil
	}
	if c.Status == "internal" {
		return fmt.Errorf("internal error")
	}
	return nil
}
func (c *ClientMock) Load(id, name, nick string, doc, phone uint64, email string) {
}

func (c *ClientMock) LoadById(id string) (bool, error) {
	if c.Status == "findbyid" {
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
	return false, nil
}
func (c *ClientMock) EmailDuplicity() (bool, error) {
	return false, nil
}
func (c *ClientMock) NickDuplicity() (bool, error) {
	return false, nil
}
func (c *ClientMock) Get() (string, string, string, uint64, uint64, string) {
	return "0", "0", "Test Name", 0, 0, "eeqwewqewqe"
}
func (c *ClientMock) GetFormatted() (string, string, string, string, string, string) {
	return "0", "0", "Test Name", "Test", "00222222222", "eeqwewqewqe"
}
func (c *ClientMock) Save() error {
	return nil
}
func (c *ClientMock) Update() error {
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

type InsertOutputDtoMock struct {
	Status string
}
func (i *InsertOutputDtoMock) Fill(id, name, nick, doc, phone, email string) {
}
func (i *InsertOutputDtoMock) Get() (string, string, string, string, string, string) {
	return "0", "0", "Test Name", "Test", "00222222222", "eeqwewqewqe"
}

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
	return "Test Name", "Test", "00222222222", "eeqwewqewqe", "eqeqwewewe"	
}



