package main

import (
	"errors"
	"fmt"
)

func main() {
	
	//NUMEROS

	//nr inteiro
	//int8 => numero até 8 bits
	//int16 => numero até 16 bits
	//int32 => numero até 32 bits
	//int64 => numero até 64 bits
	//int => cria com base no sistema operacional, se for 32 bits cria um int32, se for 64 bits cria um int64
	//o int tbm suporta numero negativos
	//unit => numero sem sinal (unsygned int) - sempre tem de ser "positivo"
	fmt.Println("INTS")
	var numero int8 = 100
	var numero2 int16 = 1000
	var numero3 int32 = 10000
	var numero4 int64 = 100000
	var numero5 int = 20000
	var numero6 uint32 = 100
	
	fmt.Println(numero, numero2, numero3, numero4, numero5, numero6)

	//ALIAS
	//int32 = RUNE
	fmt.Println("")
	fmt.Println("RUNES")
	
	var numero7 rune = 12345
	fmt.Println(numero7)

	// int8 = BYTE
	var numero8 byte = 123
	fmt.Println(numero8)

	//nr reais, numeros com casas decimais
	//para separar as casas decimais é sempre com "."
	//mesmo logica dos numeros inteiros
	fmt.Println("")
	fmt.Println("FLOATS")
	
	var float1 float32 = 123.45 
	var float2 float64 = 123.456
	float3:= 12345.67

	fmt.Println(float1, float2, float3)

	//FIM NUMEROS

	//STRINGS
	fmt.Println("")
	fmt.Println("STRINGS")
	

	var str string = "Texto"
	fmt.Println(str)
	str2:="Texto 2"
	fmt.Println(str2)

	char:='B'
	fmt.Println(char)
	//loga o valor da tabela asp
	//converte para um int/rune

	//FIM STRINGS

	//VALOR 0
	//para todos os tipos de dados vem com o valor a 0 ou empty
	//int = 0
	//string = ""
	fmt.Println("")
	fmt.Println("VALOR 0")
	
	var texto2 string
	fmt.Println(texto2)

	var texto int
	fmt.Println(texto)

	//END VALOR 0

	//BOOLEAN
	//true ou false
	fmt.Println("")
	fmt.Println("BOOLEANS")
	
	var booleano1 bool = true
	var booleano2 bool = false
	var booleano3 bool
	fmt.Println(booleano1, booleano2, booleano3)

	//FIM BOOLEAN

	//ERROR
	fmt.Println("")
	fmt.Println("ERROS")

	var error1 error
	fmt.Println(error1)//nil
	var error2 error = errors.New("Erro de exemplo")
	fmt.Println(error2)//nil

	//FIM ERROR
}