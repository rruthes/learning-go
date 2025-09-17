package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // sincroniza duas go routines

	waitGroup.Add(2) // informa que tem 2 go routines que fazem parte do wait group

	go func() {
		escrever("olá mundo")
		waitGroup.Done() // o done faz com que o contador do add seja diminuido em 1, ou seja, agora ele so espera 1 em vez de 2.
	}()
	go func() {
		escrever("Programando em go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}