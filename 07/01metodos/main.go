package main

import "fmt"

type Persona struct {
	nombre	string
	edad	int8
}

type Numero int

func (p Persona) saludar(nombre string) {
	p.nombre = nombre
	fmt.Println("HOla soy", p.nombre)
}

func (n Numero) presentarse() {
	fmt.Println("HOla yo soy un número")
}

func (p *Persona) asignarEdad(i int8) {
	if i >= 0 {
		p.edad = i
	} else {
		fmt.Println("NO es valida la edad negativa")
	}
}

func main() {
	var persona Persona
	//var numero Numero
	persona.saludar("Donatelo")
	//numero.presentarse()
	persona.asignarEdad(-4)
	fmt.Println(persona)
}