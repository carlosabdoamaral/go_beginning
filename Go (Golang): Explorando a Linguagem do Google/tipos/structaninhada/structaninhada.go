package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	uderID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {
	pedido := pedido{
		uderID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 12.10},
		},
	}

	fmt.Println("🚀 ~ file: structaninhada.go ~ line 24 ~ funcmain ~ pedido : ", pedido.valorTotal())
}
