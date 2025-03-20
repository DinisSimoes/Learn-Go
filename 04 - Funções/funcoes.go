package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// go permite ter varios returnos na mesma função
// é muito usado
// principalmente no tratamento de erros
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma:= n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(){
		fmt.Println(("função de exemplo"))
	}

	f()

	var f2 = func(txt string){
		fmt.Println(("função de exemplo com parametros: " + txt))
	}

	f2("texto exemplo")

	var f3 = func(txt string) string{
		fmt.Println(("função nova com parametros e retorno: " + txt))
		return txt
	}

	resultado := f3("texto exemplo")
	fmt.Println(resultado)


	resultadoSomma, resultadoSubtracao := (calculosMatematicos(10, 15))
	fmt.Println(resultadoSomma, resultadoSubtracao)

	//caso não queira o segundo retorno
	resultadoSomma2, _ := (calculosMatematicos(10, 15))
	fmt.Println(resultadoSomma2)

	_, resultadoSubtracao2 := (calculosMatematicos(10, 15))
	fmt.Println(resultadoSubtracao2)
	
}