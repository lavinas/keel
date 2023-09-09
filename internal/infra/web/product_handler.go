package web

import (
	"encoding/json"
	"net/http"

	"github.com/lavinas/keel/internal/usecase"
)

type ProductHandler struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductUseCase   *usecase.ListProductUseCase
}

func NewProductUseCase(createProductUseCase *usecase.CreateProductUseCase, listProductUseCase *usecase.ListProductUseCase) *ProductHandler {
	return &ProductHandler{
		CreateProductUseCase: createProductUseCase,
		ListProductUseCase:   listProductUseCase,
	}
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (h *ProductHandler) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListProductUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
