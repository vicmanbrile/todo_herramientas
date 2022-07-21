package main

type Herramienta struct{}

var Herramientas_Path string = "data/herramienta.csv"

func (h *Herramienta) NewFile() File {
	file := File{
		path: Herramientas_Path,
	}

	QuerySelector := []*struct {
		Descripcion string `csv:"descripcion"`
	}{}

	file.ReadFile(&QuerySelector)

	var data []string
	for i, _ := range QuerySelector {
		data = append(data, QuerySelector[i].Descripcion)
	}

	file.data = data

	return file

}
