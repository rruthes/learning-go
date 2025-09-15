package main

import "fmt"

func closure() func() { // em go, como as funcoes sao tipos, elas podem ser usadas como retorno.
	texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure() // funcao closure serve para, dentro de uma funcao, referenciar variaveis que foram declaradas em uma outra funcao externa.
	funcaoNova() // aqui, então, a variavel "texto" usada é a declarada dentro da funcao externa "closure", e não a variável texto declarado dentro da funcao "main".
}
