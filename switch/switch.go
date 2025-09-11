package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1: // é igual a: case numero == 1
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	default:
		return "não sei o dia kkkk"
	}
		

}

func main(){

	dia := diaDaSemana(30)
	fmt.Println(dia)
}