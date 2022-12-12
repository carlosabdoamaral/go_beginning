package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 7 && n < 9:
		return "B"
	default:
		return "C"
	}
}

func main() {
	fmt.Println(notaParaConceito(7))
}
