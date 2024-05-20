package dto

type CreateProductDTO struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"Price"`
	Stock       int     `json:"Stock"`
	Code        string  `json:"Code"`
	IsPublished bool    `json:"isPublished"`
}

type UpdatedProductDTO struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"Price"`
	Stock       int     `json:"Stock"`
	Code        string  `json:"Code"`
	IsPublished bool    `json:"isPublished"`
}
