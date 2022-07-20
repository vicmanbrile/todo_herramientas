package main

type ToDoList struct { // Our example struct, you can use "-" to ignore a field
	Herramienta string `csv:"Herramienta"`
	Trabajadores string `csv:"trabajadores"`
	Tiempo DateTime `csv:"tiempo"`
}

func (t *ToDoLists) NewFile(path string) File{
	Trabajadores_Path := "data/trabajadores.csv"

	file := File{
		path: Trabajadores,
		out: []t
	}

}
