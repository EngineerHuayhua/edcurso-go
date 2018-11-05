package main

import "fmt"

func main() {
	////Arrays
	//var nombres [3]string
	//nombres[0] = "Donato"
	//nombres[1] = "Huayhua"
	//nombres[2] = "Huayhua"

	nombres := [3]string{"Donato", "Huayhua", "Huayhua"}
	nombre := nombres[1]

	//tamaño del array
	size := len(nombres)

	fmt.Println(size)
	fmt.Println(nombre)

	//imprimir posiciones del array
	//fmt.Println(nombres[inicio:final(excluyente)])
	// se imprime las posiciones 1 y 2 se excluye la posicion 3
	fmt.Println(nombres[0:2])
}