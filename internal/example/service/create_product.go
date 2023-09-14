package service

import (
	"github.com/lavinas/keel/internal/example/domain"
)

type CreateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductService struct {
	ProductRepository domain.ProductRepository
}

func NewCreateProductService(productRepository domain.ProductRepository) *CreateProductService {
	return &CreateProductService{
		ProductRepository: productRepository,
	}
}

func (u *CreateProductService) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	product := domain.NewProduct(input.Name, input.Price)
	if err := u.ProductRepository.Create(product); err != nil {
		return nil, err
	}
	return &CreateProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil

}
