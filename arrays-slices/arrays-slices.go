package main

import "fmt"

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]string
	fmt.Println(array1)

	array2 := [5]string{"Oi", "hello", "hola", "ni hao", "bonjour"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // isso permite apenas que voce nao digite o tamanho da array, mas nao torna ela dinamica.
	fmt.Println(array3)

	slice := []int{13, 11, 12, 15, 16, 17} // slice parece um array, mas na verdade ele é como um ponteiro que aponta para um array.
	fmt.Println(slice)

	slice = append(slice, 18) // forma de adicionar valores sem ter que criar de novo o slice
	fmt.Println(slice)

	// na realidade, o slice é mesmo uma "fatia" de um array. pode ser criado dessa forma.
	slice2 := array2[1:3] // pega os indices 1 e 2
	fmt.Println(slice2)

	array2[2]= "ciao"
	fmt.Println(slice2)

	// arrays internos
	fmt.Println("=============")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // se o cap estourar, ele cria um novo array para ser referenciado e dobra o tamanho dele.

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(len(slice4))
}