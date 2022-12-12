package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //Compilador conta quantos items sao

	// Crio a variavel i e numero com base no range do array de numeros
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, numero := range numeros {
		fmt.Printf("\n%d", numero)
	}
}
