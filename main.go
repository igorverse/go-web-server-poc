package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	products := router.Group("/products")
	{
		products.GET("/", GetAll)
	}

	router.Run()
}

func GetAll(ctx *gin.Context) {
	products := []Products{}

	id := ctx.Query("id")
	name := ctx.Query("name")
	color := ctx.Query("color")
	price := ctx.Query("price")
	stock := ctx.Query("stock")
	code := ctx.Query("code")
	isPublished := ctx.Query("isPublished")
	createdAt := ctx.Query("createdAt")

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

	filteredProducts := []Products{}

	// TODO: it must be refactored to an elegant solution
	for _, p := range products {
		if id != "" {
			convId, _ := strconv.Atoi(id)

			if convId != p.Id {
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

	ctx.JSON(http.StatusOK, gin.H{"data": filteredProducts})
}
