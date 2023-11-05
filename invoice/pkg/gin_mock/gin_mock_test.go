package ginmock

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestGinMock(t *testing.T) {
	t.Run("should test a get call", func(t *testing.T) {
		g := NewGinMock(8080)
		defer g.Stop()
		g.Start("/", "GET", 200, map[string]string{"message": "hello world"})
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
		var resp map[string]string
		if err := json.Unmarshal(data, &resp); err != nil {
			t.Errorf("error: %v", err)
		}
		if resp["message"] != "hello world" {
			t.Errorf("error: %v", resp["message"])
		}
	})
	t.Run("should test a post call", func(t *testing.T) {
		g := NewGinMock(8080)
		defer g.Stop()
		g.Start("/", "POST", 200, map[string]string{"message": "hello world"})
		response, err := http.Post("http://localhost:8080", "application/json", nil)
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
		var resp map[string]string
		if err := json.Unmarshal(data, &resp); err != nil {
			t.Errorf("error: %v", err)
		}
		if resp["message"] != "hello world" {
			t.Errorf("error: %v", resp["message"])
		}
	})
	t.Run("should test a put call", func(t *testing.T) {
		g := NewGinMock(8080)
		defer g.Stop()
		g.Start("/", "PUT", 200, map[string]string{"message": "hello world"})
		request, err := http.NewRequest("PUT", "http://localhost:8080", nil)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		response, err := http.DefaultClient.Do(request)
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
		var resp map[string]string
		if err := json.Unmarshal(data, &resp); err != nil {
			t.Errorf("error: %v", err)
		}
		if resp["message"] != "hello world" {
			t.Errorf("error: %v", resp["message"])
		}
	})
	t.Run("should test a delete call", func(t *testing.T) {
		g := NewGinMock(8080)
		defer g.Stop()
		g.Start("/", "DELETE", 200, map[string]string{"message": "hello world"})
		request, err := http.NewRequest("DELETE", "http://localhost:8080", nil)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		response, err := http.DefaultClient.Do(request)
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
		var resp map[string]string
		if err := json.Unmarshal(data, &resp); err != nil {
			t.Errorf("error: %v", err)
		}
		if resp["message"] != "hello world" {
			t.Errorf("error: %v", resp["message"])
		}
	})
	t.Run("should test a default call", func(t *testing.T) {
		g := NewGinMock(8080)
		defer g.Stop()
		g.Start("/", "DEFAULT", 200, map[string]string{"message": "hello world"})
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
		var resp map[string]string
		if err := json.Unmarshal(data, &resp); err != nil {
			t.Errorf("error: %v", err)
		}
		if resp["message"] != "hello world" {
			t.Errorf("error: %v", resp["message"])
		}
	})
	t.Run("should test a port 0 call", func(t *testing.T) {
		g := NewGinMock(0)
		defer g.Stop()
		g.Start("/", "GET", 200, map[string]string{"message": "hello world"})
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
		var resp map[string]string
		if err := json.Unmarshal(data, &resp); err != nil {
			t.Errorf("error: %v", err)
		}
		if resp["message"] != "hello world" {
			t.Errorf("error: %v", resp["message"])
		}
	})

}