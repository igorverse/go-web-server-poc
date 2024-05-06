package products

type Repository interface {
	Get(id uint64) (Product, error)
	GetAll() ([]Product, error)
	Store(name string, color string, price float64, stock int, code string, isPublished bool) (Product, error)
	lastID() (uint64, error)
}

type repository struct{}

func NewRepository() Repository {
	return &MemoryRespository{}
}
