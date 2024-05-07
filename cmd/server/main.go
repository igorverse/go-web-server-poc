package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igorverse/go-web-server-poc/cmd/server/handler"
	"github.com/igorverse/go-web-server-poc/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	productHandler := handler.NewProduct(service)

	server := gin.Default()
	pr := server.Group("/products")
	pr.POST("/", productHandler.Store())
	pr.GET("/", productHandler.GetAll())
	pr.GET("/:id", productHandler.Get())
	pr.PUT("/:id", productHandler.Update())
	server.Run()
}
