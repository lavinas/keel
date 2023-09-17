package hdlr

import (
	"net/http"
	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/lavinas/keel/internal/client/core/domain"
	"github.com/lavinas/keel/internal/client/core/service"

)

type HandlerChi struct {
	r *chi.Mux
	create *service.Create
	list *service.List
}

func NewHandlerChi(c *service.Create, l *service.List) *HandlerChi {
	r := chi.NewRouter()
	h := &HandlerChi{r: r, create: c, list: l}
	r.Post("/client/create", h.CreateClient)
	r.Get("/client/list", h.ListAll)
	http.ListenAndServe(":8000", r)
	return h
}

func (h *HandlerChi) CreateClient(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateInputDto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.create.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *HandlerChi) ListAll(w http.ResponseWriter, r *http.Request) {
	output, err := h.list.ListAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}


