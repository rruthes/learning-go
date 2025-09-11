package main

import "fmt"

func main() {

	retorno := func (texto string) string {
		return fmt.Sprintf("recebido -> %s", texto) // Sprintf: recebe uma string e concatena com outros tipos de dados (%s para strings, por exemplo)
	}("passando parametro") // é executada logo após a declaração
	fmt.Println(retorno)
}
