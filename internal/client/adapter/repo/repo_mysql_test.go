package repo

import (
	"strings"
	"testing"

	"github.com/lavinas/keel/internal/client/core/domain"
)

func TestNewRepoMysql(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	// Ok
	if repo.db == nil {
		t.Errorf("Error: db should not be nil")
	}
}

func TestSave(t *testing.T) {
	// Ok
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	client := domain.NewClient(repo)
	err := client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@tets.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	id, _, _, _, _, _ := client.Get()
	b, err := repo.GetById(id, client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Client should be found")
	}
	if client.Name != "Test Xxxx" {
		t.Errorf("Error: Client name should be Test Xxxx")
	}
	// Error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 := domain.NewClient(repo2)
	err = client2.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo2.Save(client2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	client := domain.NewClient(repo)
	err := client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	client.Name = "Test Yyyy"
	err = repo.Update(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// Error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 := domain.NewClient(repo2)
	err = client2.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo2.Update(client2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestDocumentDuplicity(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	// check not duplicated
	b, err := repo.DocumentDuplicity(94786984000, "")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Document should not be duplicated")
	}
	// check duplicated
	client := domain.NewClient(repo)
	err = client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	b, err = repo.DocumentDuplicity(94786984000, "")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Document should be duplicated")
	}
	// check error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 := domain.NewClient(repo2)
	err = client2.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	_, err = repo2.DocumentDuplicity(94786984000, "")
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestEmailDuplicityQuery(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	// check not duplicated
	b, err := repo.EmailDuplicity("test@test.com", "")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Email should not be duplicated")
	}
	// check duplicated
	client := domain.NewClient(repo)
	err = client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	b, err = repo.EmailDuplicity("test@test.com", "")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Email should be duplicated")
	}
	// check error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 := domain.NewClient(repo2)
	err = client2.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	_, err = repo2.EmailDuplicity("test@test.com", "")
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestNickDuplicityQuery(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	// check not duplicated
	b, err := repo.NickDuplicity("test", "")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Nick should not be duplicated")
	}
	// check duplicated
	client := domain.NewClient(repo)
	err = client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	b, err = repo.NickDuplicity("test", "")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Nick should be duplicated")
	}
	// check error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 := domain.NewClient(repo2)
	err = client2.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	_, err = repo2.NickDuplicity("test@test.com", "")
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestLoadSet(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	// Ok
	client := domain.NewClient(repo)
	err := client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// by name
	clientset := domain.NewClientSet(repo)
	err = repo.LoadSet(1, 10, "test", "", "", "", "", clientset)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if clientset.Count() != 1 {
		t.Errorf("Error: Client set should have 1 client, got %d", clientset.Count())
	}
	// by nick
	clientset = domain.NewClientSet(repo)
	err = repo.LoadSet(1, 10, "", "test", "", "", "", clientset)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if clientset.Count() != 1 {
		t.Errorf("Error: Client set should have 1 client, got %d", clientset.Count())
	}
	// by doc
	clientset = domain.NewClientSet(repo)
	err = repo.LoadSet(1, 10, "", "", "947869", "", "", clientset)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if clientset.Count() != 1 {
		t.Errorf("Error: Client set should have 1 client, got %d", clientset.Count())
	}
	// by phone
	clientset = domain.NewClientSet(repo)
	err = repo.LoadSet(1, 10, "", "", "", "11999", "", clientset)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if clientset.Count() != 1 {
		t.Errorf("Error: Client set should have 1 client, got %d", clientset.Count())
	}
	// by email
	clientset = domain.NewClientSet(repo)
	err = repo.LoadSet(1, 10, "", "", "", "", "test", clientset)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if clientset.Count() != 1 {
		t.Errorf("Error: Client set should have 1 client, got %d", clientset.Count())
	}

	// check error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	clientset2 := domain.NewClientSet(repo2)
	err = repo2.LoadSet(1, 10, "test", "", "", "", "", clientset2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestGetById(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	client := domain.NewClient(repo)
	err := client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// Ok
	id, _, _, _, _, _ := client.Get()
	b, err := repo.GetById(id, client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Client should be found")
	}
	if client.Name != "Test Xxxx" {
		t.Errorf("Error: Client name should be Test Xxxx")
	}
	if client.Nickname != "test" {
		t.Errorf("Error: Client nick should be test")
	}
	if client.Document != 94786984000 {
		t.Errorf("Error: Client document should be 94786984000")
	}
	if client.Phone != 5511999999999 {
		t.Errorf("Error: Client phone should be 5511999999999")
	}
	if client.Email != "test@test.com" {
		t.Errorf("Error: Client email should be test@test.com")
	}
	// Not found
	b, err = repo.GetById("999999999999999999", client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Client should not be found")
	}
	// Error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 := domain.NewClient(repo2)
	_, err = repo2.GetById("999999999999999999", client2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestGetByNick(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	client := domain.NewClient(repo)
	err := client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// Ok
	client2 := domain.NewClient(repo)
	b, err := repo.GetByNick("test", client2)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Client should be found")
	}
	if client2.Name != "Test Xxxx" {
		t.Errorf("Error: Client name should be Test Xxxx, got %s", client2.Name)
	}
	if client2.Nickname != "test" {
		t.Errorf("Error: Client nick should be test, got %s", client2.Nickname)
	}
	if client2.Document != 94786984000 {
		t.Errorf("Error: Client document should be 94786984000, got %d", client2.Document)
	}
	if client2.Phone != 5511999999999 {
		t.Errorf("Error: Client phone should be 5511999999999, got %d", client2.Phone)
	}
	if client2.Email != "test@test.com" {
		t.Errorf("Error: Client email should be test@test.com")
	}
	// Not found
	b, err = repo.GetByNick("999999999999999999", client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Client should not be found")
	}
	// Error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 = domain.NewClient(repo2)
	_, err = repo2.GetByNick("999999999999999999", client2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestGetByEmail(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	client := domain.NewClient(repo)
	err := client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// Ok
	client2 := domain.NewClient(repo)
	b, err := repo.GetByEmail("test@test.com", client2)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Client should be found")
	}
	if client2.Name != "Test Xxxx" {
		t.Errorf("Error: Client name should be Test Xxxx, got %s", client2.Name)
	}
	if client2.Nickname != "test" {
		t.Errorf("Error: Client nick should be test, got %s", client2.Nickname)
	}
	if client2.Document != 94786984000 {
		t.Errorf("Error: Client document should be 94786984000, got %d", client2.Document)
	}
	if client2.Phone != 5511999999999 {
		t.Errorf("Error: Client phone should be 5511999999999, got %d", client2.Phone)
	}
	if client2.Email != "test@test.com" {
		t.Errorf("Error: Client email should be test@test.com")
	}
	// Not found
	b, err = repo.GetByEmail("999999999999999999", client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Client should not be found")
	}
	// Error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 = domain.NewClient(repo2)
	_, err = repo2.GetByEmail("999999999999999999", client2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestGetByDoc(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	client := domain.NewClient(repo)
	err := client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// Ok
	client2 := domain.NewClient(repo)
	b, err := repo.GetByDoc(94786984000, client2)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Client should be found")
	}
	if client2.Name != "Test Xxxx" {
		t.Errorf("Error: Client name should be Test Xxxx, got %s", client2.Name)
	}
	if client2.Nickname != "test" {
		t.Errorf("Error: Client nick should be test, got %s", client2.Nickname)
	}
	if client2.Document != 94786984000 {
		t.Errorf("Error: Client document should be 94786984000, got %d", client2.Document)
	}
	if client2.Phone != 5511999999999 {
		t.Errorf("Error: Client phone should be 5511999999999, got %d", client2.Phone)
	}
	if client2.Email != "test@test.com" {
		t.Errorf("Error: Client email should be test@test.com")
	}
	// Not found
	b, err = repo.GetByDoc(999999, client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Client should not be found")
	}
	// Error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 = domain.NewClient(repo2)
	_, err = repo2.GetByDoc(94786984000, client2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}

func TestGetByPhone(t *testing.T) {
	config := ConfigMock{}
	repo := NewRepoMysql(&config)
	repo.Truncate()
	defer repo.Close()
	defer repo.Truncate()
	client := domain.NewClient(repo)
	err := client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	err = repo.Save(client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// Ok
	client2 := domain.NewClient(repo)
	b, err := repo.GetByPhone(5511999999999, client2)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if !b {
		t.Errorf("Error: Client should be found")
	}
	if client2.Name != "Test Xxxx" {
		t.Errorf("Error: Client name should be Test Xxxx, got %s", client2.Name)
	}
	if client2.Nickname != "test" {
		t.Errorf("Error: Client nick should be test, got %s", client2.Nickname)
	}
	if client2.Document != 94786984000 {
		t.Errorf("Error: Client document should be 94786984000, got %d", client2.Document)
	}
	if client2.Phone != 5511999999999 {
		t.Errorf("Error: Client phone should be 5511999999999, got %d", client2.Phone)
	}
	if client2.Email != "test@test.com" {
		t.Errorf("Error: Client email should be test@test.com")
	}
	// Not found
	b, err = repo.GetByPhone(999999, client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if b {
		t.Errorf("Error: Client should not be found")
	}
	// Error connection
	ConfigFields["user"] = "error"
	defer func() {
		ConfigFields["user"] = "root"
	}()
	repo2 := NewRepoMysql(&config)
	defer repo2.Close()
	defer repo2.Truncate()
	client2 = domain.NewClient(repo2)
	_, err = repo2.GetByPhone(5511999999999, client2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	if !strings.Contains(err.Error(), "Access denied") {
		t.Errorf("Error: %s", err)
	}
}
