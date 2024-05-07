package products

import (
	"time"
)

var ps []Product
var lastID uint64

type MemoryRespository struct{}

func (m *MemoryRespository) Get(id uint64) (Product, error) {
	return ps[id-1], nil
}

func (m *MemoryRespository) GetAll() ([]Product, error) {
	return ps, nil
}

func (m *MemoryRespository) Store(name string, color string, price float64, stock int, code string, isPublished bool) (Product, error) {
	lastID++
	p := Product{
		ID:          lastID,
		Name:        name,
		Color:       color,
		Price:       price,
		Stock:       stock,
		Code:        code,
		IsPublished: isPublished,
		CreatedAt:   time.Now().Local().Format("2024-02-04"),
	}

	ps = append(ps, p)

	return p, nil
}

func (m *MemoryRespository) Update(product Product) (Product, error) {
	ps[product.ID-1] = Product{
		product.ID,
		product.Name,
		product.Color,
		product.Price,
		product.Stock,
		product.Code,
		product.IsPublished,
		product.CreatedAt,
	}

	return ps[product.ID-1], nil
}

func (m *MemoryRespository) lastID() (uint64, error) {
	return lastID, nil
}
