package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1[5] int
	fmt.Println(array1)

	array1[4] = 10
	fmt.Println(array1)

	array2 := [5] int {1, 2, 3, 4, 5}
	fmt.Println(array2)

	//os tres pontos significa que ele vai inferir o tamanho passado nas chaves
	//na pratica não é muito usado
	array3 := [...] int {1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slices
	//mais flexivel e por isso muito usado em go, mas sempre ficam "presos" a um tipo de dados
	//não se pode armazenar um int e uma string no mesmo slice
	slice := [] int {1, 2, 3, 4, 5}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 6)
	fmt.Println(slice)

	//pega o index 1 e 2, o ultimo nr é excludente
	slice2 := array2[1:3]
	fmt.Println(slice2)

	//ao mudar o valor do array, muda o slice
	array2[1] = 100
	fmt.Println(slice2)//mudou pq é uma referencia, um ponteiro
}