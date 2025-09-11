package main

import "fmt"

func soma(numeros ...int) int { // vai receber de 0 a n ints
	total := 0
	for _, numero := range(numeros){
		total += numero
	}

	return total
} 

func escrever(texto string, numeros ...int) {
	for _, numero := range(numeros){
		fmt.Println(texto, numero)
	}
}


func main(){
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 100, 200, 300, 500, 600)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 1, 2, 3, 4, 5, 6, 7)
}