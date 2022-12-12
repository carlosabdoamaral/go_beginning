package main

import (
	"math"
	"reflect"
)

func main() {
	//	Números inteiros
	println(1, 2, 1000)
	println("Literal interiro é", reflect.TypeOf(3200))

	//	Sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	println("Byte é", reflect.TypeOf(b))

	//	Com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	println("O valor máximo de int64 é:", i1)

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	println("O rune é", reflect.TypeOf(i2))

	//	Números reais (float32, float64)
	var x float32 = 49.99
	println("O tipo do literal 49.99 é", reflect.TypeOf(x))

	//	Boolean
	bo := true
	println("O tipo de bo é", reflect.TypeOf(bo))

	//	String
	s1 := "Olá sou uma string"
	println("O tamanho da string é", len(s1))

	//	String com múltiplas linhas
	s2 := `
	Essa string é
	com múltiplas
	linhas`
	println(s2)

	//	Char?
	char := 'a'
	println("O tipo de chat é", reflect.TypeOf(char))
	println(char)
}
