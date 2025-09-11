package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome" : "pedro",
		"sobrenome": "silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome" : {
			"primeiro": "João",
			"ultimo" : "Paulo",
		},
		"curso": {
			"nome": "engenharia",
			"campus": "scs",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") // comando para remover uma chave
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string { // forma de adicionar uma chave a um map ja criado
		"nome": "Gemeos",
	}
	fmt.Println(usuario2)
}