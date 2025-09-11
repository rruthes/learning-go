package auxiliar

import "fmt"

// por questao de boas praticas, é bom colocar um comentario em cima de funcoes publicas exportadas
func Escrever() {
	fmt.Println("Escrevendo do arquivo do pacote auxiliar")
	escrever2()
}