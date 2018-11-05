package main

import "fmt"

func main() {
	saludarVarios(20, "Donato", "Pedro", "Simon", "Pancho")
}

func saludarVarios(edad uint8, nombres...string) {
	//fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("Hola ", v, "tienes ", edad, "años")
	}
}