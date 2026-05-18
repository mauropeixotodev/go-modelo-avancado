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

func SavePizzas() {
	file, err := os.Create("./dados/pizza.json")
	if err != nil {
		fmt.Print("Erro ao criar o arquivo:", err)
		return
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Pizzas); err != nil {
		fmt.Print("Erro ao codificar o arquivo:", err)
		return
	}

}
