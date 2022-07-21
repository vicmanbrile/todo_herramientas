package main

type Trabajadores struct{}

var Trabajadores_Path string = "data/trabajadores.csv"

func (h *Trabajadores) NewFile() File {
	file := File{
		path: Trabajadores_Path,
	}

	QuerySelector := []*struct {
		Nombre string `csv:"nombre"`
	}{}

	file.ReadFile(&QuerySelector)

	var data []string
	for i, _ := range QuerySelector {
		data = append(data, QuerySelector[i].Nombre)
	}

	file.data = data

	return file

}
