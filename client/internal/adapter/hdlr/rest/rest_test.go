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

func TestPing(t *testing.T) {
	t.Run("should ping", func(t *testing.T) {
		l := LogMock{}
		s := ServiceMock{}
		h := NewHandlerRest(&l, &s)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		h.Ping(ctx)
		if w.Code != http.StatusOK {
			t.Errorf("Invalid result: %v", w.Code)
		}
		if w.Body.String() != "{\"message\":\"pong\"}" {
			t.Errorf("Invalid result: %v", w.Body.String())
		}
	})
}

func TestInsert(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	h := NewHandlerRest(&l, &s)
	content := map[string]interface{}{
		"name":     "Jose da Silva",
		"nickname": "jose_da_silva_222",
		"document": "206.656.600-49",
		"phone":    "+5511999999999",
		"email":    "test@test.com",
	}
	t.Run("should insert", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "ok"
		MockJsonPost(ctx, content)
		h.Insert(ctx)
		if w.Code != http.StatusCreated {
			t.Errorf("Invalid result: %v", w.Code)
		}
		if w.Body.String() != "{\"id\":\"1\",\"name\":\"Jose da Silva\",\"nickname\":\"jose_da_silva_222\",\"document\":\"206.656.600-49\",\"phone\":\"+5511999999999\",\"email\":\"test@test.com\"}" {
			t.Errorf("Invalid result: %v", w.Body.String())
		}
	})
	// Test json error - bad request - invalid json body
	t.Run("should return error when invalid json body", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "ok"
		MockJsonPost(ctx, "xxxx")
		h.Insert(ctx)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Invalid result: %v", w.Code)
		}
		if w.Body.String() != "{\"error\":\"invalid json body\"}" {
			t.Errorf("Invalid result: %v", w.Body.String())
		}
	})
	// Test parameters error - bad request
	t.Run("should return error when invalid params", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "bad request"
		MockJsonPost(ctx, content)
		h.Insert(ctx)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Invalid result: %v", w.Code)
		}
		if w.Body.String() != "{\"error\":\"bad request: invalid xxxx\"}" {
			t.Errorf("Invalid result: %v", w.Body.String())
		}
	})
	// Test internal error
	t.Run("should return error when internal error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "internal error"
		MockJsonPost(ctx, content)
		h.Insert(ctx)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("Invalid result: %v", w.Code)
		}
		if w.Body.String() != "{\"error\":\"internal error\"}" {
			t.Errorf("Invalid result: %v", w.Body.String())
		}
	})

}

func TestFind(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	// Error - no content
	t.Run("should return error no content", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "no content"
		h := NewHandlerRest(&l, &s)
		h.Find(ctx)
		if w.Code != http.StatusNoContent {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
	})
	// Ok
	t.Run("should find", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "ok"
		h := NewHandlerRest(&l, &s)
		MockJsonGet(ctx, nil)
		h.Find(ctx)
		if w.Code != http.StatusOK {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
		if w.Body.String() != "{\"page\":1,\"per_page\":10,\"clients\":[{\"id\":\"1\",\"name\":\"Jose da Silva\",\"nickname\":\"jose_da_silva_222\",\"document\":\"206.656.600-49\",\"phone\":\"+5511999999999\",\"email\":\"test@test.com\"}]}" {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
	})
	// Error - bad request
	t.Run("should return error when invalid params", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "bad request"
		h := NewHandlerRest(&l, &s)
		MockJsonGet(ctx, nil)
		h.Find(ctx)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
		if w.Body.String() != "{\"error\":\"bad request: invalid xxxx\"}" {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
	})
}

func TestUpdate(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	// Ok
	t.Run("should update", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "ok"
		h := NewHandlerRest(&l, &s)
		MockJsonPost(ctx, nil)
		h.Update(ctx)
		if w.Code != http.StatusOK {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
	})
	// Error - bad request - invalid json body
	t.Run("should return error when invalid json body", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		MockJsonPost(ctx, "xxxx")
		h := NewHandlerRest(&l, &s)
		h.Update(ctx)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
	})
	// Error - bad request - invalid params
	t.Run("should return error when invalid params", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "bad request"
		h := NewHandlerRest(&l, &s)
		MockJsonPost(ctx, nil)
		h.Update(ctx)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
		if w.Body.String() != "{\"error\":\"bad request: invalid json body\"}" {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
	})
}

func TestGet(t *testing.T) {
	l := LogMock{}
	s := ServiceMock{}
	// Ok
	t.Run("should get", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "ok"
		h := NewHandlerRest(&l, &s)
		h.Get(ctx)
		if w.Code != http.StatusOK {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
	})
	// Error - bad request - invalid params
	t.Run("should return error when invalid params", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		s.Status = "bad request"
		h := NewHandlerRest(&l, &s)
		MockJsonPost(ctx, nil)
		h.Get(ctx)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
		if w.Body.String() != "{\"error\":\"bad request: invalid json body\"}" {
			t.Errorf("Invalid result: %v %s", w.Code, w.Body.String())
		}
	})
}
