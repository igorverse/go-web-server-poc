package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igorverse/go-web-server-poc/internal/domain"
	"github.com/igorverse/go-web-server-poc/internal/dto"
	"github.com/igorverse/go-web-server-poc/internal/product"
	"github.com/igorverse/go-web-server-poc/pkg/web"
)

type ProductHandler struct {
	service product.Service
}

func NewProduct(p product.Service) *ProductHandler {
	return &ProductHandler{
		service: p,
	}
}

// ListProduct godoc
// @Summary List product
// @Tags Products
// @Description get product
// @Produce  json
// @Param id path int true "product id"
// @Success 200 {object} web.Response{data=domain.Product} "OK"
// @Failure 400 {object} web.Response "Bad Request"
// @Failure 404 {object} web.Response "Not Found"
// @Router /products/{id} [get]
func (c *ProductHandler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid product id"))
			return
		}

		p, err := c.service.Get(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Produce  json
// @Success 200 {object} web.Response{data=[]domain.Product} "OK"
// @Failure 400 {object} web.Response "Bad Request"
// @Failure 404 {object} web.Response "Not Found"
// @Router /products [get]
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

		var filteredProducts []domain.Product

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
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if len(ps) == 0 {
			ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, filteredProducts, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductDTO true "Product to store"
// @Success 201 {object} web.Response{data=domain.Product} "Created"
// @Failure 400 {object} web.Response "Bad Request"
// @Failure 422 {object} web.Response "Unprocessable Entity"
// @Router /products [post]
func (c *ProductHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productDTO dto.CreateProductDTO
		if err := ctx.ShouldBindJSON(&productDTO); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		p, err := c.service.Store(productDTO)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}

// UpdateProducts godoc
// @Summary Update products
// @Tags Products
// @Description update products
// @Accept  json
// @Produce  json
// @Param product body dto.UpdatedProductDTO true "Product to update"
// @Success 200 {object} web.Response{data=domain.Product} "OK"
// @Failure 400 {object} web.Response "Bad Request"
// @Failure 422 {object} web.Response "Unprocessable Entity"
// @Router /products/{id} [put]
func (c *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		var updateProductDTO dto.UpdatedProductDTO
		if err := ctx.ShouldBindJSON(&updateProductDTO); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		p, err := c.service.Update(int(id), updateProductDTO)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// UpdateNameAndPrice godoc
// @Summary Update products name and price
// @Tags Products
// @Description update products name and price
// @Accept  json
// @Produce  json
// @Param product body dto.UpdatedNameAndPriceDTO true "Product to update name and price"
// @Success 200 {object} web.Response{data=domain.Product} "OK"
// @Failure 400 {object} web.Response "Bad Request"
// @Failure 422 {object} web.Response "Unprocessable Entity"
// @Router /products/{id} [patch]
func (c *ProductHandler) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		var updatedNameAndPriceDTO dto.UpdatedNameAndPriceDTO
		if err := ctx.ShouldBindJSON(&updatedNameAndPriceDTO); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		p, err := c.service.UpdateNameAndPrice(int(id), updatedNameAndPriceDTO)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// DeleteProduct godoc
// @Summary Delete product
// @Tags Products
// @Description delete product
// @Produce  json
// @Param id path int true "product id"
// @Success 204 {object} web.Response{data=domain.Product} "No Content"
// @Failure 400 {object} web.Response "Bad Request"
// @Failure 404 {object} web.Response "Not Found"
// @Router /products/{id} [delete]
func (c *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, fmt.Sprintf("product %d was removed", id), ""))
	}
}
