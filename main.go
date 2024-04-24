package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messagem": "Olá, Igor!",
		})
	})

	router.GET("/products", GetAll)

	router.Run()
}

func GetAll(ctx *gin.Context) {
	products := []Products{}

	file, err := os.ReadFile("products.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal([]byte(file), &products)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(products)

	ctx.JSON(200, products)
}
