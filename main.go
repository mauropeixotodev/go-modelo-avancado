package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{models.Pizza{ID: 1, Preco: 29.99, Sabor: "Calabresa"}, models.Pizza{ID: 2, Preco: 34.99, Sabor: "Mussarela"}, models.Pizza{ID: 3, Preco: 39.99, Sabor: "Frango com Catupiry"}}

func main() {

	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", createPizza)
	router.Run()
}

func getPizzas(c *gin.Context) {

	c.JSON(200, pizzas)
}

func createPizza(c *gin.Context) {
	var pizza models.Pizza
	if err := c.ShouldBindJSON(&pizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pizzas = append(pizzas, pizza)
	c.JSON(201, gin.H{"message": "Pizza criada com sucesso!"})

}
