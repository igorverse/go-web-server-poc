package product

import "errors"

var (
	ErrNotFound = errors.New("product not found")
	ErrInternal = errors.New("internal server error")
)
