package main

type Trabajador struct {
	Nombre string `csv:"nombre"`
}

func (t *Trabajador) NewFile(path string) File{
	Trabajadores_Path := "data/trabajadores.csv"

	file := File{
		path: Trabajadores_Path,
		out: []t
	}

}
