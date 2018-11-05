package main

import "fmt"

func main() {
	//Slice
	// 1. Apuntador a un array
	// 2. Tamaño (no es fijo)
	// 3. Capacidad
	//var nombres []string
	// otra forma de crear slice con make (tipo de dato del slice, tamaño inicial, capacidad inicial)
	/*nombres := make([]string, 0)
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Donato")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Huayhua")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Huayhua")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Daniel")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Juan")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Freddy")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))

	// corrigiendo un error de datos en un slice conocido su indice
	nombres[5] = "Marcos"
*/
	nombres := []string {"Donato", "Huayhua", "Huayhua"}

	for i := 0; i < len(nombres); i++ {
		fmt.Printf("Nombre en el indice %d es %s\n",i, nombres[i])
	}
}
