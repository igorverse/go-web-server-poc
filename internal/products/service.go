package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, color string, price float64, stock int, code string, isPublished bool) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(name string, color string, price float64, stock int, code string, isPublished bool) (Product, error) {
	product, err := s.repository.Store(name, color, price, stock, code, isPublished)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
