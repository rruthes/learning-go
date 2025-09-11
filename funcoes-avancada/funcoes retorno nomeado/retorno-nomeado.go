package main

import "fmt"

func calculosMatematicos(n1, n2 int) (/* esse é o retorno nomeado */ soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2
	return	// como os retornos foram previamente nomeados, não é preciso especificar novamente no return
}

func main() {
	soma, subtracao := calculosMatematicos(1, 2)
	fmt.Println(soma, subtracao)
}