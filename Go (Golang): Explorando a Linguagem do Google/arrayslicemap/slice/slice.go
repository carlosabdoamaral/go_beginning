package main

import (
	"fmt"
	"reflect"
)

// Slice nao e um array! Slice define um pedaco de um array
func main() {
	a1 := [3]int{1, 2, 3} //Array
	s1 := []int{1, 2, 3}  // Slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // Do indice 1 ate o 3 nao incluindo o 3
	fmt.Println(a2, s2)

	s3 := a2[:2] // Novo slice porem aponta para o mesmo array
	fmt.Println(s3)

	// Voce pode imaginar um slice como
	// Tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
