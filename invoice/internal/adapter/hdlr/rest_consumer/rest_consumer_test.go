package restconsumer

import (
	"testing"

	"github.com/lavinas/keel/invoice/internal/core/dto"
)

func TestNewRestConsumer(t *testing.T) {
	t.Run("should not return nil", func(t *testing.T) {
		rc := NewRestConsumer()
		if rc == nil {
			t.Error("NewRestConsumer() should not return nil")
		}
	})
}

func TestRestConsumer_GetClientByNickname(t *testing.T) {
	t.Run("should return a client", func(t *testing.T) {
		rc := NewRestConsumer()
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
		rc := NewRestConsumer()
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
		rGetClientUrl := getClientUrl
		getClientUrl = "123"
		defer func() {
			getClientUrl = rGetClientUrl
		}()
		rc := NewRestConsumer()
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
		rGetClientUrl := getClientUrl
		getClientUrl = "http://localhost:8083/client/get_error"
		defer func() {
			getClientUrl = rGetClientUrl
		}()
		rc := NewRestConsumer()
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
