package main

import "fmt"

func inverterSinal(numero int)int {
	return numero*-1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1 // o uso dos asteriscos serve para interagir com o valor do endereço de memória
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero) // quando a funcao inverterSinal referencia a variável numero, ela não altera o endereço de memória da variável, fazendo com que seu valor não seja alterado. ela simplesmente faz uma cópia e atribui a variável numeroInvertido.
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
