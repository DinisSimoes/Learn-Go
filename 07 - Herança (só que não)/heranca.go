package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

type estudante struct {
	pessoa
	curso string
	faculta string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Joaquim", "Silva", 20}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
}