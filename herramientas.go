package main

type Trabajador struct {
	Descripcion string `csv:"descripcion"`
}

func NewFile(path string) File{
	Herramientas_Path := "data/herramienta.csv"

	file := File{
		path: Herramientas_Path,
		out: []t{}
	}

}
