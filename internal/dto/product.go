package dto

type CreateProductDTO struct {
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"Price" binding:"required"`
	Stock       int     `json:"Stock" binding:"required"`
	Code        string  `json:"Code" binding:"required"`
	IsPublished bool    `json:"isPublished" binding:"required"`
}

type UpdatedProductDTO struct {
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"Price" binding:"required"`
	Stock       int     `json:"Stock" binding:"required"`
	Code        string  `json:"Code" binding:"required"`
	IsPublished bool    `json:"isPublished" binding:"required"`
}
