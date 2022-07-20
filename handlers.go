package main

// Datos a enviar
type Profile struct {
	Message string
	Herramienta []string
	Trabajadores[]string
}

func Home(w http.ResponseWriter, rq *http.Request) {

	if rq.Method != http.MethodPost {
			home, _ := template.ParseFiles("template/home.html")

		dat := Profile{
			Message: "Herramienta Trabajadores",
			Herramienta: []string{"Martillo", "Taladro", "Niveleta", "Flexometro"},
			Trabajadores: []string{"Victor Briseño L.", "Teo", "Toñin", "Jorge", "Carlos"},
		}
		
		err := home.Execute(w, dat)
		
		if err != nil {
			fmt.Println(err)
		}
		return
  }

	datails := Todo{
		Trabajador: rq.FormValue("trabajador"),
		Herramienta: rq.FormValue("herramienta"),
		Time: time.Now(),
	}

	fmt.Println(datails)
	
	fmt.Fprint(w, "Thank You")

}
