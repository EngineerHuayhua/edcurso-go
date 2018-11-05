package main

import "fmt"

func main() {
	var a int
	var b int8

	a = 124434
	b = 23

	//casting, cambiar temporalmente el tipo de una variable
	c := a + int(b)
	fmt.Println(c)
}