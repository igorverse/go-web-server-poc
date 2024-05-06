package products

type Repository interface {
	GetAll() ([]Product, error)
	Store(name string, color string, price float64, stock int, code string, isPublished bool) (Product, error)
	lastID() (uint, error)
}

type repository struct{}

func NewRepository() Repository {
	return &MemoryRespository{}
}
