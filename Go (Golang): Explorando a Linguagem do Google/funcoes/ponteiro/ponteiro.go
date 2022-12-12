package main

import "fmt"

func inc1(n int) {
	n++
}

// Revisao: um ponteiro e representado por um *

func inc2(n *int) {
	// * e usado para desreferencias, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	// Passa uma copia do valor de n
	inc1(n) // Por valor
	fmt.Println(n)

	// & usado para obter o endereco da variavel
	// Altera de fato a variavel n e nao uma copia dela
	inc2(&n) // Por referencia
	fmt.Println(n)
}
