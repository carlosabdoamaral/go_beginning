package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao":      11325.45,
		"Gabriela Silva": 18232.45,
		"Pedro Junior":   92935.45,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
