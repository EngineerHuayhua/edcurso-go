package main

import "fmt"

func main() {
	a := 3
	// cuando se pasa solo el valor se dice paso por valor
	fmt.Println("Antes de duplicar, a = ", a)
	fmt.Println("direccion de memoria es", &a)
	duplicar(&a)
	fmt.Println("Despues de duplicar, a = ", a)
}

//paso por referencia
func duplicar(x *int) {
	*x *= 2
	fmt.Println("Dentro de la funcion duplicar x vale:", *x)

}