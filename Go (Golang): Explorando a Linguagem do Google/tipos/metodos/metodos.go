package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println("🚀 ~ file: metodos.go ~ line 22 ~ funcmain ~ p1 : ", p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Pereira")
    fmt.Println("🚀 ~ file: metodos.go ~ line 28 ~ funcmain ~ p1 : ", p1.getNomeCompleto())
}