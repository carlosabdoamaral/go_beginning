package main

import "fmt"

func main() {
	print("Mesma")
	print(" linha")

	println(" Nova")
	println("Linha.")

	x := 3.141516
	fmt.Printf("O valor de x é %.2f.", x)

	xs := fmt.Sprint(x)
	println("O valor de x é " + xs)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
