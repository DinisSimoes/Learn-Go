package main

import "fmt"

func main() {
	//ARITMETICOS
	var soma = 1 + 2            // 3
	var subtracao = 5 - 3       // 2
	var multiplicacao = 2 * 3   // 6
	var divisao = 10 / 2        // 5
	var restodadivisao = 10 % 2 // 0
	
	fmt.Println(soma, subtracao, multiplicacao, divisao, restodadivisao)

	var numero1 int16 = 10
	//var numero2 int32 = 25
	var numero2 int16 = 25

	//Erro porque os dois numeros sao diferentes, o tipo deles é diferente mesmo os dois sendo nr
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//FIM DOS ARITMETICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel:= "string2";

	fmt.Println(variavel1, variavel)

	//FIM DA ATRIBUIÇÃO

	//RELACIONAIS
	fmt.Println(1 == 1) //true
	fmt.Println(1 != 1) //false
	fmt.Println(1 > 1) //false
	fmt.Println(1 < 1) //false
	fmt.Println(1 >= 1) //true
	fmt.Println(1 <= 1) //true

	//FIM DOS RELACIONAIS

	//LOGICOS
	fmt.Println(true && false) //false
	fmt.Println(true || false) //true
	fmt.Println(!true) //false
	//FIM DOS LOGICOS

	//UNARIOS
	numero:=10
	numero++
	fmt.Println(numero) //11
	numero+=15
	fmt.Println(numero) //26
	numero--
	fmt.Println(numero) //25
	numero-=10
	fmt.Println(numero) //15
	fmt.Println(!true) //false
	fmt.Println(!false) //true
	//FIM DOS UNARIOS

	//TERNARIOS
	var idade = 18
	var podeBeber string
	if idade >= 18 {
		podeBeber = "Pode beber"
	} else {
		podeBeber = "Nao pode beber"
	}
	fmt.Println(podeBeber)
	//FIM DOS TERNARIOS
}