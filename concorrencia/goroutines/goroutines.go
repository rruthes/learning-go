package main

import (
	"fmt"
	"time"
)

// CONCORRENCIA É DIFERENCA DE PARALELISMO
// o paralelismo exige que o processador tenha dois nucleos ou mais, para poder executar os dois programas ao mesmo tempo.
// a concorrencia pode ou não fazer com que os dois programas sejam executados ao mesmo tempo, funcionando como um revezamento de tempo de execucao.

func main() {
	go escrever("Olá mundo!") // nesse caso aqui, um programa normal (sem concorrencia) executaria a primeira linha, e só executaria a segunda quando a primeira fosse executada 100%. nesse caso, a linha 14 nunca seria executada.
	escrever("Programando em go") // go routine faz com que a execucao da linha de cima nao seja 100% concluida antes de continuar executando o resto do código.
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}