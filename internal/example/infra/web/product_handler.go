package web

import (
	"encoding/json"
	"net/http"

	"github.com/lavinas/keel/internal/example/service"
)

type ProductHandler struct {
	InsertProductService *service.InsertProductService
	ListProductService   *service.ListProductService
}

func NewProductService(createProductService *service.InsertProductService, listProductService *service.ListProductService) *ProductHandler {
	return &ProductHandler{
		InsertProductService: createProductService,
		ListProductService:   listProductService,
	}
}

func (h *ProductHandler) InsertProductHandler(w http.ResponseWriter, r *http.Request) {
	var input service.InsertProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.InsertProductService.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *ProductHandler) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListProductService.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
