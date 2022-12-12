package main

import "fmt"

func main() {
	// map[chave]valor

	// Mapas devem ser inicializados
	// var aprovados map[int]string
	aprovados := make(map[int]string)
	aprovados[1234567898] = "Maria"
	aprovados[2378961298] = "Pedro"
	aprovados[1982133321] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[1982133321])
	delete(aprovados, 1982133321)
}
