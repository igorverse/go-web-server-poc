package domain

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}
