package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println("🚀 ~ file: tipoInterface.go ~ line 14 ~ funcmain ~ coisa : ", coisa)

	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println("🚀 ~ file: tipoInterface.go ~ line 18 ~ funcmain ~ coisa2 : ", coisa2)

	coisa2 = true
	fmt.Println("🚀 ~ file: tipoInterface.go ~ line 21 ~ funcmain ~ coisa2 : ", coisa2)

	coisa2 = curso{"Texto aquiii!"}
	fmt.Println("🚀 ~ file: tipoInterface.go ~ line 24 ~ funcmain ~ coisa2 : ", coisa2)
}
