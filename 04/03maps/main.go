package main

import "fmt"

func main() {

	//var names map[string]string
	//names := make(map[string]string)
	//names["es"] = "Español"
	//names["en"] = "Inglés"
	//names["fr"] = "Francés"
	//idioms := map[string]string{"es":"Español", "en":"Inglés", "fr":"Francés", "pt":"Portugués"}

	//fmt.Println(idioms)
	//delete(idioms, "en")

	//if idiom, ok := idioms["pt"]; ok {
	//	fmt.Println("la posicion pt si existe", idiom)
	//} else {
	//	fmt.Println("la posicion pt no existe", idiom)
	//}
	//fmt.Println(idioms)

	numeros := map[uint8]string{
		1: "Uno es el numero mas chiquito",
		4: "Es otro numero",
		28: "Es un numero entero positivo",
	}

	fmt.Println(numeros)
}
