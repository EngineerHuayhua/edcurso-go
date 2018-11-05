package saludar

import "fmt"

// MeVes es una variable para utilizar desde otro paquete
var MeVes string
var noMeVes string
// muy importante la S debe ser mayuscula y esta funcion se llamara exportada
// y puede ser llamadado por cualquier otro paquete
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func noVisible() {
	fmt.Println("No me ven desde otro paquete")
}