package product

import (
	"github.com/igorverse/go-web-server-poc/internal/domain"
	"github.com/igorverse/go-web-server-poc/pkg/storage"
)

type FileRepository struct {
	db storage.Storage
}

func NewFileRepository(db storage.Storage) Repository {
	return &FileRepository{
		db: db,
	}
}

func (fr *FileRepository) GetAll() ([]domain.Product, error) {
	var ps []domain.Product
	fr.db.Read(ps)

	return ps, nil
}
