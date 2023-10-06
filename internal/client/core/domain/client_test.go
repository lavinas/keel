package domain

import (
	"testing"
)

func TestGetFormated(t *testing.T) {
	client := NewClient(nil)
	client.Load("123", "Test Xxxx", "test", 94786984000, 5511999999999, "test@test.com")
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
	if doc != "94786984000" {
		t.Errorf("Error: Document should be 94786984000. It was %s", doc)
	}
	if phone != "5511999999999" {
		t.Errorf("Error: Phone should be 5511999999999. It was %s", phone)
	}
	if email != "test@test.com" {
		t.Errorf("Error: Email should be test@test.com. It was %s", email)
	}
}
