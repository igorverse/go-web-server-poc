package product

import (
	"time"

	"github.com/igorverse/go-web-server-poc/internal/domain"
	"github.com/igorverse/go-web-server-poc/internal/dto"
)

type Service interface {
	Get(id uint64) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Store(p dto.CreateProductDTO) (domain.Product, error)
	Update(id uint64, p dto.UpdatedProductDTO) (domain.Product, error)
}

type service struct {
	repository Repository
}

func (s *service) Get(id uint64) (domain.Product, error) {
	product, err := s.repository.Get(id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) GetAll() ([]domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(p dto.CreateProductDTO) (domain.Product, error) {
	product := domain.Product{
		Name:        p.Name,
		Color:       p.Color,
		Price:       p.Price,
		Stock:       p.Stock,
		Code:        p.Code,
		IsPublished: p.IsPublished,
		CreatedAt:   time.Now().Format("02-01-2006"),
	}

	return s.repository.Store(product)
}

func (s *service) Update(id uint64, p dto.UpdatedProductDTO) (domain.Product, error) {
	currentProduct, err := s.repository.Get(id)

	if err != nil {
		return domain.Product{}, err
	}

	if p.Name != "" {
		currentProduct.Name = p.Name
	}

	if p.Color != "" {
		currentProduct.Color = p.Color
	}

	if p.Price > 0 {
		currentProduct.Price = p.Price
	}

	if p.Stock >= 0 {
		currentProduct.Stock = p.Stock
	}

	if p.Code != "" {
		currentProduct.Code = p.Code
	}

	updatedProduct, err := s.repository.Update(currentProduct)
	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
