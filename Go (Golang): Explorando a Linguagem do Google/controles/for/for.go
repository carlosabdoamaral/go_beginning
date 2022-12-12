package main

import (
	"fmt"
	"time"
)

// Nao tem while
func main() {

	// Tipo while
	i := 0
	for i <= 10 { 
		fmt.Println(i)
		i++
	}

	// Com incrementador diferente
	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	// Com condicional
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print("Par ")
		} else {
			print("Impar ")
		}
	}

	// Laco infinito
	for {
		fmt.Println("Para sempre")
		time.Sleep(time.Second)
		// time.Sleep(time.Second * 5) // 5 segundos
	}
}
