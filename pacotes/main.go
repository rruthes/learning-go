package main

import (
	"modulo/auxiliar"
	"fmt"
)
func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()


}
// é possível rodar todo o código usando o código binario gerado. porem, a cada nova alteracao no codigo, é preciso rodar novamente o go build.
// como go nao é orientado a objetos, nao existe public/private/protected. para definir se sao publicas ou nao, é preciso usar a primeira letra do nome da funcao.
// se a primeira for maiuscula, ela é publica. se não, é privada.
// funcoes podem existir em dois arquivos dentro do mesmo pacote sem a necessidade de importacao.
// o comando go install instala o codigo binario na pasta raiz do projeto.