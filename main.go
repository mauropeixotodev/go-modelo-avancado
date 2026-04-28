package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{}

func main() {
	loadPizzas()

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
	pizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, pizza)
	savePizzas()

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

func loadPizzas() {

	file, err := os.Open("./dados/pizza.json")
	if err != nil {
		fmt.Print("Erro ao abrir o arquivo:", err)
		return
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Print("Erro ao decodificar o arquivo:", err)
		return
	}
	defer file.Close()

}

func savePizzas() {
	file, err := os.Create("./dados/pizza.json")
	if err != nil {
		fmt.Print("Erro ao criar o arquivo:", err)
		return
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Print("Erro ao codificar o arquivo:", err)
		return
	}

}
