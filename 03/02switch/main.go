package main

import "fmt"

func main() {
	//var a uint8
	//a = 6
	fmt.Println("¿Qué día es Hoy?\n")

	/*
	switch a {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 27:
		fmt.Println("Domingo")
	default:
		fmt.Println("Fuera de rango de los días")
	}*/
	/*
	switch a {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Estas entre la semana")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Estas en fin de semana :)")
	default:
		fmt.Println("No es un día válido")
	}
	*/

	switch a := 3; {
	case a >= 0 && a <= 5:
		fmt.Println("Estas entre la semana")
	case a >= 6 && a <= 7:
		fmt.Println("Estas en fin de semana")
	default:
		fmt.Println("No es un dia valido")
	}
}