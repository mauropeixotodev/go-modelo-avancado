package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(200, data.Pizzas)
}

func CreatePizza(c *gin.Context) {
	var pizza models.Pizza

	if err := c.ShouldBindJSON(&pizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	pizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, pizza)
	savePizzas()

	c.JSON(201, pizza)
}

func DeletePizza(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			savePizzas()
			c.JSON(200, gin.H{"message": "Pizza deletada com sucesso"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Pizza não encontrada"})
}

func UpdatePizza(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			var updatedPizza models.Pizza

			if err := c.ShouldBindJSON(&updatedPizza); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id
			savePizzas()
			c.JSON(200, data.Pizzas[i])
			return
		}
	}

	c.JSON(404, gin.H{"error": "Pizza não encontrada"})
}

func GetPizzaByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	for _, pizza := range data.Pizzas {
		if pizza.ID == id {
			c.JSON(200, pizza)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Pizza não encontrada"})
}

func savePizzas() {
	file, err := os.Create("./dados/pizza.json")
	if err != nil {
		fmt.Print("Erro ao criar o arquivo:", err)
		return
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data.Pizzas); err != nil {
		fmt.Print("Erro ao codificar o arquivo:", err)
		return
	}

}
