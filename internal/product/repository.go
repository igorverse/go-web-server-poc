package product

import "github.com/igorverse/go-web-server-poc/internal/domain"

type Repository interface {
	Get(id uint64) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Store(p domain.Product) (domain.Product, error)
	Update(p domain.Product) (domain.Product, error)
	lastID() (uint64, error)
}

type repository struct{}

func NewRepository() Repository {
	return &MemoryRespository{}
}
