package service

import (
	"github.com/lavinas/keel/internal/client/core/port"
)

// Service are services to orchestrate client domain
type Service struct {
	domain port.Domain
	log    port.Log
	repo   port.Repo
}
// NewCreate creates a new Create service
func NewService(domain port.Domain, log port.Log, repo port.Repo) *Service {
	return &Service{
		domain: domain,
		log:    log,
		repo:   repo,
	}
}
// Orquestration of Creating a new client
func (s *Service) ClientCreate(input port.CreateInputDto, output port.CreateOutputDto) error {
	return ServiceClientCreate(s.log, s.domain, input, output)
}
