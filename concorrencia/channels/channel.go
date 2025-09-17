package main

import (
	"fmt"
	"time"
)

func main () {
	canal := make(chan string) // canais sao a forma mais correta de sofisticada de sincronizar go routines.

	go escrever("Olá mundo!", canal)

	fmt.Println("depois da funcao escrever")

	for mensagem := range canal {
		fmt.Println(mensagem) // toda mensagem existente no canal vai ser printada na tela.
	}
	fmt.Println("fim do programa")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++{
		canal <- texto // mandando um valor (texto) para dentro do canal. caso a seta fosse ao contrario, o canal que estaria mandando um valor.
		time.Sleep(time.Second)
	}
	close(canal) // depois que terminar o loop, fecha o canal. sem isso, ele cai no deadlock (canal vai ficar aberto p sempre esperando um valor que nunca vai chegar)
}
