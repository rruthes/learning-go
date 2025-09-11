package main // pacotes só podem ser executados se forem do pacote main
import (
	"fmt"
)

func main() { // a funcao main precisa ser declarada dentro de um arquivo do pacote main
	var a = 1
	var b = 2
	var resultado = a + b

	fmt.Println(resultado)
}