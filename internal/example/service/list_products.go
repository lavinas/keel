package service

import "github.com/lavinas/keel/internal/example/domain"

type ListProductsOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductService struct {
	ProductRepository domain.ProductRepository
}

func NewListProductService(productRepository domain.ProductRepository) *ListProductService {
	return &ListProductService{
		ProductRepository: productRepository,
	}
}

func (u *ListProductService) Execute() ([]*ListProductsOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var productsOutput []*ListProductsOutputDto
	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil

}
