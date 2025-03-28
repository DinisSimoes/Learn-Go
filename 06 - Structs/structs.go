package main

import "fmt"

type usuario struct {
	nome string
	sobrenome string
	idade int
	endereco endereco
}

type endereco struct{
	lagradouro string
	numero int
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)
	u.nome = "Pedro"
	u.sobrenome = "Silva"
	u.idade = 20

	fmt.Println(u)

	u2 := usuario{"Maria", "Silva", 30, endereco{"Rua dos bobos", 0}}
	fmt.Println(u2)

	u3 := usuario{idade: 40}
	fmt.Println(u3)
	u4 := usuario{nome: "Joao"}
	fmt.Println(u4)
}