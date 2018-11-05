package main

import (
	"net/http"
	"html/template"
	"log"
)

// Persona contiene los datos de una persona
type Persona struct {
	Nombre	string
	Edad	uint8
}

func renderizar(w http.ResponseWriter, r *http.Request)  {
	p := &Persona{"Donato", 38}
	t, err := template.ParseFiles("./views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, p)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderizar)

	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
	log.Println("Ejecución terminada, ya no esta corriendo el servidor")
}
