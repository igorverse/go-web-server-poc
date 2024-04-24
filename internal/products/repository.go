package products

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

var ps []Product

type Repository interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) GetOne(id int) (Product, error) {
	p := Product{}

	return p, nil
}
