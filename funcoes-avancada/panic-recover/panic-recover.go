package main

import (
	"fmt"
)

func recuperarExecucao() {
	if r := recover(); r != nil{ // implementação mais comum do recover
		// se não cair no panic, ele nem é executado.
		fmt.Println("Execução foi recuperada com sucesso.")
	}
}

func alunoAprovado(n1, n2 float32)bool{
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media >6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6") // panic mata a execucao do programa e executa o que foi passado. não continua a execução após ele.
	// ele pode chamar somente as funções que tem defer.
}

func main(){
	fmt.Println(alunoAprovado(8, 6))
	fmt.Println("Pós-execução")
}
