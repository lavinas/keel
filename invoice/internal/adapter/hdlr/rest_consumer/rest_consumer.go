package restconsumer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/lavinas/keel/invoice/internal/core/port"
)

var (
	// api url
	consumer_base = os.Getenv("KEEL_INVOICE_CLIENT_URL")
)

// Rest comsumer implements ApiConsumer interface
type RestConsumer struct {
}

// NewRestComsumer is the constructor of RestComsumer
func NewRestConsumer() *RestConsumer {
	return &RestConsumer{}
}

// GetClientByNickname returns a GetClientByNicknameInputDto
func (rc *RestConsumer) GetClientByNickname(nickname string, client port.GetClientByNicknameInputDto) (bool, error) {
	if consumer_base == "" {
		return false, fmt.Errorf("error: KEEL_INVOICE_CLIENT_URL is not set")
	}
	r, err := url.JoinPath(consumer_base, nickname)
	if err != nil {
		return false, err
	}
	response, err := http.Get(r)
	if err != nil {
		return false, err
	} 
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	if response.StatusCode == http.StatusNoContent {
		return false, nil
	}
	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("error: %v", response.Status)
	}
	if err := json.Unmarshal(data, &client); err != nil {
		return false, err
	}
	return true, nil
}
