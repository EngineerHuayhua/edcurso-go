package main

import (
	"github.com/EngineerHuayhua/edcurso-go/06/07paquetes/saludar"
	"fmt"
	"github.com/EngineerHuayhua/edcurso-go/06/07paquetes/despedida"
)

func main() {
	nombre := "Donato"
	saludar.Saludar(nombre)
	saludar.MeVes = "Esto es un texto asignado desde el main"
	fmt.Println(saludar.MeVes)
	despedida.Despedirse("Donato")
}
