package main

import "fmt"

// o go nao tem heranca nativa pq nao é orientado a objetos.
// isso é o mais proximo de heranca
type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}
type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Silva", 19, 193}
	fmt.Println(p1)

	e1:= estudante{p1, "Engenharia", "Mauá"}
	fmt.Println(e1)	
}