package product

import (
	"github.com/igorverse/go-web-server-poc/internal/domain"
)

var ps []domain.Product
var lastID uint64

type MemoryRespository struct{}

func (m *MemoryRespository) Get(id uint64) (domain.Product, error) {
	return ps[id-1], nil
}

func (m *MemoryRespository) GetAll() ([]domain.Product, error) {
	return ps, nil
}

func (m *MemoryRespository) Store(p domain.Product) (domain.Product, error) {
	lastID++
	p.ID = int(lastID)

	ps = append(ps, p)

	return p, nil
}

func (m *MemoryRespository) Update(p domain.Product) (domain.Product, error) {
	ps[p.ID-1] = domain.Product{
		ID:          int(p.ID),
		Name:        p.Name,
		Color:       p.Color,
		Price:       p.Price,
		Stock:       p.Stock,
		Code:        p.Code,
		IsPublished: p.IsPublished,
		CreatedAt:   p.CreatedAt,
	}

	return ps[p.ID-1], nil
}

func (m *MemoryRespository) lastID() (uint64, error) {
	return lastID, nil
}
