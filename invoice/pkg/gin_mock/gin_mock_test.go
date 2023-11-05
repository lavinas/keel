package ginmock

import (
	"net/http"
	"io"
	"testing"
)

func TestGinMock(t *testing.T) {
	t.Run("should start and stop the http server", func(t *testing.T) {
		g := NewGinMock("8080", "do", "GET", "hello world\n")
		g.Start()
		g.Stop()
	})
	t.Run("should test a http call", func(t *testing.T) {
		g := NewGinMock("8080", "do", "GET", "hello world\n")
		g.Start()
		response, err := http.Get("http://localhost:8080")
		if err != nil {
			t.Errorf("error: %v", err)
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			t.Errorf("error: %v", response.Status)
		}
		data, err := io.ReadAll(response.Body)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if string(data) != "hello world\n" {
			t.Errorf("error: expected %v, got %v", "hello world\n", string(data))
		}
		g.Stop()
	})

}