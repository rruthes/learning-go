package main

import "fmt"

func main() {
	variavel2 := "variável 2" 
	// o go não deixa você declarar uma variável e não usá-la
	// o tipo da variavel pode estar implicito ou explicito
	var variavel1 string = "Variável 1"
	var (
		variavel3 string = "variável 3"
		variavel4 string = "oiiii"
	)

	variavel5, variavel6 := 123, "teste"
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	}	
