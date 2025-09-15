package main

import "fmt"

func generica(interf interface{}) { // dessa forma, qualquer coisa está dentro dos requisitos da interface, fazendo com que ela possa receber qualquer tipo de valor
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
}