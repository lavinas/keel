package service

import (
	"github.com/lavinas/keel/internal/example/domain"
)

type InsertProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type InsertProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type InsertProductService struct {
	ProductRepository domain.ProductRepository
}

func NewInsertProductService(productRepository domain.ProductRepository) *InsertProductService {
	return &InsertProductService{
		ProductRepository: productRepository,
	}
}

func (u *InsertProductService) Execute(input InsertProductInputDto) (*InsertProductOutputDto, error) {
	product := domain.NewProduct(input.Name, input.Price)
	if err := u.ProductRepository.Insert(product); err != nil {
		return nil, err
	}
	return &InsertProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil

}
