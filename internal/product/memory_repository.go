package product

import (
	"fmt"

	"github.com/igorverse/go-web-server-poc/internal/domain"
)

var ps []domain.Product
var lastID int

type MemoryRespository struct{}

func (m *MemoryRespository) Get(id int) (domain.Product, error) {
	if id > len(ps) {
		fmt.Println("Bugs")
		return domain.Product{}, ErrNotFound
	}

	return ps[int(id)-1], nil
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

func (m *MemoryRespository) UpdateNameAndPrice(p domain.Product) (domain.Product, error) {
	ps[p.ID-1].Name = p.Name
	ps[p.ID-1].Price = p.Price

	return ps[p.ID-1], nil
}

func (m *MemoryRespository) lastID() (int, error) {
	return lastID, nil
}
