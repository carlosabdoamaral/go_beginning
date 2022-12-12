package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {

	// Struct para JSON
	p1 := produto{1, "Notebook", 1899.0, []string{"Promocao", "Eletronicos"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println("🚀 ~ file: json.go ~ line 16 ~ funcmain ~ p1Json : ", string(p1Json))

	//JSON para Struct
	var p2 produto
	jsonString := `{"id":2, "nome": "Caneta", "Preco": 8.90, "tags":["Papelaria"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println("🚀 ~ file: json.go ~ line 28 ~ funcmain ~ p2 : ", p2.Tags[0])
}
