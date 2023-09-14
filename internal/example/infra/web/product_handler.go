package web

import (
	"encoding/json"
	"net/http"

	"github.com/lavinas/keel/internal/example/service"
)

type ProductHandler struct {
	CreateProductService *service.CreateProductService
	ListProductService   *service.ListProductService
}

func NewProductService(createProductService *service.CreateProductService, listProductService *service.ListProductService) *ProductHandler {
	return &ProductHandler{
		CreateProductService: createProductService,
		ListProductService:   listProductService,
	}
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input service.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.CreateProductService.Execute(input)
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
