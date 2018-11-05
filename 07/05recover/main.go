package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() {
		// recover() recupera lo que panic() lanzo
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f:", r)
		}
	}()
	fmt.Println("Llamando g.")
	g(5)
	fmt.Println("Retorna normalmente desde g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("AAAAAaaaaa")
		// panic() puede contener cualquier tipo de dato (string, int, ....)
		panic("El número no puede ser mayor a 3")
	}
	defer fmt.Println("Defer en la funcion g")
	fmt.Println("Imprimiento en g", i)
}