package main

import (
	"errors"
	"fmt"
)

func main() {
	// unica diferenca entre os tipos de int é o tamanho. 
	var numero int16 = 100
	fmt.Println(numero)

	// uint = unsigned int, ou seja, int sem sinal (não pode ser negativo)
	// byte = uint8
	var numero2 uint = 1000
	fmt.Println(numero2)

	// int32 = rune. esses nomes assim são simplesmente apelidos para certos tipos
	var numero3 rune = 12345
	fmt.Println(numero3)

	//float nao é um tipo que pode ser usado explicitamente. só pode ser usado desse jeito se for float32 ou float64.
	var numeroReal float32 = 12345.67
	fmt.Println(numeroReal)

	var str string = "texto"
	fmt.Println(str)

	// o mais proximo de char em go é:
	char := 'B'
	fmt.Println(char) // ele não printa o caracter, e sim o numero da tabela ascii

	// valor zero: todo tipo de dado tem seu valor zero (valor inicial). string, por exemplo, tem o valor zero como uma string vazia.
	var texto string
	fmt.Println(texto)

	var booleano bool = true
	fmt.Println(booleano)

	// no go existe um tipo de dado chamado erro. seu valor zero é nil (de nulo).
	var erro error = errors.New("Erro interno") // pacote nativo chamado errors.
	fmt.Println(erro)
}