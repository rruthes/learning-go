package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func(u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func(u usuario) salvar(){ // jeito de criar um método para a struct usuario. o "u" é só o nome de uma variavel, entao poderia ser qqr nome
	fmt.Printf("salvando os dados do usuário %s no banco de dados\n", u.nome) // nesse caso, o placeholder vai referenciar o campo nome do usuario que chamou o método salvar.
}

func(u *usuario) aniversario() { // usando a referenciacao de ponteiro, o valor da idade do usuario especifico vai mudar permanentemente. caso nao seja esse o objetivo, da pra salvar uma copia em uma outra variavel.
	u.idade++
}

func main() {
	usuario1 := usuario{"rafael", 19}
	fmt.Println(usuario1)
	usuario1.salvar()
	mdiR := usuario1.maiorDeIdade()
	fmt.Println(mdiR)
	usuario2 := usuario{"lulu", 15}
	usuario2.salvar()
	mdiL := usuario2.maiorDeIdade()
	usuario2.aniversario()
	fmt.Println(usuario2.idade)
	fmt.Println(mdiL)
}