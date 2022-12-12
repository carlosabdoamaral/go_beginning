package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {

	// Declara a variavel i como de controle
	// Depois faz a condicao
	// Tambem e suportado no switch
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Ganhou")
	}
}
