package products

type Service interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) GetOne(id int) (Product, error) {
	p := Product{}

	return p, nil
}
