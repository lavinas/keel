package dto

import (
	"testing"
)

func TestFindOutDtoAppend(t *testing.T) {
	t.Run("should append the given value to the slice", func(t *testing.T) {
		f := FindOutputDto{}
		f.Append("a", "b", "c", "d", "e", "f")
		f.Append("g", "h", "i", "j", "k", "l")
		if len(f.Clients) != 2 {
			t.Errorf("expected 2, got %d", len(f.Clients))
		}
		id, name, nick, doc, phone, email := f.Clients[0].Get()
		if id != "a" {
			t.Errorf("expected a, got %s", id)
		}
		if name != "b" {
			t.Errorf("expected b, got %s", name)
		}
		if nick != "c" {
			t.Errorf("expected c, got %s", nick)
		}
		if doc != "d" {
			t.Errorf("expected d, got %s", doc)
		}
		if phone != "e" {
			t.Errorf("expected e, got %s", phone)
		}
		if email != "f" {
			t.Errorf("expected f, got %s", email)
		}
		id, name, nick, doc, phone, email = f.Clients[1].Get()
		if id != "g" {
			t.Errorf("expected g, got %s", id)
		}
		if name != "h" {
			t.Errorf("expected h, got %s", name)
		}
		if nick != "i" {
			t.Errorf("expected i, got %s", nick)
		}
		if doc != "j" {
			t.Errorf("expected j, got %s", doc)
		}
		if phone != "k" {
			t.Errorf("expected k, got %s", phone)
		}
		if email != "l" {
			t.Errorf("expected l, got %s", email)
		}

	})
}

func TestFindOutDtoCount(t *testing.T) {
	t.Run("should return the number of elements in the slice", func(t *testing.T) {
		f := FindOutputDto{}
		f.Append("a", "b", "c", "d", "e", "f")
		f.Append("g", "h", "i", "j", "k", "l")
		if f.Count() != 2 {
			t.Errorf("expected 2, got %d", f.Count())
		}
	})
}

func TestFindOutDtoSetPage(t *testing.T) {
	t.Run("should set the page and perPage", func(t *testing.T) {
		f := FindOutputDto{}
		f.SetPage(1, 10)
		if f.Page != 1 {
			t.Errorf("expected 1, got %d", f.Page)
		}
		if f.PerPage != 10 {
			t.Errorf("expected 10, got %d", f.PerPage)
		}
	})
}
