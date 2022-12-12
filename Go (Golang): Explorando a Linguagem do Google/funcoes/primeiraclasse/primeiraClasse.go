package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(soma(2, 3))
	fmt.Println(sub(2, 3))
}
