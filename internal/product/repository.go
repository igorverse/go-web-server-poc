package product

import (
	"github.com/igorverse/go-web-server-poc/internal/domain"
	"github.com/igorverse/go-web-server-poc/pkg/storage"
)

type Repository interface {
	Get(id int) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Store(p domain.Product) (domain.Product, error)
	Update(p domain.Product) (domain.Product, error)
	UpdateNameAndPrice(p domain.Product) (domain.Product, error)
	Delete(id int) error
	LastID() (int, error)
}

type repository struct{}

func NewRepository(db storage.Storage) Repository {
	return &FileRepository{db}
}
