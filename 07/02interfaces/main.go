package main

import (
	"github.com/EngineerHuayhua/edcurso-go/07/02interfaces/animales"
	"fmt"
)

func main() {
	var p animales.Perro
	var g animales.Gato
	fmt.Println("El valor de p es:", p)
	p.Nombre = "Firulais"
	g.Nombre = "Manchas"
	//AdoptarPerro(p)
	//AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}
/*
func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}

func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}
*/

func AdoptarMascota(m animales.Mascota) {
	fmt.Println(m.Comunicarse)
	m.Comunicarse()
}