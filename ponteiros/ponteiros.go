package main

import "fmt"

// apontam para um endereço de memória
func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++ 
	fmt.Println(variavel1, variavel2)
	// quando um ponteiro é criado, o valor da variável não é jogado dentro dele, e sim o endereço da memória onde está armazenada a variável.

	var variavel3 int 
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // & indica o endereço de memória da variavel em questão

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // desreferenciacao: asterisco faz com que o valor do endereço de memória seja lido.
	
	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)

}
