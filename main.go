package main

import "github.com/gin-gonic/gin"

type Pizza struct {
	ID    int
	Preco float64
	Sabor string
}

func main() {

	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}

func getPizzas(c *gin.Context) {
	var pizzas = []Pizza{Pizza{ID: 1, Preco: 29.99, Sabor: "Calabresa"}, Pizza{ID: 2, Preco: 34.99, Sabor: "Mussarela"}, Pizza{ID: 3, Preco: 39.99, Sabor: "Frango com Catupiry"}}
	c.JSON(200, pizzas)
}
