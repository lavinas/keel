package hdlr

import (
	"net/http"
	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/service"

)
// HandlerChi is a service to handlle the client web requests
type HandlerChi struct {
	r *chi.Mux
	service *service.ClientService
}

// NewHandlerChi creates a new HandlerChi service
func NewHandlerChi(s *service.ClientService) *HandlerChi {
	r := chi.NewRouter()
	h := &HandlerChi{r: r, service: s}
	r.Post("/client/create", h.CreateClient)
	r.Get("/client/list", h.ListAll)
	http.ListenAndServe(":8000", r)
	return h
}

// CreateClient creates a new client
func (h *HandlerChi) CreateClient(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateInputDto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.service.Create(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// ListAll list all clients
func (h *HandlerChi) ListAll(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.ListAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}


