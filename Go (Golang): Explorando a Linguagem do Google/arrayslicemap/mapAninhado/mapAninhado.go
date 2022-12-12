package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva:": 12234,
			"Guga:":           23032,
		},
		"J": {
			"Jose Junior": 12421,
		},
	}

	delete(funcPorLetra, "G")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}
