package main

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

// Producto contiene los datos de un articulo
type Producto struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
	CodigoBarras string
	Precio uint
}

func main() {
	db, err := gorm.Open("mysql", "root:dnt@/edcurso?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error en la conexion a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conecto a la base de datos")

	// Creamos la tabla productos
	//db.CreateTable(&Producto{})
	//fmt.Println("Se creo la tabla productos en la base de datos")

	//// Insertamos datos a la base de datos
	//db.Create(&Producto{
	//	CodigoBarras: 	"0100101012356767",
	//	Precio: 		5000,
	//	})

	// Recuperamos datos de la base de datos
	var p Producto
	db.First(&p, 2)
	fmt.Println("Producto:", p.CodigoBarras)
	fmt.Println("Precio:", p.Precio)

	p.Precio = 4500
	db.Save(&p)
	fmt.Println("El nuevo precio del producto es:", p.Precio)
}
