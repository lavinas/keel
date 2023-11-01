package domain

import (
	"testing"
)

func TestClientGetFormated(t *testing.T) {
	client := NewClient(nil)
	client.Load("123", "Test Xxxx", "test", 1489149007, 5511999999999, "test@test.com")
	id, name, nick, doc, phone, email := client.GetFormatted()
	if id != "123" {
		t.Errorf("Error: ID should be 123. it was %s", id)
	}
	if name != "Test Xxxx" {
		t.Errorf("Error: Name should be Test. It was %s", name)
	}
	if nick != "test" {
		t.Errorf("Error: Nickname should be test. It was %s", nick)
	}
	if doc != "01489149007" {
		t.Errorf("Error: Document should be 94786984000. It was %s", doc)
	}
	if phone != "5511999999999" {
		t.Errorf("Error: Phone should be 5511999999999. It was %s", phone)
	}
	if email != "test@test.com" {
		t.Errorf("Error: Email should be test@test.com. It was %s", email)
	}
	client.Load("123", "Test Xxxx", "test", 1545838000123, 5511999999999, "test@test.com")
	_, _, _, doc, _, _ = client.GetFormatted()
	if doc != "01545838000123" {
		t.Errorf("Error: Document should be 1545838000123. It was %s", doc)
	}
}

func TestClientInsert(t *testing.T) {
	t.Run("should insert a client", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		client.Insert("Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
		if client.ID == "" {
			t.Errorf("Error: ID should not be empty")
		}
		if client.Name != "Test Xxxx" {
			t.Errorf("Error: Name should be Test. It was %s", client.Name)
		}
		if client.Nickname != "test" {
			t.Errorf("Error: Nickname should be test. It was %s", client.Nickname)
		}
		if client.Document != 94786984000 {
			t.Errorf("Error: Document should be 94786984000. It was %d", client.Document)
		}
		if client.Phone != 5511999999999 {
			t.Errorf("Error: Phone should be 5511999999999. It was %d", client.Phone)
		}
		if client.Email != "test@test.com" {
			t.Errorf("Error: Email should be test@test.com. It was %s", client.Email)
		}
	})
}

func TestClientLoad(t *testing.T) {
	t.Run("should load a client", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		client.Load("123", "Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
		if client.ID != "123" {
			t.Errorf("Error: ID should be 123. it was %s", client.ID)
		}
		if client.Name != "Test Xxxx" {
			t.Errorf("Error: Name should be Test. It was %s", client.Name)
		}
		if client.Nickname != "test" {
			t.Errorf("Error: Nickname should be test. It was %s", client.Nickname)
		}
		if client.Document != 94786984000 {
			t.Errorf("Error: Document should be 94786984000. It was %d", client.Document)
		}
		if client.Phone != 5511999999999 {
			t.Errorf("Error: Phone should be 5511999999999. It was %d", client.Phone)
		}
		if client.Email != "test@test.com" {
			t.Errorf("Error: Email should be test@test.com. It was %s", client.Email)
		}
	})
}

func TestClientLoadById(t *testing.T) {
	t.Run("should load a client by id", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		b, err := client.LoadById("123")
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if b {
			t.Errorf("Error: Client should be found")
		}
	})
}

func TestClientLoadByNick(t *testing.T) {
	t.Run("should load a client by nick", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		b, err := client.LoadByNick("test")
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if b {
			t.Errorf("Error: Client should be found")
		}
	})
}

func TestClientLoadByEmail(t *testing.T) {
	t.Run("should load a client by email", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		b, err := client.LoadByEmail("test@test.com")
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if b {
			t.Errorf("Error: Client should be found")
		}
	})
}

func TestClientLoadByDoc(t *testing.T) {
	t.Run("should load a client by doc", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		b, err := client.LoadByDoc(94786984000)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if b {
			t.Errorf("Error: Client should be found")
		}
	})
}

func TestClientLoadByPhone(t *testing.T) {
	t.Run("should load a client by phone", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		b, err := client.LoadByPhone(5511999999999)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if b {
			t.Errorf("Error: Client should be found")
		}
	})
}

func TestClientUpdate(t *testing.T) {
	t.Run("should update a client", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		err := client.Update()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
}

func TestClientSave(t *testing.T) {
	t.Run("should close a client", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		err := client.Save()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
}

func TestClientDocumentDuplicity(t *testing.T) {
	t.Run("should return true for document duplicity", func(t *testing.T) {
		repo := &RepoMock{
			ClientDocumentDuplicityReturn: true,
		}
		client := NewClient(repo)
		b, err := client.DocumentDuplicity()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if !b {
			t.Errorf("Error: Client should be found")
		}
	})
}

func TestClientEmailDuplicity(t *testing.T) {
	t.Run("should return true for email duplicity", func(t *testing.T) {
		repo := &RepoMock{
			ClientEmailDuplicityReturn: true,
		}
		client := NewClient(repo)
		b, err := client.EmailDuplicity()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if !b {
			t.Errorf("Error: Client should be found")
		}
	})
}

func TestClientNickDuplicity(t *testing.T) {
	t.Run("should return true for nick duplicity", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		b, err := client.NickDuplicity()
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if b {
			t.Errorf("Error: Client should not be found")
		}
	})
}

func TestClientGet(t *testing.T) {
	t.Run("should get a client", func(t *testing.T) {
		repo := &RepoMock{}
		client := NewClient(repo)
		client.Load("123", "Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
		id, name, nick, doc, phone, email := client.Get()
		if id != "123" {
			t.Errorf("Error: ID should be 123. it was %s", id)
		}
		if name != "Test Xxxx" {
			t.Errorf("Error: Name should be Test. It was %s", name)
		}
		if nick != "test" {
			t.Errorf("Error: Nickname should be test. It was %s", nick)
		}
		if doc != 94786984000 {
			t.Errorf("Error: Document should be 94786984000. It was %d", doc)
		}
		if phone != 5511999999999 {
			t.Errorf("Error: Phone should be 5511999999999. It was %d", phone)
		}
		if email != "test@test.com" {
			t.Errorf("Error: Email should be test@test.com. It was %s", email)
		}
	})
}
