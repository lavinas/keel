package restconsumer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/lavinas/keel/internal/invoice/core/port"
)

var (
	// api url
	getClientUrl = "http://localhost:8081/client/get"
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
	r, err := url.JoinPath(getClientUrl, nickname)
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
	if response.StatusCode == http.StatusNotFound {
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
