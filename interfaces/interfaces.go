package main

import (
	"fmt"
	"math"
)

// digamos que eu quero fazer uma funcao de calculararea, que seja usada pelos dois structs. eu nao conseguiria fazer isso, dado que não seria possivel passar simultaneamente as variaveis do struct circulo e do struct retangulo.
// isso é possível usando as interfaces.
type forma interface { // recebe apenas assinaturas de métodos
	area() float64
}

func escreverArea(f forma){
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	lado1 float64
	lado2 float64
}

func(r retangulo) area() float64{
	return r.lado1 * r.lado2
}

type circulo struct {
	raio float64
}

func(c circulo) area() float64 {
	return c.raio * c.raio * math.Pi
}

func main() {
	circulo1 := circulo{10}
	retangulo1 := retangulo{10, 15}
	escreverArea(circulo1)
	escreverArea(retangulo1)
}
