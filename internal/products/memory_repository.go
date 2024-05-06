package products

import (
	"time"
)

var ps []Product
var lastID uint

type MemoryRespository struct{}

func (m *MemoryRespository) GetAll() ([]Product, error) {
	return ps, nil
}

func (m *MemoryRespository) Store(name string, color string, price float64, stock int, code string, isPublished bool) (Product, error) {
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
	lastID++

	return p, nil
}

func (m *MemoryRespository) lastID() (uint, error) {
	return lastID, nil
}
