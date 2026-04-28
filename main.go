package main

import (
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Preco: 29.99, Sabor: "Calabresa"},
	{ID: 2, Preco: 34.99, Sabor: "Mussarela"},
	{ID: 3, Preco: 39.99, Sabor: "Frango com Catupiry"},
}

func main() {

	router := gin.Default()

	router.GET("/pizzas", getPizzas)
	router.GET("/pizzas/:id", getPizzaByID)
	router.POST("/pizzas", createPizza)

	router.Run(":8080")
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

	c.JSON(201, pizza)
}

func getPizzaByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(200, pizza)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Pizza não encontrada"})
}
