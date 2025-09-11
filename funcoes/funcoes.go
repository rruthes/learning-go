package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}


func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
	
	var f = func(text string) string {
		fmt.Println(text)
		return text
	}
	
	resultado := f("texto da funcao 1")
	fmt.Println(resultado)

	// se voce por algum motivo nao precisa dos dois resultado no caso dessa funcao, voce pode usar um underline no lugar da variavel que nao vai ser usada.
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	_, resultado4 := calculosMatematicos(11, 22)
	fmt.Println(resultadoSoma, resultadoSubtracao)
	fmt.Println(resultado4)
}

 