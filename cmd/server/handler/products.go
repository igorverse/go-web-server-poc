package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igorverse/go-web-server-poc/internal/products"
)

type CreateRequestDTO struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"Price"`
	Stock       int     `json:"Stock"`
	Code        string  `json:"Code"`
	IsPublished bool    `json:"isPublished"`
}

type ProductHandler struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductHandler {
	return &ProductHandler{
		service: p,
	}
}

func (c *ProductHandler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Get(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ps, err := c.service.GetAll()

		id := ctx.Query("id")
		name := ctx.Query("name")
		color := ctx.Query("color")
		price := ctx.Query("price")
		stock := ctx.Query("stock")
		code := ctx.Query("code")
		isPublished := ctx.Query("isPublished")
		createdAt := ctx.Query("createdAt")

		var filteredProducts []products.Product

		// TODO: it must be refactored to an elegant solution
		for _, p := range ps {
			if id != "" {
				convId, _ := strconv.Atoi(id)

				if convId != int(p.ID) {
					continue
				}
			}
			if name != "" && p.Name != name {
				continue
			}
			if color != "" && p.Color != color {
				continue
			}
			if price != "" {
				convPrice, _ := strconv.ParseFloat(price, 64)

				if convPrice != p.Price {
					continue
				}
			}
			if stock != "" {
				convStock, _ := strconv.Atoi(stock)

				if convStock != p.Stock {
					continue
				}
			}
			if code != "" && p.Code != code {
				continue
			}
			if isPublished != "" {
				convIsPublished, _ := strconv.ParseBool(isPublished)

				if convIsPublished != p.IsPublished {
					continue
				}
			}
			if createdAt != "" && p.CreatedAt != createdAt {
				continue
			}

			filteredProducts = append(filteredProducts, p)
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, filteredProducts)
	}
}

func (c *ProductHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateRequestDTO
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.IsPublished)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, p)
	}
}

func (c *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req CreateRequestDTO
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Update(id, req.Name, req.Color, req.Price, req.Stock, req.Code, req.IsPublished)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, p)
	}
}
