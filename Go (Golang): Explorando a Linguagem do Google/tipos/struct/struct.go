package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Metodo: Funcao com receiver(receptor)
func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 Produto
	produto1 = Produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println("🚀 ~ file: struct.go ~ line 29 ~ funcmain ~ produto1 : ", produto1, produto1.precoComDesconto())

	produto2 := Produto{"Caneta", 2, 0.5}
	fmt.Println("🚀 ~ file: struct.go ~ line 27 ~ funcmain ~ produto2 : ", produto2, produto1.precoComDesconto())
}
