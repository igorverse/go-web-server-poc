package product

import (
	"github.com/igorverse/go-web-server-poc/internal/domain"
	"github.com/igorverse/go-web-server-poc/pkg/storage"
)

type FileRepository struct {
	db storage.Storage
}

func NewFileStorage(db storage.Storage) Repository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) Get(id int) (domain.Product, error) {
	var ps []domain.Product
	r.db.Read(&ps)

	for _, p := range ps {
		if p.ID == id {
			return p, nil
		}
	}

	return domain.Product{}, ErrNotFound
}

func (r *FileRepository) GetAll() ([]domain.Product, error) {
	var ps []domain.Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *FileRepository) Store(p domain.Product) (domain.Product, error) {
	var ps []domain.Product
	r.db.Read(&ps)

	lastInsertedID := 1

	if len(ps) > 0 {
		lastInsertedID = ps[len(ps)-1].ID + 1
	}

	p.ID = lastInsertedID
	ps = append(ps, p)

	err := r.db.Write(ps)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *FileRepository) Update(p domain.Product) (domain.Product, error) {
	var ps []domain.Product
	r.db.Read(&ps)

	for i, up := range ps {
		if up.ID == p.ID {
			ps[i] = domain.Product{
				ID:          int(p.ID),
				Name:        p.Name,
				Color:       p.Color,
				Price:       p.Price,
				Stock:       p.Stock,
				Code:        p.Code,
				IsPublished: p.IsPublished,
				CreatedAt:   p.CreatedAt,
			}

			err := r.db.Write(ps)
			if err != nil {
				return domain.Product{}, err
			}

			return ps[i], nil
		}
	}

	return domain.Product{}, ErrInternal
}

func (r *FileRepository) UpdateNameAndPrice(p domain.Product) (domain.Product, error) {
	var ps []domain.Product
	r.db.Read(&ps)

	for i, up := range ps {
		if up.ID == p.ID {
			ps[i].Name = p.Name
			ps[i].Price = p.Price

			err := r.db.Write(ps)
			if err != nil {
				return domain.Product{}, err
			}

			return ps[i], nil
		}
	}

	return domain.Product{}, ErrInternal
}

func (r *FileRepository) Delete(id int) error {
	var ps []domain.Product
	r.db.Read(&ps)

	for i, p := range ps {
		if p.ID == id {
			ps = append(ps[:i], ps[i+1:]...)
			err := r.db.Write(ps)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return ErrNotFound
}

func (r *FileRepository) LastID() (int, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil
}
