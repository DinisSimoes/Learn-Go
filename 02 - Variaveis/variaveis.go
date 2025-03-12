package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1" // variavel com declaração explícita
	variavel2 := "Variavel 2" // variavel com declaração implícita
	variavel3 := 3
	variavel4, variavel5 := "Variavel 4", 5
	fmt.Println(variavel1)
	fmt.Println(variavel2, variavel3)
	fmt.Println(variavel4, variavel5)

	const const1 string = "Constante 1"
	fmt.Println(const1)

	//para trocar valores de valores - funciona só com variaveis do mesmo tipo
	variavel5, variavel3 = variavel3, variavel5
	fmt.Println(variavel5, variavel3)
}
