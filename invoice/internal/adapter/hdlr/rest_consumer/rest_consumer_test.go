package restconsumer

import (
	"net/http"
	"testing"

	"github.com/lavinas/keel/invoice/internal/core/dto"
	"github.com/lavinas/keel/invoice/pkg/gin_mock"
)

var (
	default_consumer = map[string]string{
		"id":       "1",
		"name":     "John Doe",
		"nickname": "consumer_doe",
		"document": "12345678900",
		"phone":    "11999999999",
		"email":    "test@test.com",
	}
)

func TestNewRestConsumer(t *testing.T) {
	t.Run("should not return nil", func(t *testing.T) {
		rc := NewRestConsumer(&ConfigMock{})
		if rc == nil {
			t.Error("NewRestConsumer(&ConfigMock{}) should not return nil")
		}
	})
}

func TestRestconsumer_baseByNickname(t *testing.T) {
	t.Run("should return a client", func(t *testing.T) {
		g := ginmock.NewGinMock(8085)
		g.Start("/client/get/nickname/:nickname", "GET", http.StatusOK, default_consumer)
		defer g.Stop()
		rc := NewRestConsumer(&ConfigMock{})
		rc.consumer_base = "http://localhost:8085/client/get/nickname"
		dto := dto.NewGetClientByNicknameInputDto()
		b, err := rc.GetClientByNickname("consumer_doe", dto)
		if err != nil {
			t.Errorf("Expected nil, got error: %s", err.Error())
		}
		if b == false {
			t.Error("Expected true, got false")
		}
		if dto.GetId() != "1" {
			t.Errorf("Expected 1, got %s", dto.GetId())
		}
	})
	t.Run("should not return a client", func(t *testing.T) {
		g := ginmock.NewGinMock(8085)
		g.Start("/client/get/nickname/:nickname", "GET", http.StatusNoContent, default_consumer)
		defer g.Stop()
		rc := NewRestConsumer(&ConfigMock{})
		rc.consumer_base = "http://localhost:8085/client/get/nickname"
		dto := dto.NewGetClientByNicknameInputDto()
		b, err := rc.GetClientByNickname("consumer_doe_not_found", dto)
		if err != nil {
			t.Errorf("Expected nil, got error: %s", err.Error())
		}
		if b == true {
			t.Error("Expected false, got true")
		}
		if dto.GetId() != "" {
			t.Errorf("Expected empty string, got %s", dto.GetId())
		}
	})
	t.Run("should return a url malformed error", func(t *testing.T) {
		rc := NewRestConsumer(&ConfigMock{})
		rconsumer_base := rc.consumer_base
		rc.consumer_base = "123"
		defer func() {
			rc.consumer_base = rconsumer_base
		}()
		dto := dto.NewGetClientByNicknameInputDto()
		b, err := rc.GetClientByNickname("consumer_doe", dto)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if b == true {
			t.Error("Expected false, got true")
		}
	})
	t.Run("should return a http get error", func(t *testing.T) {
		rc := NewRestConsumer(&ConfigMock{})
		rconsumer_base := rc.consumer_base
		rc.consumer_base = "http://localhost:8083/client/get_error"
		defer func() {
			rc.consumer_base = rconsumer_base
		}()
		dto := dto.NewGetClientByNicknameInputDto()
		b, err := rc.GetClientByNickname("consumer_doe", dto)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if b == true {
			t.Error("Expected false, got true")
		}
	})
	t.Run("should return a http status error", func(t *testing.T) {
		g := ginmock.NewGinMock(8085)
		g.Start("/client/get/nickname/:nickname", "GET", http.StatusNotFound, default_consumer)
		defer g.Stop()
		rc := NewRestConsumer(&ConfigMock{})
		rc.consumer_base = "http://localhost:8085/client/get/nickname"
		dto := dto.NewGetClientByNicknameInputDto()
		b, err := rc.GetClientByNickname("consumer_doe", dto)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if b == true {
			t.Error("Expected false, got true")
		}
	})
	t.Run("should return a json unmarshal error", func(t *testing.T) {
		g := ginmock.NewGinMock(8085)
		g.Start("/client/get/nickname/:nickname", "GET", http.StatusOK, "invalid json")
		defer g.Stop()
		rc := NewRestConsumer(&ConfigMock{})
		rc.consumer_base = "http://localhost:8085/client/get/nickname"
		dto := dto.NewGetClientByNicknameInputDto()
		b, err := rc.GetClientByNickname("consumer_doe", dto)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if b == true {
			t.Error("Expected false, got true")
		}
	})
	t.Run("should return a KEEL_INVOICE_CLIENT_URL not set error", func(t *testing.T) {
		rc := NewRestConsumer(&ConfigMock{})
		rconsumer_base := rc.consumer_base
		rc.consumer_base = ""
		defer func() {
			rc.consumer_base = rconsumer_base
		}()
		dto := dto.NewGetClientByNicknameInputDto()
		b, err := rc.GetClientByNickname("consumer_doe", dto)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if b == true {
			t.Error("Expected false, got true")
		}
	})
	


}
