package main

import "fmt"

func main() {
	//anonima := func () {
	//	fmt.Println("Me imprimo estando en una variable llamada anonima")
	//}
	//
	//anonima()
	secuencia1 := secuencia(1)
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println("------------")
	secuencia2 := secuencia(2)
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
}

func secuencia(n int32) func() int32 {
	var x int32
	return func() int32 {
		x += n
		return x
	}
}