package main

import "fmt"

func main() {
	fmt.Println("estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("maior que 15")
	} else { // nao precisa ter o else. se nao cair na condicao do if, ele simplesmente é ignorado.
		fmt.Println("menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 { // if init: inicializacao de uma nova variavel dentro da estrutura do if. só existe dentro do if.
		fmt.Println("numero é maior que 0")
	} else {
		fmt.Println("numero menor que 0")
	}
}