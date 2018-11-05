package main

import "fmt"

func saludar(nombre string, edad uint8) {
	fmt.Printf("HOla %s tienes %d años de edad\n", nombre, edad)

	if edad > 30 {
		return
	}
	fmt.Println("Eres joven")
}

func main() {
	saludar("Donatelo", 25)
}
