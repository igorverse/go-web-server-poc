package product

import (
	"time"

	"github.com/igorverse/go-web-server-poc/internal/domain"
	"github.com/igorverse/go-web-server-poc/internal/dto"
)

type Service interface {
	Get(id int) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Store(p dto.CreateProductDTO) (domain.Product, error)
	Update(id int, p dto.UpdatedProductDTO) (domain.Product, error)
	UpdateNameAndPrice(id int, p dto.UpdatedNameAndPriceDTO) (domain.Product, error)
}

type service struct {
	repository Repository
}

func (s *service) Get(id int) (domain.Product, error) {
	return s.repository.Get(id)
}

func (s *service) GetAll() ([]domain.Product, error) {
	return s.repository.GetAll()
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

func (s *service) Update(id int, p dto.UpdatedProductDTO) (domain.Product, error) {
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

	return s.repository.Update(currentProduct)
}

func (s *service) UpdateNameAndPrice(id int, p dto.UpdatedNameAndPriceDTO) (domain.Product, error) {
	currentProduct, err := s.repository.Get(id)
	if err != nil {
		return domain.Product{}, err
	}

	if p.Name != "" {
		currentProduct.Name = p.Name
	}

	if p.Price > 0 {
		currentProduct.Price = p.Price
	}

	return s.repository.UpdateNameAndPrice(currentProduct)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
