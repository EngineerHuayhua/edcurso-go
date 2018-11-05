package main

import "fmt"

//func suma(a, b int8) int8 {
//	return a + b
//}
//
//func tipoEdad(edad int8) string {
//	var tipo string
//
//	switch {
//	case edad < 12:
//		tipo = "niñez"
//	case edad < 18:
//		tipo = "adolescencia"
//	default:
//		tipo = "madurez"
//	}
//
//	return tipo
//}

func maxymin(n []int8) (max, min int8) {

	max, min = n[0], n[0]

	for _, v := range n {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return
}

func main() {
	//var n1, n2 int8
	//n1 = 4
	//n2 = 35
	//r := suma(n1, n2)
	//fmt.Println("El resultado es: \n", r)
	//
	//fmt.Printf("tu estas en la etapa de %s",tipoEdad(20))

	num := []int8{23,45,2,3,6,3,44,67,56}
	maximo, minimo := maxymin(num)
	fmt.Printf("el maximo es: %d\n",maximo)
	fmt.Printf("el minimo es: %d",minimo)
}
