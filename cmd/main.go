package main

import (
	"pizzaria/internal/data"
	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()

	router := gin.Default()

	router.GET("/pizzas", handler.GetPizzas)
	router.GET("/pizzas/:id", handler.GetPizzaByID)
	router.POST("/pizzas", handler.CreatePizza)
	router.DELETE("/pizzas/:id", handler.DeletePizza)
	router.PUT("/pizzas/:id", handler.UpdatePizza)
	router.POST("/pizzas/:id/reviews", handler.CreateReview)

	router.Run(":8080")
}
