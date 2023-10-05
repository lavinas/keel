package hdlr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	content := map[string]interface{}{
		"name":     "Jose da Silva",
		"nickname": "jose_da_silva_222",
		"document": "206.656.600-49",
		"phone":    "+55 (11) 99999-9999",
		"email":    "test@test.com.br",
	}
	MockJsonPost(ctx, content)
	l := LogMock{}
	s := ServiceMock{}

	h := NewHandlerGin(&l, &s)
	h.ClientCreate(ctx)

	if w.Code != http.StatusCreated {
		t.Errorf("Invalid result: %v", w.Code)
	}

}

func TestClientList(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	l := LogMock{}
	s := ServiceMock{}

	h := NewHandlerGin(&l, &s)
	h.ClientList(ctx)

	if w.Code != http.StatusOK {
		t.Errorf("Invalid result: %v", w.Code)
	}

	if w.Body.String() != "{\"clients\":null}" {
		t.Errorf("Invalid result: %v", w.Body.String())
	}

}
