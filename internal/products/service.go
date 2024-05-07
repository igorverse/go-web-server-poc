package products

type Service interface {
	Get(id uint64) (Product, error)
	GetAll() ([]Product, error)
	Store(name string, color string, price float64, stock int, code string, isPublished bool) (Product, error)
	Update(id uint64, name string, color string, price float64, stock int, code string, isPublished bool) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) Get(id uint64) (Product, error) {
	product, err := s.repository.Get(id)
	if err != nil {
		return Product{}, err
	}

	return product, nil
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

func (s *service) Update(id uint64, name string, color string, price float64, stock int, code string, isPublished bool) (Product, error) {
	currentProduct, err := s.repository.Get(id)

	if err != nil {
		return Product{}, err
	}

	if name != "" {
		currentProduct.Name = name
	}

	if color != "" {
		currentProduct.Color = color
	}

	if price > 0 {
		currentProduct.Price = price
	}

	if stock >= 0 {
		currentProduct.Stock = stock
	}

	if code != "" {
		currentProduct.Code = code
	}

	updatedProduct, err := s.repository.Update(currentProduct)
	if err != nil {
		return Product{}, err
	}

	return updatedProduct, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
