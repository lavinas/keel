package restcomsumer

import (
	"github.com/lavinas/keel/internal/invoice/core/port"
)

// Rest comsumer implements ApiConsumer interface
type RestComsumer struct {
}

// NewRestComsumer is the constructor of RestComsumer	
func NewRestComsumer() *RestComsumer {
	return &RestComsumer{}
}

// GetClientByNickname returns a GetClientByNicknameInputDto
func (rc *RestComsumer) GetClientByNickname(nickname string, client port.GetClientByNicknameInputDto) error {
	return nil
}



