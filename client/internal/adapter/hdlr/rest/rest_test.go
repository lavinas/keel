package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRun(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	h := NewHandlerRest(&l, &s)
	h.Run()
}

func TestInsert(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	h := NewHandlerRest(&l, &s)
	// Test ok
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	s.Status = "ok"
	content := map[string]interface{}{
		"name":     "Jose da Silva",
		"nickname": "jose_da_silva_222",
		"document": "206.656.600-49",
		"phone":    "+5511999999999",
		"email":    "test@test.com",
	}
	MockJsonPost(ctx, content)
	h.Insert(ctx)
	if w.Code != http.StatusCreated {
		t.Errorf("Invalid result: %v", w.Code)
	}
	if w.Body.String() != "{\"id\":\"1\",\"name\":\"Jose da Silva\",\"nickname\":\"jose_da_silva_222\",\"document\":\"206.656.600-49\",\"phone\":\"+5511999999999\",\"email\":\"test@test.com\"}" {
		t.Errorf("Invalid result: %v", w.Body.String())
	}
	// Test json error - bad request - invalid json body
	w = httptest.NewRecorder()
	ctx = GetTestGinContext(w)
	s.Status = "ok"
	MockJsonPost(ctx, "xxxx")
	h.Insert(ctx)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Invalid result: %v", w.Code)
	}
	if w.Body.String() != "{\"error\":\"invalid json body\"}" {
		t.Errorf("Invalid result: %v", w.Body.String())
	}
	// Test parameters error - bad request
	w = httptest.NewRecorder()
	ctx = GetTestGinContext(w)
	s.Status = "bad request"
	MockJsonPost(ctx, content)
	h.Insert(ctx)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Invalid result: %v", w.Code)
	}
	// Test internal error
	w = httptest.NewRecorder()
	ctx = GetTestGinContext(w)
	s.Status = "internal error"
	MockJsonPost(ctx, content)
	h.Insert(ctx)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Invalid result: %v", w.Code)
	}

}

func TestFind(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	// Error - no content
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	s.Status = "no content"
	h := NewHandlerRest(&l, &s)
	h.Find(ctx)
	if w.Code != http.StatusNoContent {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
	// Ok
	w = httptest.NewRecorder()
	ctx = GetTestGinContext(w)
	s.Status = "ok"
	h = NewHandlerRest(&l, &s)
	MockJsonGet(ctx, nil)
	h.Find(ctx)
	if w.Code != http.StatusOK {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
	if w.Body.String() != "{\"page\":1,\"per_page\":10,\"clients\":[{\"id\":\"1\",\"name\":\"Jose da Silva\",\"nickname\":\"jose_da_silva_222\",\"document\":\"206.656.600-49\",\"phone\":\"+5511999999999\",\"email\":\"test@test.com\"}]}" {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
	// Error - bad request
	w = httptest.NewRecorder()
	ctx = GetTestGinContext(w)
	s.Status = "bad request"
	h = NewHandlerRest(&l, &s)
	MockJsonGet(ctx, nil)
	h.Find(ctx)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
}

func TestUpdate(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	// Ok
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	s.Status = "ok"
	h := NewHandlerRest(&l, &s)
	MockJsonPost(ctx, nil)
	h.Update(ctx)
	if w.Code != http.StatusOK {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
	// Error - bad request - invalid json body
	w = httptest.NewRecorder()
	ctx = GetTestGinContext(w)
	MockJsonPost(ctx, "xxxx")
	h.Update(ctx)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
	// Error - bad request - invalid params
	w = httptest.NewRecorder()
	ctx = GetTestGinContext(w)
	s.Status = "bad request"
	h = NewHandlerRest(&l, &s)
	MockJsonPost(ctx, nil)
	h.Update(ctx)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
}

func TestGet(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	// Ok
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	s.Status = "ok"
	h := NewHandlerRest(&l, &s)
	h.Get(ctx)
	if w.Code != http.StatusOK {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
	// Error - bad request - invalid params
	w = httptest.NewRecorder()
	ctx = GetTestGinContext(w)
	s.Status = "bad request"
	h = NewHandlerRest(&l, &s)
	MockJsonPost(ctx, nil)
	h.Get(ctx)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
	}
}
