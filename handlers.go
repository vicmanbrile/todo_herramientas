package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Datos a enviar
type Profile struct {
	Message      string
	Herramienta  []string
	Trabajadores []string
}

func Home(w http.ResponseWriter, rq *http.Request) {

	if rq.Method != http.MethodPost {
		home := template.Must(template.ParseFiles("templates/home.html"))

		Herra := &Herramienta{}

		HerraF := NewFile(Herra)

		Trab := &Trabajadores{}

		TrabF := NewFile(Trab)

		dat := Profile{
			Message:      "Herramienta Trabajadores",
			Herramienta:  HerraF.data,
			Trabajadores: TrabF.data,
		}

		err := home.Execute(w, dat)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	datails := Todo{
		Trabajador:  rq.FormValue("trabajador"),
		Herramienta: rq.FormValue("herramienta"),
		Time:        time.Now(),
	}

	fmt.Println(datails)

	fmt.Fprint(w, "Thank You")

}

type Todo struct {
	Trabajador  string
	Herramienta string
	Time        time.Time
}
