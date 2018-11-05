package main

import "fmt"
// Persona es una estructura
type Persona struct {
	Nombre 	string
	Edad	uint8
	Email	[]string
}

func main() {
	/*
	var persona1 Persona

	persona1.Nombre = "Donato"
	persona1.Edad	= 30

	fmt.Println(persona1.Nombre)
	*/
	/*
	persona2 := Persona{}
	persona2.Nombre = "Pedro"
	persona2.Edad	= 50

	fmt.Println(persona2)
	*/

	emails := []string{"a@b.com", "b@b.com"}
	persona2 := Persona{
		"Donato",
		45,
		emails,
	}
	fmt.Println(persona2)
}
