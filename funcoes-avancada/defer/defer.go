package main

import "fmt"

func funcao1(){
	fmt.Println("executando a funcao 1")
}

func funcao2(){
	fmt.Println("executando a funcao 2")
}

func alunoAprovado(n1, n2 float32) bool{
	defer fmt.Println("Média calculada. Resultado será calculado") // muito util principalmente para lidar com banco de dados
	fmt.Println("entrando na funcao de aluno aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		// vai executar o print com defer antes do próximo retorno.
		return true
	} 
	return false
}

func main(){
	fmt.Println(alunoAprovado(7, 8))
}
