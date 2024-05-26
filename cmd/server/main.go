package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/igorverse/go-web-server-poc/cmd/server/handler"
	"github.com/igorverse/go-web-server-poc/docs"
	"github.com/igorverse/go-web-server-poc/internal/product"
	"github.com/igorverse/go-web-server-poc/pkg/storage"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db := storage.NewFileStorage("file", "products.json")
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	productHandler := handler.NewProduct(service)

	server := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := server.Group("/products")
	pr.POST("/", productHandler.Store())
	pr.GET("/", productHandler.GetAll())
	pr.GET("/:id", productHandler.Get())
	pr.PUT("/:id", productHandler.Update())
	pr.PATCH("/:id", productHandler.UpdateNameAndPrice())
	pr.DELETE("/:id", productHandler.Delete())
	server.Run()
}
