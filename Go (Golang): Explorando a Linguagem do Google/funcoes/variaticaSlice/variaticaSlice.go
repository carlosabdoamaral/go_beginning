package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")

	for i, aprovado := range aprovados {
		fmt.Println(i, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Joao", "Pedro"}
	imprimirAprovados(aprovados...)
}
