package product

import (
	"github.com/igorverse/go-web-server-poc/internal/domain"
	"github.com/igorverse/go-web-server-poc/pkg/storage"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
}

func NewRepository(db storage.Storage) Repository {
	return &FileRepository{db}
}
