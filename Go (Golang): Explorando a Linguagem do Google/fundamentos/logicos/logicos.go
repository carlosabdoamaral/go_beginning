package main

import "fmt"

func compras(t1, t2 bool) (bool, bool, bool) {
	comprarTv50 := t1 && t2
	comprarTv32 := t1 != t2 // Ou exclusivo
	comprarSorvete := t1 || t2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
