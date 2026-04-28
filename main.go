package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	var pizzas = []models.Pizza{models.Pizza{ID: 1, Preco: 29.99, Sabor: "Calabresa"}, models.Pizza{ID: 2, Preco: 34.99, Sabor: "Mussarela"}, models.Pizza{ID: 3, Preco: 39.99, Sabor: "Frango com Catupiry"}}
	c.JSON(200, pizzas)
}
