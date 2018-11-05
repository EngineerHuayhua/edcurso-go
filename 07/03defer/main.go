package main

import "fmt"

func main() {
	fmt.Println("Conectando a la base de datos...")
	defer cerrarDB()

	return

	fmt.Println("Consultamos información, set de datos")
	defer cerrarSetDeDatos()
}

func cerrarDB() {
	fmt.Println("Cerrando la base de datos")
}

func cerrarSetDeDatos() {
	fmt.Println("Cerrar el set de datos")
}