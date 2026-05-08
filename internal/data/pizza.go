package data

import (
	"encoding/json"
	"fmt"

	"os"
	"pizzaria/internal/models"
)

var Pizzas = []models.Pizza{}

func LoadPizzas() {

	file, err := os.Open("./dados/pizza.json")
	if err != nil {
		fmt.Print("Erro ao abrir o arquivo:", err)
		return
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Print("Erro ao decodificar o arquivo:", err)
		return
	}
	defer file.Close()

}
