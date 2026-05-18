package models

type Pizza struct {
	ID     int      `json:"id"`
	Preco  float64  `json:"preco"`
	Sabor  string   `json:"sabor"`
	Review []Review `json:"reviews"`
}
