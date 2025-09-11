package main

import "fmt"

func main() {
	fmt.Println("loops")
	i := 0
	// única estrutura de loop do go é o for.
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
	}

	fmt.Println(i)

	for j:= 0; j < 10; j++ {
		fmt.Println("Incrementando j")

	}

	nomes := [3]string{"João", "Rafael", "Vitor"}

	for _, nome := range nomes { // o underline é muito comumente usado para nao ter que criar a variavel "i" ou "indice"
		fmt.Println(nome)
	}

	usuario := map[string]string {
		"nome": "Rafael",
		"sobrenome": "Ruthes",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	// nao é possível usar range em structs, apenas nesses mostrados
}
