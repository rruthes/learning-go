package main

import "fmt"

// struct é basicamente uma colecao de campos.
type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

// da pra ter struct dentro de struct
type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario 
	u.nome, u.idade = "Rafael", 19
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0} 
	// o jeito mais rapido de fazer é usando a inferencia de tipos
	usuario2 := usuario{"Rafael", 19, enderecoExemplo}
	fmt.Println(usuario2)

	// quando voce quer criar um usuario mas nao tem todos os dados
	usuario3 := usuario{nome: "Rafael"}
	fmt.Println(usuario3)

}
