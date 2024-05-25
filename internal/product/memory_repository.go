package product

import (
	"github.com/igorverse/go-web-server-poc/internal/domain"
)

var ps []domain.Product
var lastID int

type MemoryRespository struct{}

func (m *MemoryRespository) Get(id int) (domain.Product, error) {
	isFound := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			isFound = true
		}
	}

	if !isFound {
		return domain.Product{}, ErrNotFound
	}

	return ps[index], nil
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

func (m *MemoryRespository) Delete(id int) error {
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
		}
	}

	ps = append(ps[:index], ps[index+1:]...)

	return nil
}

func (m *MemoryRespository) LastID() (int, error) {
	return lastID, nil
}
